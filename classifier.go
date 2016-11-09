package nbc

type Classifier struct {
	labels []string
}

func NewClassifier(labels []string) *Classifier {
	return &Classifier{labels: labels}
}

func (c *Classifier) AddTrainingItem(label string, content []interface{}) {

}
