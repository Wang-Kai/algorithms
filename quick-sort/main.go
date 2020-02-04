package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr = []int{34, 9, 0, 3, 4, 5, 43, 90, 23, 28, 923328, 48923, 3, 458, 3, 6, 345, 6}
	quickSort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quickSort(arr []int, from, to int) {
	if from < to {

		q := randomizedPartiion(arr, from, to)

		quickSort(arr, from, q-1)
		quickSort(arr, q+1, to)
	}
}

func partition(arr []int, from, to int) int {
	fmt.Printf("Handled Array ==> %v\n", arr)

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

// randomizedPartiion 将随机的选择主元
func randomizedPartiion(arr []int, from, to int) int {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(to-from+1) + from

	arr[randomIndex], arr[to] = arr[to], arr[randomIndex]

	return partition(arr, from, to)
}
