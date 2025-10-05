# 🐍 → 🦫 給 Python 開發者的 Go 入門導覽

這份筆記主要讓你快速理解：

> 「Go 跟 Python 到底哪裡一樣、哪裡不同」
> 「為什麼 Go 有這樣的語法設計」

---

## 🚀 1. Go 的檔案結構與進入點

### Python

```python
# main.py
print("Hello world")
```

### Go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

🔹 Go 的「執行起點」永遠是：

- `package main`
- 裡面的 `func main()`

這是 Go 約定俗成的 entry point。
一般慣例會放在：

```
cmd/
 └── main.go
```

執行方式：

```bash
go run cmd/main.go
```

---

## 🧱 2. package 與 import 的概念

### Python

```python
import os
import mylib.utils
```

### Go

```go
import (
    "fmt"
    "myproject/utils"
)
```

- `package` 是 Go 的模組單位，類似 Python 的「module / package」。
- 每個 `.go` 檔都要宣告它屬於哪個 `package`。
- 相同資料夾下的檔案，**package 名稱必須一致**。
- 其他資料夾的函式，要 `import` 才能使用。

---

## 🧩 3. 專案初始化：`go mod`

Python 用 `venv + requirements.txt`
Go 用 `go mod` 管理依賴。

### 初始化模組

```bash
go mod init myproject
```

這會建立一個 `go.mod`：

```go
module myproject
go 1.22
```

### 執行程式

```bash
go run cmd/main.go
```

這兩行幾乎是所有 Go 專案的起點。

---

## 🧮 4. 變數與型別

### Python

```python
a = 10
b = "hi"
```

### Go

```go
var a int = 10
b := "hi" // 自動推斷型別
```

Go 是靜態型別（compile time 決定型別），
所以你不能在執行中改變變數的型別。

---

## 🧺 5. slice 是什麼？（動態陣列）

### Python

```python
nums = [1, 2, 3]
nums.append(4)
```

### Go

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)
```

🔹 slice 是 Go 裡「動態大小的陣列」

- 可以 `append`
- 可以 `len(nums)`
- 可以用 `for range` 迴圈

---

## 🗺️ 6. map（字典）

### Python

```python
capitals = {"Taiwan": "Taipei", "Japan": "Tokyo"}
```

### Go

```go
capitals := map[string]string{
    "Taiwan": "Taipei",
    "Japan":  "Tokyo",
}
```

🔹 用法幾乎跟 Python dict 一樣
🔹 刪除用 `delete(capitals, "Japan")`

---

## ⚙️ 7. 函式宣告與回傳值

### Python

```python
def add(a, b):
    return a + b
```

### Go

```go
func add(a int, b int) int {
    return a + b
}
```

Go 的函式宣告順序是：

```
func 名稱(參數型別) 回傳型別
```

---

## 📦 8. 結構體（struct）取代 class

Go 沒有 class，但有 struct。

### Python

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
```

### Go

```go
type Person struct {
    Name string
    Age  int
}
```

建立：

```go
p := Person{Name: "Amy", Age: 20}
fmt.Println(p.Name)
```

---

## 🧨 9. 錯誤處理（error）

Go 沒有 try/except，用「回傳值 + error」處理。

### Python

```python
def divide(a, b):
    if b == 0:
        raise ValueError("division by zero")
    return a / b
```

### Go

```go
import "errors"

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

---

## 🧭 10. 為什麼很多 Go 專案都有 `cmd/`

慣例設計：

```
cmd/
 ├── api/       ← HTTP 伺服器入口
 └── worker/    ← 背景任務入口
internal/
 └── service/
```

- 每個 `cmd/` 底下是一個可執行程式入口（有 `main()`）。
- `internal/` 放邏輯程式碼，外部不能 import。

---

## 🧰 11. Go 的哲學

- **明確比魔法好**（沒有隱式轉換）
- **每個錯誤都要處理**
- **程式結構清楚勝過炫技**
- **簡潔比靈活重要**

---

## 🧪 快速試跑區

你可以直接試：

```bash
go mod init hello-go
echo 'package main; import "fmt"; func main(){ fmt.Println("Hi!") }' > main.go
go run .
```

---

## 🧠 小結

| Python 概念   | Go 對應                  |
| ------------- | ------------------------ |
| list          | slice                    |
| dict          | map                      |
| class         | struct + methods         |
| import module | import package           |
| virtualenv    | go.mod                   |
| try/except    | return error             |
| main.py       | package main + func main |

---

📘 建議檔名：
`go_for_python.md` 或 `go_guide_for_python.md`
