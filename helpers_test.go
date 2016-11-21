package nbc

import (
	"testing"
)

func TestSplitText(t *testing.T) {
	t1 := SplitText("test foo")
	if len(t1) != 2 || t1[0] != "test" || t1[1] != "foo" {
		t.Fail()
	}
}
