package main

import "fmt"

func main() {
	var arr = []int{34, 9, 0, 3, 4, 5, 43, 90, 23, 283458, 3, 6, 3456}

	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quickSort(arr []int, from, to int) {
	if from < to {
		q := partition(arr, from, to)
		quickSort(arr, from, q-1)
		quickSort(arr, q+1, to)
	}
}

func partition(arr []int, from, to int) int {
	key := arr[to]

	var i = from - 1

	for j := from; j < to; j++ {
		if arr[j] < key {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[to] = arr[to], arr[i+1]
	return i + 1
}
