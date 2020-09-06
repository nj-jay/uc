package main


import (

    "fmt"
    "github.com/nj-jay/uc"

)

func main(){

    err := uc.PlayAudio("以父之名-周杰伦.mp3")

    if err != nil {

        fmt.Println(err)
    }

}

