package main

import "fmt"

func main() {
	// declare int
	var i int = 1234

	// int => int32
	var u uint32 = uint32(i)

	// uint => float32
	var f float32 = float32(u)

	// int => string
	var s string = string(i)

	// string型から配列（スライス）型に変換
	var b []byte = []byte("abc")

	// output
	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)
}
