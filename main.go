package main

import "fmt"

// 足し算する
func add(a int, b int) int {
	return a + b
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
	var x, y int
	fmt.Print("1つ目の数字を入力してください: ")
	fmt.Scan(&x)
	fmt.Print("2つ目の数字を入力してください: ")
	fmt.Scan(&y)
	
	sum := add(x, y)
	difference := subtract(x, y)
	product := multiply(x, y)
	
	fmt.Printf("合計: %d\n", sum)
	fmt.Printf("差: %d\n", difference)
	fmt.Printf("積: %d\n", product)
}
