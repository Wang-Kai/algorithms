package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 6, 5, 4, 9009, 3, 23, 2398, 23, 4, 6, 36, 77, 647, 34673}
	heapSort(arr)
}

// heapSort 堆排序
// 1. 首先先建最大堆
// 2. 然后取到最大值与最后一个元素交换
// 3. 针对第一个元素保有最大堆性质
// 4. 依次往复
func heapSort(a []int) {
	buildMaxHeap(a)

	// 循环 length - 1 次
	for i := 0; i < len(a)-1; i++ {
		heapSize := len(a) - i

		a[0], a[heapSize-1] = a[heapSize-1], a[0]

		MAX_HEAPIFY(a[0:heapSize-1], 0)
	}

	fmt.Println(a)
}

// 时间复杂度 O(n)
func buildMaxHeap(a []int) {
	// 在一个堆中，叶节点从 floor(n/2) + 1 开始
	// 减一的目的是数组的元素从 0 开始
	var start = int(math.Floor(float64(len(a)/2))) - 1

	for i := start; i >= 0; i-- {
		MAX_HEAPIFY(a, i)
	}
}

// leftChild return left child index
// index from 0
func leftChild(i int) int {
	return (i+1)*2 - 1
}

// rightChild return right child index
// index from 0
func rightChild(i int) int {
	return (i + 1) * 2
}

// parent return parent index
// index from 0
func parent(i int) int {
	return int(math.Floor(float64(i+1)/2)) - 1
}

// MAX_HEAPIFY 保持最大堆性质
// 时间复杂度 O(lgn)
func MAX_HEAPIFY(a []int, i int) {
	fmt.Printf("Heapify array ==> %v\n", a)

	l, r := leftChild(i), rightChild(i)

	var largest int = i

	if l < len(a) && a[l] > a[i] {
		largest = l
	}

	if r < len(a) && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]

		MAX_HEAPIFY(a, largest)
	}
}
