package main


import (

    "fmt"
    "github.com/nj-jay/uc/play"

)

func main(){

    err := play.PlayAudio("七里香-周杰伦.mp3")

    if err != nil {

        fmt.Println(err)
    }

}

