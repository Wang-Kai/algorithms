package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 6, 5, 4, 90, 33}

	buildMaxHeap(arr)

	fmt.Printf("MAX HEAP ==> %v \n", arr)

	// fmt.Printf("HEAP_MAXIMUM => %d \n", HEAP_MAXIMUM(arr))
	fmt.Printf("==> %p \n", arr)
	fmt.Printf("Extract max ==> %d \n new heap is %v \n", HEAP_EXTRACT_MAX(arr), arr)
}

// HEAP_EXTRACT_MAX 去掉并返回堆中的最大值
func HEAP_EXTRACT_MAX(H []int) int {
	fmt.Printf("The HEAP is %v\n", H)

	var max = H[0]
	H[0], H[len(H)-1] = H[len(H)-1], H[0]
	heapSize := len(H) - 1

	fmt.Printf("Before MAX HEAPIFY: %v\n", H)

	MAX_HEAPIFY(H[0:heapSize], 0)

	return max
}

// HEAP_MAMXMUM 返回具有最大关键字的元素
func HEAP_MAXIMUM(H []int) int {
	return H[0]
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
	fmt.Printf("Do HEAPIFY for %v, and Index is %d \n", a, i)

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
