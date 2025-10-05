package exercises

// ## 6. function 定義與呼叫

func SumSlice1(amounts []int) int {
	temp := 0
	for i := 0; i < len(amounts); i++ {
		temp += amounts[i]
	}
	return temp
}

func SumSlice2(amounts []int) int {
	temp := 0
	for i := range amounts {
		temp += amounts[i]
	}
	return temp
}
func SumSlice3(amounts []int) int {
	temp := 0
	for _, v := range amounts {
		temp += v
	}
	return temp
}
