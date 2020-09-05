package uc

import (

	"fmt"

	"strconv"
)

func Add(a ...interface{}) interface{} {

	var (

		sumInt int = 0

		sumInt32 int32 = 0.0

		sumInt64 int64 = 0

		sumFloat64 float64 = 0.0

		sumString string = ""

		sumSliceInt []int
	)

	for index, v := range a {

		switch v.(type) {

		case int:

			if _, ok := v.(int); ok {

				sumInt += v.(int)

				if len(a) == index+1 {

					return sumInt
				}
			}

		case int64:

			if _, ok := v.(int64); ok {

				sumInt64 += v.(int64)

				if len(a) == index+1 {

					return sumInt64

				}
			}

		case int32:

			if _, ok := v.(int32); ok {

				sumInt32 += v.(int32)

				if len(a) == index+1 {

					return sumInt32

				}
			}

		case float64:

			if _, ok := v.(float64); ok {

				sumFloat64 += v.(float64)

				if len(a) == index+1 {

					return parseFloat64(sumFloat64)

				}
			}

		case string:

			if _, ok := v.(string); ok {

				sumString += v.(string)

				if len(a) == index + 1 {

					return sumString

				}
			}

		case []int:

			if _, ok := v.([]int); ok {

				sumSliceInt = append(sumSliceInt, v.([]int)...)

				if len(a) == index + 1 {

					return sumSliceInt
				}

			}

		default:

			return nil
		}
	}

	return 0
}

func parseFloat64(f float64) float64 {

	str := fmt.Sprintf("%.2f", f)

	sumFloat, _ := strconv.ParseFloat(str, 64)

	return sumFloat

}
