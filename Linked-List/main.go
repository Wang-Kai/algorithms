package main

import "fmt"

func main() {
	L := SingleLinkedList{}

	L.Insert("kiko")
	L.Insert("jobs1")
	L.Insert("jobs2")
	L.Insert("jobs3")
	L.Insert("jobs4")
	L.Insert("hello")
	L.Insert("steve")

	L.List()

	fmt.Println("========================")

	n := FindNodeFromTail(L, 5)

	fmt.Printf("==> %+v", n)
}

// 查找单链表倒数第 order 个元素
func FindNodeFromTail(L SingleLinkedList, order int) *Node {
	// 定义快慢指针
	var slowPointer, fastPointer = L.Head, L.Head

	// 快指针要先走 order 步
	for order > 0 && fastPointer != nil {
		fastPointer = fastPointer.Next
		order--
	}

	for fastPointer != nil {
		fastPointer = fastPointer.Next
		slowPointer = slowPointer.Next
	}
	return slowPointer
}
