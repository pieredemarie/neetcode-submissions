func twoSum(nums []int, target int) []int {
    count := make(map[int]int)

	for i, num := range nums {
		difference := target - num

		if j, found := count[difference]; found {
			return []int{j,i}
		}

		count[num] = i
	}

	//it is said that there's always solution
	return []int{0,0}
}
