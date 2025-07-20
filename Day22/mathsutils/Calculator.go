package mathsutils

type Number interface{
	int | float32 | int64
}

func Add[T Number](a, b T) T {
	return a+b
}
func Subtract[T Number](a, b T) T {
	return a - b
}
func Multiply[T Number](a, b T) T {
	return a * b
}
func Divide[T Number](a, b T) T {
	return a / b
}