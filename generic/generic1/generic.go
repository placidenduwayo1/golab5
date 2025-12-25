package generic1

type Number interface {
	byte | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64
}

func ComputeSum[T Number](slice []T) T {
	var result T = 0 // result is of type T
	for _, v := range slice {
		result = result + v
	}

	return result
}
