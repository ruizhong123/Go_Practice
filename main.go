package main

import (
	"fmt"
)

type gasEngine struct {
	gallons float32
	mpg     float32
}

type elerticEngine struct {
	kwh float32
	mph float32
}

type car[T gasEngine | elerticEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {

	var gascar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     23,
		},
	}

	var Eletriccar = car[elerticEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: elerticEngine{
			kwh: 12.3,
			mph: 22,
		},
	}

	fmt.Println(gascar)
	fmt.Println(Eletriccar)
}

// generic 是指函式允許不同類型的資料被單一個函式計算

// func main() {

// 	// 任何類型都能跑
// 	Print(42)      // T = int
// 	Print("hello") // T = string
// 	Print(3.14)    // T = float64

// }

// func Print[T any](value T) {
// 	fmt.Println(value)
// }
