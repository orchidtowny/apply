package util

func MustNotError(err error) {
	if err != nil {
		panic(err)
	}

	return
}

func MustReturn[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}
