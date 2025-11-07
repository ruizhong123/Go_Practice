package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE = 5

func main() {
	var chickenchannel = make(chan string)
	var tofuchannel = make(chan string)
	var website = []string{"walmart.com", "costco.com", "wholesfood.com"}

	for i := range website {
		go checkchickencnannel(website[i], chickenchannel)
		go checkTofuchannel(website[i], tofuchannel)

	}

	sendMessage(chickenchannel, tofuchannel)

}

func checkTofuchannel(website string, Tofuchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			Tofuchannel <- website
			break
		}
	}
}

func checkchickencnannel(website string, chickenchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenchannel <- website
			break
		}
	}
}

func sendMessage(chickenchannel chan string, Tofuchannel chan string) {
	select {
	case website := <-chickenchannel:
		fmt.Printf("Found deal on chicken at %v", website)
	case website := <-Tofuchannel:
		fmt.Printf("Found deal on chicken at %v", website)

	}

}

// func main() {

// 	var c = make(chan int)
// 	go process(c)

// 	// 從主線道發送資料
// 	for i := range c {
// 		fmt.Printf("\nthe value have been exited: %v", i)
// 	}

// }

// // 使用 process 讓channel 接收資料
// func process(c chan int) {
// 	// 當管道執行完接收資料資料時就立即關閉
// 	defer close(c)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// }
