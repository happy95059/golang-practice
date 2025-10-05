package main

import (
	"fmt"
	"go_practice/exercises"
)

func main() {
	// ## 1. Hello, World!
	fmt.Println("## 1. Hello, World!")
	fmt.Println()

	// *****實現*****
	exercises.HelloWorld()
	// *****實現*****

	// ## 2. 變數、常數、型別宣告
	fmt.Println("\n## 2. 變數、常數、型別宣告")
	fmt.Println()

	// *****實現*****
	exercises.DefineVar()
	// *****實現*****

	// ## 3. for 迴圈 + if 判斷
	fmt.Println("\n## 3. for 迴圈 + if 判斷")
	fmt.Println()

	// *****實現*****
	exercises.TryForLoop()
	// *****實現*****

	// ## 4. slice（動態陣列）基本操作
	fmt.Println("\n## 4. slice（動態陣列）基本操作")
	fmt.Println()

	// *****實現*****
	exercises.FindMaxNum()
	// *****實現*****

	// ## 5. map（字典）應用
	fmt.Println("\n## 5. map（字典）應用")
	fmt.Println()

	// *****實現*****
	exercises.UseMap()
	// *****實現*****

	// ## 6. function 定義與呼叫
	fmt.Println("\n## 6. function 定義與呼叫")
	fmt.Println()

	// *****實現*****
	dailyMoney := []int{1, 3, 5, 7, 1}
	totalMoney1 := exercises.SumSlice1(dailyMoney)
	totalMoney2 := exercises.SumSlice2(dailyMoney)
	totalMoney3 := exercises.SumSlice3(dailyMoney)
	fmt.Println(totalMoney1)
	fmt.Println(totalMoney2)
	fmt.Println(totalMoney3)
	// *****實現*****

	// ## 7. struct（結構體）基本用法
	fmt.Println("\n## 7. struct（結構體）基本用法")
	fmt.Println()

	// *****實現*****
	mitch := exercises.CreatePerson("mitch hsu", 25)
	fmt.Println(mitch.Name)
	fmt.Println(mitch.Age)
	// *****實現*****

	// ## 8. error 返回基礎
	fmt.Println("\n## 8. error 返回基礎")
	fmt.Println()

	// *****實現*****
	intA := 7
	intB := 0
	result, err := exercises.DivideInt(intA, intB)
	if err != nil {
		fmt.Println("炸裂之王", err)
	} else {
		fmt.Println(result)
	}
	// *****實現*****

}
