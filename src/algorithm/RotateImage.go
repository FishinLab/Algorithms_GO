package algorithm

import "fmt"

func RotateImage() {
	image := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	ret := rotateImage(image)
	fmt.Print(ret)
}

func rotateImage(image [][]int) [][]int {
	if image == nil {
		return nil
	}
	rowLength := len(image[0])
	colLength := len(image)
	ret := make([][]int, colLength)

	for j := 0; j < colLength; j++ {
		for i := 0; i < rowLength; i++ {
			ret[rowLength - j - 1] = make([]int, colLength)
			ret[colLength - i - 1][rowLength - j - 1] = image[i][j]
		}
	}
	return ret
}
