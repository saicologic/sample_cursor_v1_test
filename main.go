package main

import "fmt"

// 足し算する
func add(a int, b int) int {
	return a + b
}

// 3つの数字を足し算する
func addThree(a int, b int, c int) int {
	return a + b + c
}

// 4つの数字を足し算する
func addFour(a int, b int, c int, d int) int {
	return a + b + c + d
}

// 引き算する
func subtract(a int, b int) int {
	return a - b
}

// 掛け算する
func multiply(a int, b int) int {
	return a * b
}

func main() {
	var x, y, z, w int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)
	fmt.Scan(&w)

	sum := add(x, y)
	sumThree := addThree(x, y, z)
	sumFour := addFour(x, y, z, w)
	difference := subtract(x, y)
	product := multiply(x, y)

	fmt.Printf("合計: %d\n", sum)
	fmt.Printf("3つの数字の合計: %d\n", sumThree)
	fmt.Printf("4つの数字の合計: %d\n", sumFour)
	fmt.Printf("差: %d\n", difference)
	fmt.Printf("積: %d\n", product)
}
