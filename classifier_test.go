package nbc

import (
	"testing"
)

func TestNewClassifier(t *testing.T) {
	c := NewClassifier([]string{"label1"})

	if len(c.labels) != 1 {
		t.Error("labels not set")
	}

}
