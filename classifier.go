package nbc

type Classifier struct {
	trainingSet map[string]map[string]float64
}

func (c *Classifier) Train(items []ClassItem) {
	c.trainingSet = make(map[string]map[string]float64)

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

	//TODO normalization
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
	return nil
}

type Match struct {
	Class       string
	Probability float64
}

type ClassItem struct {
	Class   string
	Content []string
}
