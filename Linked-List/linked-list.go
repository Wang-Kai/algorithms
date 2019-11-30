package main

import (
	"errors"
	"fmt"
)

var Name string

/*
	单链表实现
*/

type Node struct {
	Key  string
	Next *Node
}

type SingleLinkedList struct {
	Head *Node
}

func (l *SingleLinkedList) Search(key string) *Node {
	// 依次遍历查找 Key 等于 key 的节点
	var tmpNode = l.Head
	for tmpNode != nil && tmpNode.Key != key {
		tmpNode = tmpNode.Next
	}

	if tmpNode == nil {
		return nil
	} else {
		return tmpNode
	}
}

func (l *SingleLinkedList) Insert(key string) {
	// 当前单链表为空
	if l.Head == nil {
		l.Head = &Node{
			Key:  key,
			Next: nil,
		}
		return
	}

	// 当前单链表不为空
	n := &Node{
		Key:  key,
		Next: l.Head,
	}
	l.Head = n
}

func (l *SingleLinkedList) List() {
	var tmpNode = l.Head
	var num int
	for tmpNode != nil {
		fmt.Printf("%d ==> Pointer: %p \t\t NextPointer: %p \t\t Key: %s \n", num, tmpNode, tmpNode.Next, tmpNode.Key)
		tmpNode = tmpNode.Next
		num++
	}
}

func (l *SingleLinkedList) Delete(key string) error {
	// 定义快慢指针，快指针比慢指针快 1 步
	var fastPointer, slowPointer = l.Head.Next, l.Head

	var delActonDone bool

	for fastPointer != nil {
		// 找到了要删除的节点
		if fastPointer.Key == key {
			delActonDone = true

			slowPointer.Next = fastPointer.Next
			fastPointer = fastPointer.Next
		} else {
			// 不是要删除的节点
			fastPointer = fastPointer.Next
			slowPointer = slowPointer.Next
		}
	}

	// 根据是否有执行 delete 操作
	if delActonDone {
		return nil
	} else {
		return errors.New("Not Found this node!")
	}
}
