package main

// i want to ignore this comment
func add(x, y int) int {
	var z int
	z = (x + y) + 2
	return z
}

func searchLinear(arr []int, x int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}
