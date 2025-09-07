package utils_pointer

func ValueOrDefault[T any](ptr *T, def T) T {
	if ptr != nil {
		return *ptr
	}

	return def
}
