# Go 新手語法練習題（含提示）

---

## 1. Hello, World!

目標：熟悉 go 程式主結構、印出文字
**提示：**

- Go 程式的 entry point 在 `main()` 函數，需放在 `main` 這個 package 裡
- 印出字串用 `fmt.Println`

---

## 2. 變數、常數、型別宣告

目標：學會宣告變數，型別要明確，可以做基本運算
**提示：**

- 變數宣告語法：`var 變數名 型別 = 值`
- 字串串接用 `+`
- 也可簡寫：`a := 5`（型別自動推斷）

---

## 3. for 迴圈 + if 判斷

目標：體驗 go 的 for 與 if
題目：請寫一個程式，計算 1~100 之間所有「3 的倍數」的總和，並把最後的總和印出來。
**提示：**

- for 迴圈格式：`for i := 1; i <= 100; i++ { ... }`
- if 判斷格式：`if 條件 { ... } else { ... }`

---

## 4. slice（動態陣列）基本操作

目標：建立一個 int slice，找最大值
題目：建立一個 int slice，例如 []int{8, 22, 5, 13, 7}。寫一個 for 迴圈找到最大值，並印出最大值是多少。
**提示：**

- 宣告 slice：`nums := []int{3, 9, 4, 2, 8}`
- 取長度：`len(nums)`
- 迴圈：`for i := 0; i < len(nums); i++ { ... }`
- 指定元素：`nums[i]`

---

## 5. map（字典）應用

目標：使用 map 存 key/value，增刪查
題目：請建立一個 map，key 為字串，value 為 int。加入三組 key/value，印出全部內容，查詢其中一個 key 的值，然後刪除一個 key，再印一次全部內容。

**提示：**

- 宣告 map：`capitals := make(map[string]string)`
- 指定：`capitals["Taiwan"] = "Taipei"`
- 刪除：`delete(capitals, "France")`
- 用 `for k, v := range capitals { ... }` 走訪所有元素

---

## 6. function 定義與呼叫

目標：寫 function 並回傳結果
題目：請寫一個 function，名稱為 SumSlice，參數為 int slice，回傳全部元素的總和。
在 main 裡呼叫這個 function 並印出結果。
**提示：**

- 定義 function：`func max(a int, b int) int { ... }`
- 回傳：`return xxx`

---

## 7. struct（結構體）基本用法

目標：定義 struct，建立實體並取欄位值
題目：定義一個 struct 叫 Person，有兩個欄位：Name (string)、Age (int)。在 main 裡建立一個 Person 實體，指定欄位值，然後把 Name 和 Age 印出來。
**提示：**

- 定義：`type Student struct { Name string; Score int }`
- 建立物件：`s := Student{Name: "Amy", Score: 95}`
- 印：`fmt.Println(s.Name, s.Score)`

---

## 8. error 返回基礎

目標：用 error 回傳異常
題目：請寫一個 function，名稱為 Divide，接收兩個 int（a, b），回傳 (int, error)。
如果 b=0，就回傳 error，不然回傳 a/b。
在 main 裡測試 b=0 跟 b≠0 的情境，正確印出結果或錯誤訊息。
**提示：**

- function 可以 `return 值, err`
- error 用 `errors.New("錯誤訊息")` 產生
- 要 `import "errors"`
