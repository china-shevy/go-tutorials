package arrayslice

import "testing"

func Test(t *testing.T) {
	array := [3]int{1, 2, 3}
	t.Logf("%v, %T", array, array)

	array2 := [2]int{1, 2}
	t.Logf("%v, %T", array2, array2)

	array3 := [...]int{1, 2, 3, 4}
	t.Logf("%v, %T", array3, array3)

	array4 := [4]int{1, 2}
	t.Logf("%v, %T", array4, array4)

	slice := []int{1, 2}
	t.Logf("%v, %T", slice, slice)
	t.Log("len", len(slice), "cap", cap(slice))

	setArray(array)
	t.Logf("%v, %T", array, array)

	setSlice(slice)
	t.Logf("%v, %T", slice, slice)

	s := array[:2]
	setSlice(s)
	t.Logf("%v, %T", array, array)
	t.Log("len", len(s), "cap", cap(s))

	setArray2(&array)
	t.Logf("%v, %T", array, array)
}

func setArray(a [3]int) {
	a[0] = 10
}

func setArray2(a *[3]int) {
	a[0] = 12
}

func setSlice(s []int) {
	s[0] = 10
}
