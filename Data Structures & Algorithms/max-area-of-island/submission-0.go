func isInBounds(rows,cols,newRow,newCol int) bool {
	return newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols
}
func bfs(startRow, startCol int, visited [][]bool, grid [][]int) int {
	area := 0
	rows := len(grid)
	cols := len(grid[0])

	q := list.New()
	q.PushBack([2]int{startRow,startCol})
	visited[startRow][startCol] = true 
	area++
	directions := [][2]int{{-1,0}, {1,0}, {0,-1}, {0,1}}  


	for q.Len() > 0 {
		front := q.Front()
		coords := front.Value.([2]int)
		q.Remove(front)

		row := coords[0]
		col := coords[1]

		for _, dir := range directions {
			newRow := row + dir[0]
			newCol := col + dir[1]

			if isInBounds(rows,cols,newRow,newCol) {
				if grid[newRow][newCol] == 1 && !visited[newRow][newCol] {
					visited[newRow][newCol] = true 
					q.PushBack([2]int{newRow, newCol})
					area++
				}
			}
		}
	}

	return area
}
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for i := 0;i < rows;i++ {
		for j := 0;j<cols;j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				area := bfs(i,j,visited,grid)
				maxArea = max(area, maxArea)
			}
		}
	}

	return maxArea   
}
