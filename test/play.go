package main


import (

    "fmt"
    "github.com/nj-jay/uc"

)

func main(){

    err := uc.Play("./七里香.mp3")
    if err != nil {
        fmt.Println(err)
    }
}

