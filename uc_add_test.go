package uc

import (

	"testing"
)

func TestAdd(t *testing.T) {


	if Add(1, 2) != 3 {

		t.Error("Add(1, 2) must be 3")
	}

    //if Add(1.1, 2,2) != 3.3 {

      //  t.Error("Add(1.1, 2.2) must be 3.3")

    //}
}



