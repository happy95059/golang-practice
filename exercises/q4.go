package exercises

import (
	"fmt"
)

// ## 4. slice（動態陣列）基本操作
func FindMaxNum() {
	nums := []int{3, 9, 4, 2, 8}
	temp := 0
	for i := 0; i < len(nums); i++ {
		if temp < nums[i] {
			temp = nums[i]
		}
	}
	fmt.Println(temp)
}
