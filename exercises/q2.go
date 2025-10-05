package exercises

import (
	"fmt"
)

// ## 2. 變數、常數、型別宣告
func DefineVar() {
	test_qucik := 3.14
	var test_int int = 3
	var test_str string = "33"
	var test_str2 string = "12"
	test_auto := map[string]int{
		test_str: 22,
		"qq":     11,
	}

	fmt.Println(test_qucik)
	fmt.Println(test_int)
	fmt.Println(test_str)
	fmt.Println(test_auto)

	fmt.Printf("type: %T\n", test_int)
	fmt.Printf("type: %T\n", test_str)
	fmt.Printf("type: %T\n", test_auto)
	fmt.Println(test_str + test_str2)
}
