package main

import (
	"fmt"
	"os"
)

func main() {
	// make directory
	if err := os.Mkdir("hoge", 0777); err! = nil {
		fmt.Println(err)
	}

	// make sub directory
	if err := os.MkdirAll("hoge/fuga", 0777); err! = nil {
		fmt.Println(err)
	}
}