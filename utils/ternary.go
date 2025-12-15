package utils

func Ternary(condition bool, a, b any) any {
	if condition {
		return a
	}

	return b
}
