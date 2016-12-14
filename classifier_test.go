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

func TestNormalizeTrainedData(t *testing.T) {
	c := &Classifier{}

	c.Train([]ClassItem{
		ClassItem{
			Class:   "test",
			Content: SplitText("test example 1"),
		},
		ClassItem{
			Class:   "test",
			Content: SplitText("foo example bar"),
		},
	})

	if c.trainingSet["test"]["test"] != 0.5 {
		t.Fail()
	}

	if c.trainingSet["test"]["example"] != 1 {
		t.Fail()
	}

}

func TestClampMissingToMinWeight(t *testing.T) {
	c := &Classifier{}

	c.Train([]ClassItem{
		ClassItem{
			Class:   "foo",
			Content: SplitText("foo"),
		},
		ClassItem{
			Class:   "bar",
			Content: SplitText("bar"),
		},
	})

	if c.trainingSet["foo"]["bar"] != c.minWeight {
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
	a := c.Validate(d, 1)

	if a[0] < 1 {
		t.Fail()
	}
}

func TestFullProbability(t *testing.T) {
	d := []ClassItem{
		ClassItem{
			Class:   "l1",
			Content: SplitText("1 2 3 4 5"),
		},
	}

	c := &Classifier{}

	c.Train(d)

	m := c.PredictAll(SplitText("2 3 4"))

	if m[0].Probability != 1 {
		t.Fail()
	}
}

func TestTrainWithContentOverlapping(t *testing.T) {
	d := []ClassItem{
		ClassItem{
			Class:   "label2",
			Content: SplitText("label2 test example"),
		},
		ClassItem{
			Class:   "label1",
			Content: SplitText("example of label1 test"),
		},
	}

	c := &Classifier{}

	c.Train(d)

	m := c.PredictAll(SplitText("label1 example test"))

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

func TestPerformanceOfTrainingSetIsPerfect(t *testing.T) {
	l1 := "test text 1"
	l2 := "foo bar"

	d := []ClassItem{
		ClassItem{
			Class:   "1",
			Content: SplitText(l1),
		},
		ClassItem{
			Class:   "2",
			Content: SplitText(l2),
		},
	}

	c := &Classifier{}

	c.Train(d)

	a := c.Accuracy(d)

	if a < 1 {
		t.Fail()
	}
}

func TestPerformanceOf50Percent(t *testing.T) {
	l1 := "test text 1"
	l2 := "foo bar"

	d := []ClassItem{
		ClassItem{
			Class:   "1",
			Content: SplitText(l1),
		},
		ClassItem{
			Class:   "2",
			Content: SplitText(l2),
		},
	}

	c := &Classifier{}

	c.Train(d)

	acc := c.Validate([]ClassItem{
		ClassItem{
			Class:   "1",
			Content: SplitText(l1),
		},
		ClassItem{
			Class:   "1",
			Content: SplitText(l2),
		},
	}, 1)

	if acc[0] != 0.5 {
		t.Fail()
	}
}

func TestAccuracyByZeroItemsGetsPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fail()
		}
	}()

	c := &Classifier{}
	c.Accuracy([]ClassItem{})

}
