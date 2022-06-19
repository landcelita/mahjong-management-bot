package testutil

func First[T any] (val *T, e error) T {
	return *val
}