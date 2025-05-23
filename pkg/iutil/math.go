package iutil

// CeilInts 对整数除法结果向上取整
func CeilInts[T int | int32 | int64 | uint | uint32 | uint64](a, b T) T {
	return (a + b - 1) / b
}
