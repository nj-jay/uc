package uc

import (

    "fmt"

    "strconv"
)

var (

	sumInt int

    sumFloat64 float64
    
    sumString string
    
)

func Add(a ...interface{}) interface{} {

    for index, v := range a {

        switch v.(type) {

        case int:

            if _, ok := v.(int); ok {

                sumInt += v.(int)

                if len(a) == index+1 {

                    return sumInt
                }
            }

        case float64 :

            if _, ok := v.(float64); ok {

                sumFloat64 += v.(float64)

                if len(a) == index + 1 {

                    return parseFloat(sumFloat64)
                    
                }
            }
	
		case string :

			if _, ok := v.(string); ok {
			
				sumString += v.(string)
				
				if len(a) == index + 1 {
				
					return sumString
					
				}
			}

        default:

            return nil

        }
    }
    
    return 0
}

func parseFloat(f float64) float64 {

    str := fmt.Sprintf("%.1f", f)

    sumFloat64, _ := strconv.ParseFloat(str, 64)

	return sumFloat64

}
