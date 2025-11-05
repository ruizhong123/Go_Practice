package main

import "fmt"

func main() {
	// // 限制矩陣使用整數，並且可容納三位數
	// var intArr [3]int32

	// // 矩陣第一位數是123
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// // 矩陣命為[1,2,3]
	intArr := [3]int32{1, 2, 3}
	fmt.Println(intArr)

	// 矩陣[1,2,3] 另一寫法
	another_intArr := [...]int32{1, 2, 3}
	fmt.Println(another_intArr)

	// slice 是底層矩陣最大可容許的資料長度
	// cap 為大容納資料長度
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// intSlice3 :使用make使陣列資料長度為 3 容量為 8
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v", len(intSlice3), cap(intSlice3))

	// 使用 map 建造 key鍵 以及 value 鍵
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) // 默認jason 初始值 0

	// age 回傳預設年齡 ok 回傳 false
	var age, ok = myMap2["Jason"]
	fmt.Println(age, ok)

	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invaild Value")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name:%v,Value:%v ", name, age)
	}

	// 從列表回傳 index 跟值
	for i, v := range intArr {
		fmt.Printf("Index:%v,Value:%v", i, v)
	}

	// 迴圈累加1的1不同寫法
	var i int = 0 // 變數命名為 0

	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i = i + 1
	}

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

}
