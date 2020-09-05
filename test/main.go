package main

import (

	"fmt"

	"github.com/nj-jay/uc"

)

func main(){

    fmt.Println(uc.Add(1, 2, 3))
    fmt.Println(uc.Add(1, 2))
	fmt.Println(uc.Add(1.1, 2.2))
    fmt.Println(uc.Add("I", " ", "love", " ", "go!"))
    fmt.Println(uc.Add([]int{1, 2}, []int{3, 4}))
}
