package main

import (

	"fmt"

	"strconv"
)

func main() {

	fmt.Println(Sum())

}

func Sum() float64 {

	sum := 1.1 + 2.22

	var str string

	var sumFloat64 float64

	str = fmt.Sprintf("%.2f", sum)

	sumFloat64, _ = strconv.ParseFloat(str, 64)

	return sumFloat64
}
