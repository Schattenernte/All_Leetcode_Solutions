package main

func twoSum(nums []int, target int) []int {
	my_map := make(map[int]int)
	for index, number := range nums {
		if map_idx, pass := my_map[number]; pass {
			return []int{index, map_idx}
		}
		my_map[target-number] = index
	}
	return []int{}
}
