package util

func Factorial(target int) int {
	if target == 1 {
		return target
	}
	return target * Factorial(target - 1)
}
