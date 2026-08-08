func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int) 

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		graph[u] = append(graph[u],v)
	}

	return graph
}

func buildIndegree(graph map[int][]int, numCourses int) map[int]int {
    indegree := make(map[int]int)
    for i := 0; i < numCourses; i++ {
        indegree[i] = 0
    }
    
   
    for _, neighbors := range graph {
        for _, nei := range neighbors {
            indegree[nei]++
        }
    }
    return indegree

}


func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(prerequisites)
	indegree := buildIndegree(graph,numCourses)

	q := list.New()
	for node, deg := range indegree {
		if deg == 0 {
			q.PushBack(node)
		}
	}

	order := []int{}

	for q.Len() > 0 {
		front  := q.Front()
		node := front.Value.(int)
		q.Remove(front)

		order = append(order, node)
		for _, nei := range graph[node] {
			indegree[nei]--
			if indegree[nei] == 0 {
				q.PushBack(nei)
			}
		}
	}

	return len(order) == numCourses
}
