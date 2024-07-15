package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dir, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range dir {
		fmt.Println(file.Name())
	}
}
