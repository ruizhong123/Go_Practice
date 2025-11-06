package main

import "fmt"

func main() {
	var p *int32 = new(int32) // 可改
	var i int32               // 不可改
	fmt.Printf("\nThe value p point to is: %v", *p)
	fmt.Printf("\nThe value i is %v", i)

	p = &i // p 表示 i 的記憶體位置
	*p = 1 // 定義為記憶體位置上面的值
	// 可以發現到在相同記憶位置上 i、p 的值是一致
	fmt.Printf("\nThe value p point to is :%v", *p)
	fmt.Printf("\nThe value i is %v", i)

	// 從陣列上發現如果slice跟sliceCopy相同，則位置跟值一致，
	// 如果sliceCopy改動的話 slice 跟著改動

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

}
