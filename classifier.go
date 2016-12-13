package nbc

import (
	"sort"
)

type Classifier struct {
	trainingSet map[string]map[string]float64
	minWeight   float64
}

func (c *Classifier) Train(items []ClassItem) {
	c.trainingSet = make(map[string]map[string]float64)

	c.minWeight = 0.05

	//fill classes and values
	for _, val := range items {
		if _, ok := c.trainingSet[val.Class]; !ok {
			c.trainingSet[val.Class] = make(map[string]float64)
		}
	}

	cnt := 0

	for _, val := range items {
		for _, w := range val.Content {
			for l, _ := range c.trainingSet {

				if _, ok := c.trainingSet[l][w]; !ok {
					c.trainingSet[l][w] = 0
				}

				if val.Class == l {
					c.trainingSet[l][w]++
					cnt++
				}
			}
		}
	}

	for _, d := range c.trainingSet {
		for i, _ := range d {
			if d[i] == 0 {
				//clamp to minWeight to avoid multiply by zero
				d[i] = c.minWeight
			} else {
				d[i] /= float64(len(items))
			}
		}
	}
}

func (c *Classifier) Classes() []string {

	if c.trainingSet == nil {
		return []string{}
	}

	l := make([]string, len(c.trainingSet))
	i := 0

	for k, _ := range c.trainingSet {
		l[i] = k
		i++
	}

	return l
}

// Validates adjusts the threshold for each labels n times. Where n is given by iterations.
// It returns an Array of float64 which gives the accuracy between each iteration/adjustmens as well as before and after the Validation. (len() == iterations + 1)
func (c *Classifier) Validate(items []ClassItem, iterations int) []float64 {
	acc := make([]float64, iterations+1)

	for i := 0; i <= iterations; i++ {
		acc[i] = c.Accuracy(items)
	}

	return acc
}

func (c *Classifier) Accuracy(items []ClassItem) float64 {
	matches := 0

	for _, iv := range items {
		m := c.PredictBest(iv.Content)

		if m.Class == iv.Class {
			matches++
		}
	}

	return float64(matches) / float64(len(items))
}

func (c *Classifier) PredictBest(content []string) Match {
	return c.PredictAll(content)[0]
}

func (c *Classifier) PredictAll(content []string) []Match {
	matches := make(Matches, len(c.trainingSet))

	i := 0

	for k, _ := range c.trainingSet {
		matches[i] = Match{
			k,
			1,
		}

		i++
	}

	sum := 0.

	for i, _ := range matches {
		for _, w := range content {
			matches[i].Probability *= c.trainingSet[matches[i].Class][w]
		}

		sum += matches[i].Probability
	}

	for i, _ := range matches {
		matches[i].Probability /= sum
	}

	sort.Sort(matches)

	return matches
}

type Match struct {
	Class       string
	Probability float64
}

type ClassItem struct {
	Class   string
	Content []string
}

type Matches []Match

func (m Matches) Len() int {
	return len(m)
}

func (m Matches) Less(i, j int) bool {
	return m[i].Probability > m[j].Probability
}

func (m Matches) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
