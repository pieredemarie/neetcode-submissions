func maxArea(heights []int) int {
	left, right := 0, len(heights)-1

	maxWater := 0

	for left < right {
		width := right - left 
		height := min(heights[left], heights[right])

		water := width * height
		if water > maxWater {
			maxWater = water
		}
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}

	return maxWater
}
