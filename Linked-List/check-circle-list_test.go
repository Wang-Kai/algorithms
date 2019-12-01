package main

import (
	"testing"
)

func TestCheckRing(t *testing.T) {
	var hNode = new(Node)

	var n1 = &Node{
		Key: "first node1",
	}
	var n2 = &Node{
		Key:  "first node2",
		Next: n1,
	}
	var n3 = &Node{
		Key:  "first node3",
		Next: n2,
	}
	var n4 = &Node{
		Key:  "first node4",
		Next: n3,
	}
	var n5 = &Node{
		Key:  "first node5",
		Next: n4,
	}
	var n6 = &Node{
		Key:  "first node6",
		Next: n5,
	}
	hNode = &Node{
		Key:  "first node",
		Next: n6,
	}
	n1.Next = n5

	list := &SingleLinkedList{
		Head: hNode,
	}

	// 检查链表是否有环
	hasRing := CheckRing(list)
	t.Log(hasRing)

	// 求环的长度
	ringLen := GetRingLength(list)
	t.Logf("The length of ring ==> %d \n", ringLen)

	// // 求环点
	crossPoint := FindCrossPoint(list, 5)
	t.Logf("Cross Point ==> %+v \n", crossPoint)

}

/*
	FindCrossPoint 已知一个单链表有环，并且知道环长度，求环点

	思路：
	通过快慢指针的方式，快指针先走换长度步数，然后快慢指针以同样的速度走，当两个指针相遇的时候，慢指针所在的就是环点
*/
func FindCrossPoint(l *SingleLinkedList, ringLen int) *Node {
	var (
		quickP = l.Head
		slowP  = l.Head
	)

	// 快指针先走环长步数
	for quickP.Next != nil && ringLen > 0 {
		quickP = quickP.Next
		ringLen--
	}

	// 快慢指针以相同的速度走，直到相遇
	for quickP != slowP {
		quickP = quickP.Next
		slowP = slowP.Next
	}

	return slowP
}

/*
	GetRingLength	求单链表环的长度

	思路：
	通过快慢指针的方式找到指针的交点，接着快指针不动，慢指针继续走，当两个指针再次相遇的时候，慢指针所走过的长度就是环长。
*/
func GetRingLength(l *SingleLinkedList) int {
	var ringLen int

	var (
		quickP = l.Head.Next.Next
		slowP  = l.Head
	)

	for quickP != nil && quickP.Next != nil {
		// 两个指针相交，该链表有环
		if quickP == slowP {
			// TODO
			ringLen = measuringRingLen(quickP, slowP)
			break
		}

		// 快指针每次走两步，慢指针每次走一步
		quickP = quickP.Next.Next
		slowP = slowP.Next
	}

	return ringLen
}

// 开始测量环长
func measuringRingLen(a, b *Node) int {
	b = b.Next
	var ringLen int = 1
	for a != b {
		b = b.Next
		ringLen++
	}

	return ringLen
}

/*
	CheckRing 检查一个单链表是否有环

	思路：
	设置一对快慢指针，慢指针每次移动一步，快指针每次移动两步，假使该单链表有环，则两个指针必定相交
*/
func CheckRing(l *SingleLinkedList) bool {
	var (
		quickP  = l.Head
		slowP   = l.Head
		hasRing bool
	)

	for quickP != nil && quickP.Next != nil {
		slowP = slowP.Next
		quickP = quickP.Next.Next

		if slowP == quickP {
			hasRing = true
			break
		}
	}

	return hasRing
}
