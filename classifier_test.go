package nbc

import (
	"testing"
)

func TestTrainIncreaseClassesLen(t *testing.T) {
	c := &Classifier{}

	if len(c.Classes()) != 0 {
		t.Fail()
	}

	c.Train([]ClassItem{
		ClassItem{
			Class:   "test",
			Content: SplitText("test text"),
		},
	})

	if len(c.Classes()) != 1 {
		t.Fail()
	}
}

func TestSameForValidation(t *testing.T) {
	d := []ClassItem{
		ClassItem{
			Class:   "test1",
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
	t.SkipNow()

	d := []ClassItem{
		ClassItem{
			Class:   "label1",
			Content: SplitText("example of label1 test"),
		},
		ClassItem{
			Class:   "label2",
			Content: SplitText("label2 test example"),
		},
	}

	c := &Classifier{}

	c.Train(d)

	m := c.GetMatches(SplitText("label1 example test"))

	if len(m) != 2 {
		t.FailNow()
	}

	res := m[0]

	if m[0].Probability < m[1].Probability {
		res = m[1]
	}

	if res.Class != "label1" {
		t.Fail()
	}
}
