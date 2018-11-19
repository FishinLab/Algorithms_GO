package algorithm

import "fmt"

func GetYanghuiTrangle() {
	target := 9
	for _, line := range getYanghuiTrangle(target) {
		for _, num := range line {
			fmt.Print(num)
			fmt.Print(SPACE)
		}
		fmt.Println()
	}
}

func getLine(line []int) []int {
	length := len(line)
	if length <= 0 {
		return nil
	}
	newLine := make([]int, length + 1)
	newLine[0] = line[0]
	newLine[length] = line[length - 1]
	for i := 1; i < length; i++ {
		newLine[i] = line[i - 1] + line[i]
	}
	return newLine
}

func getYanghuiTrangle(target int) [][]int {

	if target <= 0 {
		return nil
	}
	var ret = make([][]int, target)
	ret[0] = []int{1, 1}
	for i := 1; i < target; i++ {
		ret[i] = getLine(ret[i - 1])
	}

	return ret
}
