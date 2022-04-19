package utils

// If 三元运算符
// a, b := 2, 3
// max := If(a > b, a, b).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
