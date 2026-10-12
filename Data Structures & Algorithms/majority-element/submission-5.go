func majorityElement(nums []int) int {
    seen := make(map[int]int)
	maj := len(nums)/2

	for _, num := range nums {
		seen[num]++
	}

	for num,count := range seen {
		if count > maj {
			return num
		}
	}

	return -1 // majority element always exists but we need to return something
}
