package main

import (
    "github.com/nj-jay/uc/shell"
    "os"
)

func main(){
    
    url := os.Args[1]
    shell.Wget(url)
}

