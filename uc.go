package uc


func Add(a ...interface{}) interface{} {

    var sumInt int

    var sumFloat64 float64

    for index, v := range a {

        switch v.(type) {

        case int:

            if _, ok := v.(int); ok {

                sumInt += v.(int)

                if len(a) == index+1 {

                    return sumInt
                }
            }

        case float64:

            if _, ok := v.(float64); ok {

                sumFloat64 += v.(float64)

                if len(a) == index +1 {

                    return sumFloat64
                }
            }

        default:

            return nil

        }
    }
    return 0
}

