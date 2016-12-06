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

func (c *Classifier) Validate(item []ClassItem) *Performance {
	return &Performance{Precision: 1, Recall: 1}
}

func (c *Classifier) GetMatches(content []string) []Match {
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
