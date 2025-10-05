package exercises

import (
	"fmt"
)

// ## 3. for 迴圈 + if 判斷
func TryForLoop() {
	var temp int = 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			temp += i
		}
	}

	fmt.Println(temp)
}
