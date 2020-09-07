package main

import (
	"fmt"

	"github.com/nj-jay/uc/cal"
)

func main() {

	var a int32
	var b int64
	a = 5
	b = 4
	fmt.Println(cal.Add(a, a))
	fmt.Println(cal.Add(b, b))
	fmt.Println(cal.Add(1, 2, 3))
	fmt.Println(cal.Add(1, 2))
	fmt.Println(cal.Add(1.1, 2.2))
	fmt.Println(cal.Add(2.22, 3.33))
	fmt.Println(cal.Add("I", " ", "love", " ", "go!"))
	fmt.Println(cal.Add("https://", "nj-jay.com"))
	fmt.Println(cal.Add([]int{1, 2}, []int{3, 4}))
    fmt.Printf("%T\n", cal.Add([]int{5, 6}, []int{7, 8}))

}
