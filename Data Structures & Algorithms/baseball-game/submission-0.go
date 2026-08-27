func calPoints(operations []string) int {
	stack := make([]int,0, len(operations))
    res := 0
    for _, oper := range operations {
        if oper == "+" {
            first := stack[len(stack)-1]
            second := stack[len(stack)-2]
            sum := second + first
            stack = append(stack,sum)
        } else if oper == "D" {
            val := stack[len(stack)-1]
            stack = append(stack, val*2)
        } else if oper == "C" {
            stack = stack[:len(stack)-1]
        } else {
            num, _ := strconv.Atoi(oper)
            stack = append(stack, num)
        }
    }

    for _, num := range stack {
        res += num 
    }

    return res
}
