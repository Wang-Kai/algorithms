package main

import "fmt"

// 活动选择问题
// N 场演唱会, 以 [{startTime, endTime}…] 的形式给出, 计算出最多能听几场演唱会
func main() {
	var activities = []Activity{
		Activity{
			From: 1,
			To:   4,
		},
		Activity{
			From: 3,
			To:   5,
		},
		Activity{
			From: 0,
			To:   6,
		},
		Activity{
			From: 5,
			To:   7,
		},
		Activity{
			From: 3,
			To:   9,
		},
		Activity{
			From: 5,
			To:   9,
		},
		Activity{
			From: 6,
			To:   10,
		},
		Activity{
			From: 8,
			To:   11,
		},
		Activity{
			From: 8,
			To:   12,
		},
		Activity{
			From: 2,
			To:   14,
		},
		Activity{
			From: 12,
			To:   16,
		},
	}

	// 以结束时间，对所有活动从早到晚排序
	sArr, fArr := sortActivities(activities)
	activirySelector(sArr, fArr)
}

func sortActivities(activities []Activity) ([]int, []int) {
	for i := 1; i < len(activities); i++ {
		key := activities[i]

		j := i - 1

		for j >= 0 && activities[j].To > key.To {
			activities[j+1] = activities[j]
			j--
		}

		activities[j+1] = key
	}
	var fromTimes = make([]int, len(activities))
	var finishTimes = make([]int, len(activities))

	for _, a := range activities {
		fmt.Printf("StartTime: %d \t EndTime: %d\n", a.From, a.To)

		fromTimes = append(fromTimes, a.From)
		finishTimes = append(finishTimes, a.To)
	}

	return fromTimes, finishTimes
}

type Activity struct {
	From int
	To   int
}

func activirySelector(s, f []int) {
	a1 := &Activity{
		From: s[0],
		To:   f[0],
	}

	fmt.Printf("==> %v\n", a1)

	var currentFinishTime = f[0]

	for i := 1; i < len(s); i++ {

		if s[i] > currentFinishTime {
			fmt.Printf("==> %v\n", &Activity{
				From: s[i],
				To:   f[i],
			})
			currentFinishTime = f[i]
		}

	}
}
