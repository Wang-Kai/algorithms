package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr = []int{0, 10, 2, 13, 44, 5, 6, 9997, 108, 89}

	val := randomizedSelect(arr, 0, len(arr)-1, 10)

	println(val)
}

// randomizedSelect 从 arr[from, to] 中找出第 i 小的元素
func randomizedSelect(arr []int, from, to, i int) int {
	pivotIndex := randomizedPartition(arr, from, to)
	findIndex := from + i - 1

	fmt.Printf("==> Current Array: %v \n", arr)
	fmt.Printf("==> From: %d\t To: %d\t Pivot: %d \t FindIndex: %d\n", from, to, pivotIndex, findIndex)

	if findIndex == pivotIndex {
		return arr[findIndex]
	}

	if findIndex > pivotIndex {
		i = i - (pivotIndex - from + 1)
		return randomizedSelect(arr, pivotIndex+1, to, i)
	} else {
		return randomizedSelect(arr, from, pivotIndex-1, i)
	}
}

// randomizedPartition 随机选择一个元素作为 pivot
// 然后将指定的数组区间跨分为大于pivot & 小于pivot 的两部分
func randomizedPartition(arr []int, from, to int) int {
	rand.Seed(time.Now().Unix())
	randomIndex := rand.Intn(to-from+1) + from

	arr[to], arr[randomIndex] = arr[randomIndex], arr[to]

	var i = from - 1
	var key = arr[to]

	for j := from; j < to; j++ {
		if arr[j] < key {
			arr[i+1], arr[j] = arr[j], arr[i+1]
			i++
		}
	}

	arr[i+1], arr[to] = arr[to], arr[i+1]

	return i + 1
}

// findMaxAndMinNum 同时找到一个数组中的最大值 & 最小值
func findMaxAndMinNum(arr []int) (max, min int) {
	var from int

	if len(arr)%2 == 1 {
		// 数组长度为奇数
		from = 1
		max = arr[0]
		min = arr[0]
	} else {
		// 数组长度为偶数
		from = 2
		if arr[0] > arr[1] {
			max = arr[0]
			min = arr[1]
		} else {
			max = arr[1]
			min = arr[0]
		}
	}

	for i := from; i < len(arr); i += 2 {
		var relativeMax, relativeMin int

		if arr[i] > arr[i+1] {
			relativeMax = arr[i]
			relativeMin = arr[i+1]
		} else {
			relativeMax = arr[i+1]
			relativeMin = arr[i]
		}

		if relativeMax > max {
			max = relativeMax
		}

		if relativeMin < min {
			min = relativeMin
		}

	}

	return
}
