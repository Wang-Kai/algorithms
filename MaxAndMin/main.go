package main

import "fmt"

/*
	同时找出一个数组中的最大值 & 最小值
*/

var arr = []int{89, 0, 8, 23, 4, 2, 42, 3, 4, 2, 34, 23, 4}

func main() {
	max, min := SearchMaxMinPro(arr)
	fmt.Printf("Max: %d \t Min: %d \n", max, min)
}

/*
	思路二：
	方案一要比较 2N 次，可以把整个数组中的元素两两分组，比如 arr[2]、arr[3]，二者先比较做比较，然后大的和 max 比较，小的和 min 比较
	方案一中挪动两个位置要比较 4 次，现在仅需比较 3 次
*/

func SearchMaxMinPro(arr []int) (max, min int) {
	if len(arr) == 0 {
		return
	}

	// 如果数组只有一个值
	if len(arr) == 1 {
		min = arr[0]
		max = arr[0]
	}

	// 如果数组中只有两个值
	if arr[0] > arr[1] {
		max = arr[0]
		min = arr[1]
	} else {
		max = arr[1]
		min = arr[0]
	}

	if len(arr) == 2 {
		return
	}

	// 处理数组长度为奇数的情况
	var lenIsOdd int

	if len(arr)%2 == 1 {
		lenIsOdd = 1

		lastNum := arr[len(arr)-1]

		if lastNum > max {
			max = lastNum
		}
		if lastNum < min {
			min = lastNum
		}
	}

	for i := 2; i < len(arr)-lenIsOdd; i += 2 {
		var tmpMax, tmpMin int
		if arr[i] > arr[i+1] {
			tmpMax = arr[i]
			tmpMin = arr[i+1]
		} else {
			tmpMax = arr[i+1]
			tmpMin = arr[i]
		}

		if tmpMax > max {
			max = tmpMax
		}
		if tmpMin < min {
			min = tmpMin
		}
	}

	return
}

/*
	思路一：
	比较数组的前两个数，大的给 max，小的赋予 min
	指针从第 2 位开始往后走，分别与 min & max 比较，假使比 max 大就替换 max，假使比 min 小就替换 min，否则继续往后走
*/

func SearchMaxMin(arr []int) (max, min int) {
	if len(arr) == 0 {
		return
	}

	// 如果数组只有一个值
	if len(arr) == 1 {
		min = arr[0]
		max = arr[0]
	}

	// 如果数组中只有两个值
	if arr[0] > arr[1] {
		max = arr[0]
		min = arr[1]
	} else {
		max = arr[1]
		min = arr[0]
	}

	if len(arr) == 2 {
		return
	}

	for i := 2; i < len(arr); i++ {
		currentNum := arr[i]

		if currentNum > max {
			max = currentNum
		} else if currentNum < min {
			min = currentNum
		}
	}

	return
}

func max(arr []int) int {
	var max = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func min(arr []int) int {
	var min = arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}
