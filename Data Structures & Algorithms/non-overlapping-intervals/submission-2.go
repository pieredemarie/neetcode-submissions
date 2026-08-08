func isOverlapping(a, b []int) bool {
    return max(a[0], b[0]) < min(a[1], b[1])  
}

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })

    result := [][]int{intervals[0]}  
    ans := 0
    
    for i := 1; i < len(intervals); i++ {
        last := result[len(result)-1]
        curr := intervals[i]
        
        if isOverlapping(last, curr) {
            ans++
        } else {
            result = append(result, curr)
        }
    }
    return ans
}