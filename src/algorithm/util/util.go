package util

func Exists(tmp []int, num int) bool {
	for _, n := range tmp {
		if num == n {
			return true
		}
	}
	return false
}

func Remove(line []int, target int) []int {
	var ret []int
	for _, num := range line {
		if num != target {
			ret = append(ret, num)
		}
	}
	return ret
}

func Swap(array []int, i int, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
