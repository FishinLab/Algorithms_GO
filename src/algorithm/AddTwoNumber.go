package algorithm

import (
	"algorithm/datastruct"
)

func AddTwoNumber() {
	one := datastruct.NewLinkNode()
	two := datastruct.NewLinkNode()

	one.BuildWithValues([]int{1,2,3})
	two.BuildWithValues([]int{3,4,5})
	addTwoNumber(one, two).PrintNodes()

	one.Clear()
	two.Clear()

	one.BuildWithValues([]int{1,2,3})
	two.BuildWithValues([]int{9,9,9})
	addTwoNumber(one, two).PrintNodes()

	one.Clear()
	two.Clear()

	one.BuildWithValues([]int{1,2,3})
	two.BuildWithValues([]int{9,9})
	addTwoNumber(one, two).PrintNodes()
}

func addTwoNumber(one *datastruct.LinkNode, two *datastruct.LinkNode) (*datastruct.LinkNode) {
	if one == nil || two == nil {
		return datastruct.NewLinkNode()
	}
	one.Reverse()
	two.Reverse()

	oneP := one
	twoP := two
	header := datastruct.NewLinkNode()

	if one.Len() > two.Len() {
		tmp := 0
		for oneP.Next != nil {
			p := datastruct.NewLinkNode()
			if twoP != nil {
				p.Val = oneP.Val + twoP.Val + tmp
			} else {
				p.Val = oneP.Val + tmp
			}

			if p.Val >= 10 {
				p.Val = p.Val - 10
				tmp = 1
			} else {
				tmp = 0
			}
			header.Append(p)
			oneP = oneP.Next
			twoP = twoP.Next
		}
		if tmp != 0 {

		}
	} else {
		for twoP.Next != nil {
			tmp := 0
			p := datastruct.NewLinkNode()
			if oneP!= nil {
				p.Val = twoP.Val + oneP.Val + tmp
			} else {
				p.Val = twoP.Val + tmp
			}

			if p.Val >= 10 {
				p.Val = p.Val - 10
				tmp = 1
			} else {
				tmp = 0
			}
			header.Append(p)
			oneP = oneP.Next
			twoP = twoP.Next
		}
	}
	header.Reverse()
	return header
}
