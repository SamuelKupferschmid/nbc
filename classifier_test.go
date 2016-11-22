package nbc

import (
	"testing"
)

func TestTrainIncreaseLabelLen(t *testing.T) {
	c := &Classifier{}

	if len(c.Labels()) != 0 {
		t.Fail()
	}

	c.Train([]LabelItem{
		LabelItem{
			Label:   "test",
			Content: SplitText("test text"),
		},
	})

	if len(c.Labels()) != 1 {
		t.Fail()
	}
}

func TestSameForValidation(t *testing.T) {
	d := []LabelItem{
		LabelItem{
			Label:   "test1",
			Content: SplitText("example text"),
		},
	}

	c := &Classifier{}

	c.Train(d)
	p := c.Validate(d)

	if p.Precision < 1 || p.Recall < 1 {
		t.Fail()
	}
}

func TestTrainWithContentOverlapping(t *testing.T) {

	//TODO not ready yet
	return

	d := []LabelItem{
		LabelItem{
			Label:   "label1",
			Content: SplitText("example of label1 test"),
		},
		LabelItem{
			Label:   "label2",
			Content: SplitText("label2 test example"),
		},
	}

	c := &Classifier{}

	c.Train(d)

	m := c.GetMatches(SplitText("label1 example test"))

	if len(m) != 2 {
		t.Fail()
	}

	res := m[0]

	if m[0].Probability < m[1].Probability {
		res = m[1]
	}

	if res.Label != "label1" {
		t.Fail()
	}
}
