package uc

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	var a float64

	var b int32

	var c int32

	var d int64

	var e int64

	a = 1.1

	b = 3

	c = 6

	d = 2

	e = 4

	sliceInt := []int{1, 2, 3, 4}

	if reflect.DeepEqual(Add([]int{1, 2}, []int{3, 4}), sliceInt) != true {

		t.Error("must be []int{1, 2, 3, 4}")

	}

	if Add(b, b) != c {

		t.Error("Add(3, 3) must be 6")
	}

	if Add(d, d) != e {

		t.Error("Add(2, 2) must be 4")
	}

	if Add(a, a) != 2.2 {

		t.Error("Add(1.1, 1.1) must be 2.2")
	}

	if Add("hello", "world") != "helloworld" {

		t.Error("must be helloworld")
	}

}
