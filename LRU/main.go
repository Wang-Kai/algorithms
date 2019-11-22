package main

import "fmt"

/*
LRU 实现

LRU Cache 即 Least Recently Used 缓存策略，最近最少使用只是思考问题的角度，LRU算法的设计原则是：如果一个数据在最近一段时间没有被访问到，那么在将来它被访问的可能性也很小。也就是说，当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰。
*/

const CACHE_SIZE = 10

func main() {
	c := &Cache{}

	c.Set("na23me", "steve")
	c.Set("a23ge", "23")
	c.Set("ag34e", "23")
	c.Set("na12me", "steve")
	c.Set("ag98e", "23")
	c.Set("age", "23")
	c.Set("lodve", "kiko")
	c.Set("lovde", "kikasdfo")
	c.Set("lo2ve", "kikafasdfo")
	c.Set("love", "kika31243sdfo")
	c.Set("lov3de", "kika234234sdfo")
	c.Set("love", "kikas234dfo")
	c.Set("lov23e", "23423")
	c.Set("lo34ve", "kikas4324==dfo")
	c.Set("lo3423ve", "kikas4324==dfo")
	c.Set("lo3445ve", "kikas4324==dfo")

	c.listAll()
}

type Page struct {
	prev *Page
	next *Page
	val  string
	key  string
}

type Cache struct {
	head *Page
	tail *Page
	len  int
}

func (c *Cache) listAll() {
	p := c.head
	var num int
	for p != nil {
		fmt.Printf("%d ==> Pointer: %p \t\t nextPointer: %p \t\t key: %s \t\t value: %s \n", num, p, p.next, p.key, p.val)
		p = p.next

		num++
	}
}

func (c *Cache) Set(key, val string) {
	// cache 为空，添加第一个节点
	if c.head == nil {
		c.head = &Page{
			prev: nil,
			next: nil,
			val:  val,
			key:  key,
		}
		c.tail = c.head
		c.len++
		return
	}

	// 检查是否有重复
	var tmpPage = c.head

	for tmpPage != nil && tmpPage.key != key {
		tmpPage = tmpPage.next
	}

	// 有 key 重复情况发生
	if tmpPage != nil {
		var newHead = &Page{
			key:  key,
			val:  val,
			prev: nil,
		}

		if tmpPage.prev == nil {
			// 该节点为第一个节点
			// 该节点被删除
			newHead.next = tmpPage.next
		} else if tmpPage.next == nil {
			// 该节点为最后一个节点
			tmpPage.prev.next = nil
			newHead.next = c.head
			c.tail = newHead
		} else {
			tmpPage.prev.next = tmpPage.next
			tmpPage.next.prev = tmpPage.prev
			newHead.next = c.head
			c.head.prev = newHead
		}

		c.head = newHead
		tmpPage = nil // 移除这个节点
		return
	}

	// cache 满了，移除最后一个，把新的节点加入到最前面
	if c.len == CACHE_SIZE {
		tailPage := c.tail
		tailPage.prev.next = nil
		c.tail = tailPage.prev

		newPage := &Page{
			prev: nil,
			next: c.head,
			key:  key,
			val:  val,
		}
		c.head.prev = newPage
		c.head = newPage

	} else if c.len < CACHE_SIZE {
		// cache 没满，新节点插入到最前
		var newHead = &Page{
			key:  key,
			val:  val,
			next: c.head,
			prev: nil,
		}
		c.head.prev = newHead
		c.head = newHead
		c.len++
	}
}

func (c *Cache) Get(key string) string {
	var currentPage = c.head

	for currentPage != nil && currentPage.key != key {
		currentPage = currentPage.next
	}

	return currentPage.val
}
