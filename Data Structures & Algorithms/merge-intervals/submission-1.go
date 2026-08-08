func isOverlapping(a,b [] int) bool {
	return max(a[0], b[0]) <= min(a[1],b[1])
}
func mergeTwo(a,b []int) []int {
	return []int{min(a[0],b[0]),max(a[1],b[1])}
}
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals);i++ {
		last := res[len(res)-1]
		curr := intervals[i]

		if isOverlapping(last,curr) {
			res[len(res)-1] = mergeTwo(last,curr)
		} else {
			res = append(res, curr)
		}
	}
	
	return res
}
