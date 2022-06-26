package testutil

func FirstPtoV[T any] (val *T, e error) T {
	if(e != nil) {
		panic(e)
	}
	
	return *val
}

func FirstPtoP[T any] (val *T, e error) *T {
	if(e != nil) {
		panic(e)
	}

	return val
}