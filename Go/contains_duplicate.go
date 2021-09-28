package leetCode

// first think
func containsDuplicate_1(nums []int) bool {
	// originalArrayLen = len(nums)
	numsLen := len(nums)
	if numsLen <= 0 {
		return false
	}
	var set []int
	set = append(set, nums[0])
	for i := 1; i < numsLen; i++ {
		for _, v := range set {
			if v == nums[i] {
				return true
			}
		}
		set = append(set, nums[i])
	}
	return false
}

// Second think

func containsDuplicate_2(nums []int) bool {
	dict := make(map[int]int)
	for _, v := range nums {
		dict[v] += 1
		if dict[v] == 2 {
			return true
		}
	}
	return false

}
