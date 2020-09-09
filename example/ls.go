package main

import (

	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

func main(){

	if len(os.Args) == 1 {

		Ls(".")

	} else {

		path := os.Args[1]

		Ls(path)

	}
}

func Ls(path string) {

	info, err := ioutil.ReadDir(path)

	if err != nil {

		fmt.Println("no this dictionary")

		os.Exit(0)
	}

    for _, f := range info {

        length := f.Size()

        if ok := f.IsDir(); ok {

        	color.Blue("%s", f.Name())

		} else {

			if length > 1024 * 1024 {

				fmt.Println(f.Name(), " ", fmt.Sprintf("%.2f", float64(length)/(1024*1024)), "Mb")

			} else {

				fmt.Println(f.Name(), " ", fmt.Sprintf("%.1f", float64(length)/1024), "Kb")
			}
		}

		}
}
