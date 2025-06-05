package main

import "fmt"

func add(a int, b int) int {
	return a - b
}

func main() {
	var x, y int
	fmt.Print("1つ目の数字を入力してください: ")
	fmt.Scan(&x)
	fmt.Print("2つ目の数字を入力してください: ")
	fmt.Scan(&y)
	sum := add(x, y)
	fmt.Printf("合計: %d\n", sum)
}
