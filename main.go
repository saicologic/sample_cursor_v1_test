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

// 引き算する
func subtract(a int, b int) int {
	return a - b
}

// 掛け算する
func multiply(a int, b int) int {
	return a * b
}

func main() {
	var x, y, z int
	fmt.Print("1つ目の数字を入力してください: ")
	fmt.Scan(&x)
	fmt.Print("2つ目の数字を入力してください: ")
	fmt.Scan(&y)
	fmt.Print("3つ目の数字を入力してください: ")
	fmt.Scan(&z)

	sum := add(x, y)
	sumThree := addThree(x, y, z)
	difference := subtract(x, y)
	product := multiply(x, y)

	fmt.Printf("2つの数字の合計: %d\n", sum)
	fmt.Printf("3つの数字の合計: %d\n", sumThree)
	fmt.Printf("差: %d\n", difference)
	fmt.Printf("積: %d\n", product)
}
