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

// 3つの数字を引き算する
func subtractThree(a int, b int, c int) int {
	return a - b - c
}

// 掛け算する
func multiply(a int, b int) int {
	return a * b
}

// 割り算する
func divide(a int, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("0で割ることはできません")
	}
	return float64(a) / float64(b), nil
}

// 3つの数字を割り算する
func divideThree(a int, b int, c int) float64 {
	return float64(a) / float64(b) / float64(c)
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
	quotient, err := divide(x, y)
	quotientThree := divideThree(x, y, z)

	fmt.Printf("合計: %d\n", sum)
	fmt.Printf("3つの数字の合計: %d\n", sumThree)
	fmt.Printf("差: %d\n", difference)
	fmt.Printf("積: %d\n", product)

	if err != nil {
		fmt.Printf("商: エラー - %s\n", err.Error())
	} else {
		fmt.Printf("商: %.2f\n", quotient)
	}

	fmt.Printf("3つの数字の商: %.2f\n", quotientThree)
}
