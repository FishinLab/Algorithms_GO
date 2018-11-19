package algorithm

import (
	"algorithm/util"
	"fmt"
)

func FullPermutation() {
	target := []int{1, 2, 3}
	for _, line := range fullPermutation(target) {
		for _, num := range line {
			fmt.Print(num)
			fmt.Print(SPACE)
		}
		fmt.Println()
	}
}

func DFS(array []int, cursor int, end int, ret [][]int) {
	if cursor == end {
		ret = append(ret, array)
	} else {
		for i := cursor; i <= end; i++ {
			util.Swap(array, cursor, i)
			DFS(array, cursor + 1, end, ret)
		}
	}
}

func fullPermutation(target []int) [][]int {
	if len(target) <= 1 {
		return [][]int{}
	}
	var ret [][]int
	DFS(target, 0, len(target) - 1, ret)
	return ret
}