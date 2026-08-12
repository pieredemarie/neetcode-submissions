
// Rotten Fruit
func isInBounds(newRow, newCol, n, m int) bool {
	return newRow >= 0 && newRow < n && newCol >= 0 && newCol < m
}
func orangesRotting(grid [][]int) int {
	// 0 - empty cell
	// 1 - fruit
	// 2 - rotten fruit
	n := len(grid)
	m := len(grid[0])

	q := [][2]int{}
	minutes := -1
	hasFresh := false
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				hasFresh = true
			}
			if grid[i][j] == 2 {
				q = append(q, [2]int{i, j})
			}
		}
	}
	if !hasFresh {
		return 0
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	//BFS
	for len(q) > 0 {

		size := len(q)

		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]

			row, col := curr[0], curr[1]

			for _, d := range directions {
				newRow, newCol := row+d[0], col+d[1]

				if isInBounds(newRow, newCol, n, m) && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2 // fruit got rotten
					q = append(q, [2]int{newRow, newCol})
				}
			}
		}
		minutes++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}