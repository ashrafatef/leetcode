package main

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v] += 1
		if dict[v] == 2 {
			return true
		}
	}
	return false

}
