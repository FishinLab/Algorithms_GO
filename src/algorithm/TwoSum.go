package algorithm

import "fmt"

func TwoSum() {
	//var target int
	//nums := make([]int, size)

	target := 4
	nums := []int{1, 2, 3, 4}


	for _, i := range twoSum(nums, target) {
		fmt.Println(i)
	}
}

func twoSum(nums []int, target int) []int {
	var ret []int
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] + nums[j] == target {
				ret = []int{nums[i], nums[j]}
				break
			}
		}
	}
	return ret
}
