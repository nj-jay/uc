package uc

import (

	"testing"
)

func TestAdd(t *testing.T) {



	if Add(1, 2) != 3 {

		t.Error("Add(1, 2) must be 3")
	}
}


