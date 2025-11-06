package main

import (
	"fmt"
	"strings"
)

func main() {

	// string 讀半字，不消耗記憶體，但index容易錯位
	var long_version_String = "r你esum"

	var indexed1 = long_version_String[0]
	fmt.Printf("%v,%T\n", indexed1, indexed1)

	for i, v := range long_version_String {
		fmt.Println(i, v)
	}

	// rune 讀全字 ，消耗記憶體，但index 不容易錯位
	var myString = []rune("r你esume") //
	var indexed = myString[0]
	fmt.Printf("%v,%T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// 最慢沒效率蒐集字串方式

	var strslice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr1 = ""
	for i := range strslice {
		catStr1 += strslice[i]
	}
	fmt.Printf("\n%v", catStr1)

	// 有效率蒐集字串方式
	var strBuilder strings.Builder
	for i := range strslice {
		strBuilder.WriteString(strslice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)

}
