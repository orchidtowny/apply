package util

func IsNil[T any](i *T) bool {
	return i == nil
}

func IfNil[T any](i *T, a *T) *T {
	if i == nil {
		return a
	}

	return i
}
