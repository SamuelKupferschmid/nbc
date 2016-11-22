package nbc

type Classifier struct {
	trainingSet map[string]map[string]float64
}

func (c *Classifier) Train(items []LabelItem) {
	c.trainingSet = make(map[string]map[string]float64)

	//fill labels and values
	for _, val := range items {
		if _, ok := c.trainingSet[val.Label]; !ok {
			c.trainingSet[val.Label] = make(map[string]float64)
		}
	}

	cnt := 0

	for _, val := range items {
		for _, w := range val.Content {
			for l, _ := range c.trainingSet {

				if _, ok := c.trainingSet[l][w]; !ok {
					c.trainingSet[l][w] = 0
				}

				if val.Label == l {
					c.trainingSet[l][w]++
					cnt++
				}
			}
		}
	}

	//TODO normalization
}

func (c *Classifier) Labels() []string {

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

func (c *Classifier) Validate(item []LabelItem) *Performance {
	return &Performance{Precision: 1, Recall: 1}
}

func (c *Classifier) GetMatches(content []string) []Match {
	return nil
}

type Match struct {
	Label       string
	Probability float64
}

type LabelItem struct {
	Label   string
	Content []string
}
