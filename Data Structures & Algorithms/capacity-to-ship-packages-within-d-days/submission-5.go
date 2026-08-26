func CountDays(weights []int, cap int) int {
	sum := 0
	days := 1
	for _, wei := range weights {
		if sum + wei > cap {
			days++
			sum = wei
		} else {
			sum += wei
		}
	}

	return days
}

func shipWithinDays(weights []int, days int) int {
	r := 0
	l := 0

	for _, wei := range weights {
		l = max(l, wei)
		r += wei
	}

	for l < r {
		m := (r-l)/ 2 + l
		countedDays := CountDays(weights,m) 
		if countedDays > days {
			l = m + 1
		} else  {
			r = m
		} 
	}

	return l
}
