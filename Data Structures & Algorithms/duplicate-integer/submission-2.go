func hasDuplicate(nums []int) bool {
    for i := range nums {
		for j := range nums {
			if i != j && nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
