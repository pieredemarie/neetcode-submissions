func isInBounds(newRow,newCol,n,m int) bool {
	return newRow >= 0 && newRow < n && newCol >= 0 && newCol < m
}
func islandsAndTreasure(grid [][]int) {
    n := len(grid)
	m := len(grid[0])


	inf := 2147483647 
	q := [][2]int{}

	for i :=0;i<n;i++ {
		for j := 0;j<m;j++ {
			if grid[i][j] == 0 {
				q = append(q,[2]int{i,j})
			}
		}
	}
	directions := [][2]int{{0,1}, {0,-1}, {1,0}, {-1,0}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		row, col := curr[0],curr[1]

		for _, d := range directions {
			newRow, newCol := row+d[0],col+d[1]

			if isInBounds(newRow,newCol, n,m) && grid[newRow][newCol] == inf {
				grid[newRow][newCol] = grid[row][col] + 1
				q = append (q, [2]int{newRow,newCol})
			}
		}
	}
}
