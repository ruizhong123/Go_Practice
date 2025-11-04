package datatype // 定義式包裝成套件包

import "fmt"
import "unicode/utf8"

func Demo() {
	// 命名變數 : var + 變數名 + 變數類型 = 指定數
	var intNum int16 = 32758
	intNum = intNum + 1
	fmt.Println(intNum)

	// 小數名稱命名
	var floatNum float64 = 12345.9000
	fmt.Println(floatNum)

	//如果數字相加，確保數字型態相同之後再相加
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// 字串 : 跟python不同的地方是 go 字串 遵守 ascii25
	var mystring1 string = "a"
	fmt.Println(mystring1)
	var mtstring2 string = "Hello" + " " + "World"
	fmt.Println(mtstring2)

	// 字串長度
	fmt.Println(len("b"))                       // => 11010111
	fmt.Println(utf8.RuneCountInString("test")) // => 4

	//布林
	var myBoolean bool = false
	fmt.Println(myBoolean)

	// 命名變數(簡化版)
	myVar := "text"
	fmt.Println(myVar)

	//同時命名兩個變數
	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// 如果想要隨時改變變數名稱使用const
	const myconst string = "the sting"
	fmt.Println(myconst)

}
