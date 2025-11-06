// struct 可以指定資料類型
// interface 定義多個函式的回傳值
package main

import (
	"fmt"
)

// 指定參數類型為 uint8
type gasEngine struct {
	mpg    uint8
	gallon uint8
}

type eletricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// e 為 gasEngine 指定類型 ， miles 為uint8
func (e gasEngine) milesLeft() uint8 {
	return e.gallon * e.mpg
}

func (e eletricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh

}

type engine interface {
	milesLeft() uint8
}

func canMake(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there !!")
		fmt.Println()
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {
	// 從 外呼叫gasEngine
	var mygasEngine gasEngine = gasEngine{25, 15}
	fmt.Println(mygasEngine.mpg, mygasEngine.gallon)

	var mygasEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(mygasEngine2)
	fmt.Println(mygasEngine2.gallons)
	fmt.Println(mygasEngine2.mpg)

	//interface 表示
	var myEngine3 eletricEngine = eletricEngine{25, 15}
	canMake(myEngine3, 50)

}
