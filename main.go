package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// concurrency 兩個任務同一個cpu  parrellel executaion 處理核心分開
var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"1d1", "1d2", "1d3", "id4", "id5"}
var results = []string{}

// 使用迴圈呼叫五個函式進行完成
func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// 啟動執行序
		wg.Add(1)
		// 呼叫dbcail函式執行
		go dbCail(i)
	}
	// 等待所有執行任務完成以後，輸出完成時間
	wg.Wait()
	fmt.Printf("\nTotal executation time:%v", time.Since(t0))

}

// 呼叫資料庫 ID
func dbCail(i int) {
	var delay float32 = rand.Float32() * 5000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("the result from the database:", dbData[i])

	// 使用資料上鎖避免資料再回傳的時候避免被其他已完成函式干擾結果
	save()

	// 完成任務
	wg.Done()
}

// 儲存資料
func save(i int) {
	m.Lock()
	// 完成資料連接才能留給下一個使用著
	results = append(results, dbData[i])
	m.Unlock()

}

// 列印出儲存結果
func log() {
	fmt.Printf("\nThe result are %v", results)

}
