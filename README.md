# nbc

[![Build Status](https://travis-ci.org/SamuelKupferschmid/nbc.svg?branch=master)](https://travis-ci.org/SamuelKupferschmid/nbc)
[![Coverage Status](https://coveralls.io/repos/github/SamuelKupferschmid/nbc/badge.svg)](https://coveralls.io/github/SamuelKupferschmid/nbc)


**this Code is still under construction and not ready for use!**

Implementation of a simple naive bayes classifier.

This Library was developed for education purpose and was not used yet in production by the author.

This classifier could be used to detect if a text is spam or not or to detect a language of a given text.

For this a large training and validationset is required.

 
## usage

```go
d := []LabelItem{
	LabelItem{
		label:   "test1",
		content: SplitText("example text"),
	},
}

c := &Classifier{}

c.Train(d)
p := c.Validate(d)

if p.Precision < 1 || p.Recall < 1 {
	t.Fail()
}
```