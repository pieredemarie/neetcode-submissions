func countHours(piles []int, k int) int {
	count := 0
	for _, banana := range piles {
		if banana > k {
			count += banana/k
			if banana % k != 0 {
				count++
			}
		} else {
			count++
		}
	}

	return count 
}
func minEatingSpeed(piles []int, h int) int {
	maxK := 0
	for _, banana := range piles {
		maxK = max(maxK, banana)
	}

	l, r := 1, maxK

	for l < r {
		mid := l + (r-l)/2
		hours := countHours(piles,mid)

		if hours <= h {
			r = mid 
		} else {
			l = mid + 1
		}
	}

	return l
}
