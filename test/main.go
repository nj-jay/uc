package main

import (
	"fmt"

	"github.com/nj-jay/uc"
)

func main() {

	var a int32
	var b int64
	a = 5
	b = 4
	fmt.Println(uc.Add(a, a))
	fmt.Println(uc.Add(b, b))
	fmt.Println(uc.Add(1, 2, 3))
	fmt.Println(uc.Add(1, 2))
	fmt.Println(uc.Add(1.1, 2.2))
	fmt.Println(uc.Add(2.22, 3.33))
	fmt.Println(uc.Add("I", " ", "love", " ", "go!"))
	fmt.Println(uc.Add("https://", "nj-jay.com"))
	fmt.Println(uc.Add([]int{1, 2}, []int{3, 4}))
}
