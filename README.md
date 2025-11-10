# Go_Practice

## API基本架構
&nbsp;&nbsp; GO 語言在goroutineu以及channel來實現輕量級併發，在api json的表現超越python、node.js甚至速度媲美
C語言，所以GO語言在API中是被重視的語言之一


### - API module
&nbsp;&nbsp;&nbsp;&nbsp;API 基本上並須要有資料名稱以及類型，這樣會比較好確認資料的來源以及所給予的資訊，而做的資料主要主要分為兩類的資料，一個是餘額參數以及餘額傳過來的授權代碼，這類是屬於使用著的資訊，而另一個相對重要的資訊是叫做錯誤回傳的資訊，主要是能夠在發生錯誤時明確的知道。
 
&nbsp;&nbsp;&nbsp;&nbsp; 這次範本示範的有四種類型的資料CoinBalanceParam(使用著名稱)
CoinBalanceResponse(使用著餘額)以及Error(錯誤回傳訊息格式)，主要是由type來去訂義資料，
而其中Errror回傳的資訊主要有分兩大主要的資訊請求錯誤回傳另一個是網路消失時會回傳的錯誤，而我
們把這兩類資料定義成Errorwriter帶入函式的變數，而這變數會被middle中的authorization.go文件
作為模組使用。

### - internal module 
&nbsp;&nbsp;&nbsp;&nbsp; 主要掌控api連線資訊，齊文件有handler(連接網路)、middleware(中介軟體)以及tool(資料庫套件
)，其中我們會使用handler做為主要連線資訊的根據並以go語言中chi.Mux的 Route、Use以及Get 對餘額帳戶進行連線設定、授權使用以及資訊獲取，而其中控制的文件有middleware(中介軟體)的authorization授權函式以及get_coin_balance，我們使用這裡的授權軟體，如果通過的話得到getcoinbalance的資料。

### - cmd 執行命令

 









  

