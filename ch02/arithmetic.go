package main

import "fmt"

func main() {
	fmt.Println("1 + 2 =", 1+2)

	fmt.Println("\"abc\" + \"XYZ\" = ", "abc"+"XYZ")

	fmt.Println("3 - 2=", 3-2)

	fmt.Println("2 * 3", 2*3)

	fmt.Println("-5 / 2", -5/2)

	fmt.Println("5 % 2 =", 5%2)

	fmt.Println("3 & 6 =", 3&6)

	fmt.Println("3 | 6 =", 3|6)

	// ビットクリア(and not)
	fmt.Println("3 &^ 6 =", 3&^6)

	// 左シフト演算
	fmt.Println("2 << 1 =", 2<<1)

	// 右シフト演算
	fmt.Println("2 >> 1 =", 2>>1)
}
