//结构体 链表
package main

import (
	"fmt"
)

//链表节点
type LinkNode struct {
	data interface{}
	next *LinkNode
}

//链表
type Link struct {
	head *LinkNode //头节点
	tail *LinkNode //尾节点
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	p.tail.next = node
	p.tail = node
}

func (p *Link) Trans() {
	q := p.head
	for q != nil {
		fmt.Println(q.data)
		q = q.next
	}
}
func main() {
	var link Link
	for i :=0; i<10;i++{
		link.
	}
}
