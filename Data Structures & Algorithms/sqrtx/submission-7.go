func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else if x == 4 {
		return 2
	}
	l, r := 0, x

	for l < r {
		m := (r-l)/2 + l

		if m*m == x {
			return m
		} else if m*m <= x {
			l = m + 1
		} else {
			r = m
		}
	}

	return l - 1
}
