package main

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}

func twoSum(nums []int, target int) []int {
	num2idx := make(map[int]int)

	for idx, num := range nums {
		diff := target - num

		_, ok := num2idx[diff]
		if ok {
			return []int{num2idx[diff], idx}
		}
		num2idx[num] = idx

	}

	return []int{}
}
