package exercises

import (
	"fmt"
)

// ## 5. map（字典）應用
func UseMap() {
	// 宣告
	nameMoneyMap := make(map[string]int)

	//賦值
	nameMoneyMap["mitch"] = 200
	nameMoneyMap["andy"] = 300
	nameMoneyMap["emma"] = 400
	for k, v := range nameMoneyMap {
		fmt.Printf("key = %s, value = %d\n", k, v)
	}
	delete(nameMoneyMap, "mitch")
	fmt.Println()
	for k, v := range nameMoneyMap {
		fmt.Printf("key = %s, value = %d\n", k, v)
	}
}
