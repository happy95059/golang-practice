package exercises

import "errors"

func DivideInt(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide zero")
	}
	return a / b, nil

}
