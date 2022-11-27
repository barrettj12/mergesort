package main

// non concurrent, recursive mergesort algorithm
func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	m := len(arr) / 2
	arr1 := mergesort(arr[:m])
	i1 := 0
	arr2 := mergesort(arr[m:])
	i2 := 0
	sorted := []int{}

	add1 := func() {
		sorted = append(sorted, arr1[i1])
		i1++
	}
	add2 := func() {
		sorted = append(sorted, arr2[i2])
		i2++
	}

	for i1 < len(arr1) || i2 < len(arr2) {
		if i1 >= len(arr1) {
			add2()
		} else if i2 >= len(arr2) {
			add1()
		} else {
			if arr1[i1] <= arr2[i2] {
				add1()
			} else {
				add2()
			}
		}
	}
	return sorted
}
