package main

// concurrent, recursive mergesort algorithm
func Concsort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	ch := make(chan int, len(arr))
	go concsort(arr, ch)

	sorted := make([]int, 0, len(arr))
	for a := range ch {
		sorted = append(sorted, a)
	}
	return sorted
}

// sorted elements are returned on the given channel
func concsort(arr []int, out chan int) {
	if len(arr) == 1 {
		out <- arr[0]
		close(out)
		return
	}

	m := len(arr) / 2
	arr1 := arr[:m]
	in1 := make(chan int, len(arr1))
	arr2 := arr[m:]
	in2 := make(chan int, len(arr2))

	go concsort(arr1, in1)
	go concsort(arr2, in2)

	// Keep track of next elements from each channel
	// Once we've pulled from both channels, compare and take the smallest
	var next1, next2 *int
	open1 := true
	open2 := true

	for open1 || open2 {
		// Pull next elements
		if open1 && next1 == nil {
			var n1 int
			n1, open1 = <-in1
			next1 = &n1
		} else if open2 && next2 == nil {
			var n2 int
			n2, open2 = <-in2
			next2 = &n2
		} else {
			if !open1 {
				out <- *next2
				next2 = nil
			} else if !open2 {
				out <- *next1
				next1 = nil
			} else {
				// compare elements
				if *next1 <= *next2 {
					out <- *next1
					next1 = nil
				} else {
					out <- *next2
					next2 = nil
				}
			}
		}
	}

	close(out)
}
