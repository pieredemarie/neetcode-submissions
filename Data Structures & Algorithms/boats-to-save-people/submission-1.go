import "slices"
func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	boats := 0

	left, right := 0, len(people)-1

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
			right--
		} else if people[left]+people[right] > limit {
			right--
		}

		boats++
	}

	return boats
}
