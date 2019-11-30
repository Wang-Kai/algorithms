package main

import "fmt"

/*
	问题：
	其一个数组中，子数组之和的最大值。比如数组 [0,-9,3,4]，哪一段和最大呢？当然是 [3,4]，其和为 7，子数组必须是连续的
*/

func main() {
	maxSum := dynamicGetMaxSubArraySumPro(arr)

	println(maxSum)
}

var arr = []int{1, -2, 3, 5, -3, 2}

/*
	思路四：
	反思方法三的解法，其主要通过如下思路求解：通过求解 max{A[0], A[0, i], A[i, j]} 找到子数组和的最大值。通过求解一个小规模的子问题，要提供原问题的求解依据。
	在方案三中，我们分配了 start、all 两个长度为 n 的空间，目的是记录子问题的求解结果，但是我们实际急需要记住上一个问题的求解结果即可，不需要知道子子问题的解，遂把分配数组空间变为分配临时变量。
*/

func dynamicGetMaxSubArraySumPro(arr []int) int {
	var all, start int

	var n = len(arr)

	all = arr[n-1]
	start = arr[n-1]

	for i := n - 2; i >= 0; i-- {
		start = max(arr[i], arr[i]+start)
		all = max(start, all)
	}

	return all
}

/*
	思路三：
	运用动态规划的思想，拆分子问题，一个长度为 n 的数组 A，其和最大一定是求以下三种情况的最大值：
	1. A[0], 第一个数就是最大和子数组
	2. A[0, i] 其中 i <= n
	3. A[i, j] 其中 0 < i < j <= n

	假定 A[0, n] 最大子数组和为 all[0], 以 0 开始的最大子数组和为 Start[0],
*/

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func dynamicGetMaxSubArraySum(arr []int) int {

	var n = len(arr)
	var start, all = make([]int, n), make([]int, n)

	start[n-1] = arr[n-1]
	all[n-1] = arr[n-1]

	for i := n - 2; i >= 0; i-- {
		start[i] = max(arr[i], arr[i]+start[i+1])
		all[i] = max(start[i], all[i+1])

	}

	return all[0]
}

/*
	思路二：
	从三次遍历优化为两次遍历
*/

func sumMaxSubArrayPro(arr []int) int {
	var maxSum int

	for i := 0; i < len(arr); i++ {
		var sum = arr[i]

		fmt.Printf("Start from %dth \n", i)

		for j := i + 1; j < len(arr); j++ {
			fmt.Printf("Current Sum: %d \t, Next subSum %d \n", sum, sum+arr[j])

			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}

	}

	return maxSum
}

/*
	思路一：
	粗暴遍历，把所有的子数组全部穷举出来，比对哪个子数据的和最大
*/
func sumMaxSubArray(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	var maxSum int
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			var sum int
			for k := i; k <= j; k++ {
				sum += arr[k]
			}

			if maxSum < sum {
				maxSum = sum
			}
		}
	}
	return maxSum
}
