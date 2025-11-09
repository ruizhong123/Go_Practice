# Go_Practice

## API基本架構
&nbsp;&nbsp; GO 語言在多併發式任務執行 concurrency 以及各平台的資料傳送效率上超過node.js以及python語言並且可以跟C語言媲美，且語法相對而言必較簡單這也為什麼GO語言在開發平台的後端無論是網站開發擔任後端的重要角色。這也為 API 在GO 語言中扮演相當重要的角色。而這次主要要來建立餘額登錄系統來測試來做為api的範本，而這次的處理架構分別是主要分為以下四類模組:
___________________
|________|________
| - api 
|

### - API folder 
&nbsp;&nbsp;&nbsp;&nbsp;API 基本上並須要有資料名稱以及類型，這樣會比較好確認資料的來源以及所給予的資訊，而做的資料主要主要分為兩類的資料，一個是餘額參數以及餘額傳過來的授權代碼，這類是屬於使用著的資訊，而另一個相對重要的資訊是叫做錯誤回傳的資訊，主要是能夠在發生錯誤時明確的知道。
 
&nbsp;&nbsp;&nbsp;&nbsp; 這次範本示範的有四種類型的資料CoinBalanceParam(使用著名稱)
CoinBalanceResponse(使用著餘額)以及Error(錯誤回傳訊息格式)，主要是由type來去訂義資料，
而其中Errror回傳的資訊主要有分兩大主要的資訊請求錯誤回傳另一個是網路消失時會回傳的錯誤，而我
們把這兩類資料定義成Errorwriter帶入函式的變數，而這變數會被middle中的authorization.go文件
作為模組使用。

### - middle folder 
&nbsp;&nbsp;&nbsp;&nbsp; middle







  

