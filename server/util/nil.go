package util

func IfNil[T any](i *T, a *T) *T {
	if i == nil {
		return a
	}

	return i
}
