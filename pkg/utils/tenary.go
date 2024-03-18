package utils

func Tenary[T any](v bool, trueValue T, falseValue T) T {
	if v {
		return trueValue
	}
	return falseValue
}
