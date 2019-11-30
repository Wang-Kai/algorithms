package main

import "fmt"

/*
	切钢板问题：

	一段长为 n 的钢板，寻找一种切割方法，使切割完的每个小段总计收入最高。假使钢板价格表如下：
	1英寸：1
	2：5
	3：8
	4：9
	5：10
	6：17
	7：17
	8：20
	9：24
	10：30
	现手上有一段长为 n 英寸的钢板，该如何切割呢？
*/

var p = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

func main() {
	bottomUpCutRodPlan(p, 4)

}

/*
	拓展:

	给出最大收益是不够的，关键要告知最大收益时的切法是怎样的
*/

func bottomUpCutRodPlan(p []int, n int) {
	if n == 0 {
		return
	}

	var CutPlan = make(map[int]int)
	var maxIncomeMemo = make(map[int]int)
	maxIncomeMemo[0] = 0

	for i := 0; i <= n; i++ {
		// 子问题规模为 i 的时候求最优解

		var maxIncome int
		for j := 1; j <= i; j++ {
			if maxIncome < p[j]+maxIncomeMemo[i-j] {
				CutPlan[i] = j
				maxIncome = p[j] + maxIncomeMemo[i-j]
			}
		}

		maxIncomeMemo[i] = maxIncome
	}

	fmt.Printf("CutPlan ==> %v \n", CutPlan)
	fmt.Printf("MaxIncome ==> %v \n", maxIncomeMemo)
}

/*
	思路三：
	方法一、二都是自顶向下的递归调用求解，是否可以自下而上的求解呢？
	要求解长度为 n 的钢板的最有解，先把其所有子问题求解完毕，并计入备忘录，最有求解 n 规模的解
*/

func bottomUpCutRod(p []int, n int) int {
	// 长度为 0 的钢板，收益也为 0
	if n == 0 {
		return 0
	}

	var maxIncomeMemo = make(map[int]int)
	maxIncomeMemo[0] = 0

	// 最优解中，左侧从第 i 英寸开始切割
	for i := 1; i <= n; i++ {
		var maxIncome int

		// 先求解该问题的所有子问题
		for j := 1; j <= i; j++ {
			maxIncome = max(maxIncome, p[j]+maxIncomeMemo[i-j])
		}
		maxIncomeMemo[i] = maxIncome
	}

	return maxIncomeMemo[n]
}

/*
	思路二：
	方案一的递归调用时间复杂度为 2 的 n 次方，很多计算任务花在了求解相同子任务上
	方法二将做一个备忘录，对于求解过的子问题将不再求解，直接从缓存中取答案
*/

var maxIncomeMemory = make(map[int]int)

func memoryCutRod(p []int, n int) int {
	// 如果钢板总长度为 0， 则直接返回
	if n == 0 {
		return 0
	}

	// 检查备忘录中是否已有
	if maxIncomeMemory[n] > 0 {
		return maxIncomeMemory[n]
	}

	var maxIncome int
	for i := 1; i <= n; i++ {
		maxIncome = max(maxIncome, p[i]+memoryCutRod(p, n-i))
	}

	// 记一笔备忘录
	maxIncomeMemory[n] = maxIncome

	return maxIncome
}

/*
	思路一：
	长度为 i 的钢板，价格为 Pi，长度为 n-i 的钢板，最大收益为 R(n-i)
	先把整条钢板切为两段，从左至右，第一段收益为 Pi, 该切法的最大收益为 Pi + R(n-i)
	第一段的长度 0 <= i <= n，遍历所有情况求最大值
*/

// p 为价格表，n 为钢板的长度
func cutROD(p []int, n int) int {
	// 长度为 0 的钢板，收益为 0
	if n == 0 {
		return 0
	}

	var maxIncome int
	for i := 1; i <= n; i++ {
		maxIncome = max(maxIncome, p[i]+cutROD(p, n-i))
	}

	return maxIncome
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
