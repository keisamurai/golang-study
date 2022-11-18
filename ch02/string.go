package main

import "fmt"

func main() {
	var str string
	str = "fafafa"
	str = str + "faffafadf"
	str += "X"
	fmt.Println(str)
}
