package main

import (
	"error"
	"fmt"
)

func main() {
	// 從其他定義是呼叫到 main 主程式
	var printValue string = "Hello World"
	printMe(printValue)

	// 從 main 召回 isDivision
	var numerator int = 11
	var denomintor int = 2
	var result, remainder, err = inDivision(numerator, denomintor)

	// if ... else ... 邏輯式簡寫
	// go 的 && 代表 and , || 代表 or

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of theinteger division is %v with remainder %v", result, remainder)
	}

	// 使用 Printf 以 result remainder 入 %v
	fmt.Printf("the result of the integer division is %v with reminder %v", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func inDivision(numerator int, denomintor int) (int, int) {
	var result int = numerator / denomintor
	var remainder int = numerator % denomintor

	return result, remainder
}
