package nbc

type Performance struct {
	Precision float64
	Recall    float64
}

func (p *Performance) F1() float64 {
	return (2 * p.Precision * p.Recall) / (p.Precision + p.Recall)
}
