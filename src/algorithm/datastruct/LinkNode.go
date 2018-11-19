package datastruct

import (
	"fmt"
	"math/rand"
)

type LinkNode struct {
	Val int
	length int
	Next *LinkNode
}

func NewLinkNode() *LinkNode {
	var header *LinkNode
	header = new(LinkNode)
	header.Val = -1;
	header.length = 0
	header.Next = nil

	return header
}

func (header *LinkNode)Build(length int) *LinkNode {
	p := header

	for l := 0; l < length; l++ {
		node := new(LinkNode)
		node.Val = rand.Int()
		node.Next = nil

		p.Next = node
		p = p.Next
	}
	header.length = length
	return header
}

func (header *LinkNode)BuildWithValues(values []int) *LinkNode {
	p := header

	for l := 0; l < len(values); l++ {
		node := new(LinkNode)
		node.Val = values[l]
		node.Next = nil

		p.Next = node
		p = p.Next
	}
	header.length = len(values)
	return header
}

func (header *LinkNode)PrintNodes() {
	pointer := header.Next
	count := 0
	for pointer != nil {
		fmt.Printf("val: %d, count: %d \n", pointer.Val, count)
		pointer = pointer.Next
		count++
	}
}

func (header *LinkNode)last() (*LinkNode) {
	p := header.Next
	q := header
	for p != nil {
		p = p.Next
		q = q.Next
		if p.Next == nil {
			q.Next = nil
			break
		}
	}
	return p
}

func (header *LinkNode)Reverse() {
	p := header
	for p.Next != nil {
		q := header.last()

		q.Next = p.Next
		p.Next = q

		p = p.Next
	}
}

func (header *LinkNode)Append(p *LinkNode) {
	last := header
	for last != nil {
		if last.Next == nil {
			last.Next = p
			break
		}
		last = last.Next
	}
	header.length++
}

func (header *LinkNode)Clear() {
	header.Next = nil
	header.length = 0
}

func (header *LinkNode)Len() int {
	return header.length
}

func TestBuild() {
	header := NewLinkNode()
	header.Build(5)
	header.PrintNodes()

	values := []int{1, 2, 3, 4, 5}
	header.BuildWithValues(values)
	header.PrintNodes()

	header.Reverse()
	header.PrintNodes()
}
