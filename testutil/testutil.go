package testutil

func First[T any] (val *T, e error) T {
	if(e != nil) {
		panic(e)
	}
	
	return *val
}