package utils

import "github.com/mylxsw/go-utils/array"

// If 三元运算符
// a, b := 2, 3
// max := If(a > b, a, b).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func ArrayIn() {
	array.In(1, []int{23})
}
