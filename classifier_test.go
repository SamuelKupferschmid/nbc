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
