func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false 
	}

	m := len(matrix[0])
	l,r  := 0, m*n-1

	for l <= r {
		mid := l + (r-l)/2

		row := mid / m
		col := mid % m 

		if matrix[row][col] > target {
			r = mid - 1
		} else if matrix[row][col] < target {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
