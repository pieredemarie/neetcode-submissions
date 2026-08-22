func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte,0,len(word1)+len(word2))

	n,m := len(word1), len(word2)

	p1, p2 := 0,0 

	for p1 != n && p2 != m {
		res = append(res,word1[p1])
		res = append(res,word2[p2])
		p1++
		p2++
	}

	for p1 != n {
		res = append(res, word1[p1])
		p1++
	}

	for p2 != m {
		res = append(res,word2[p2])
		p2++
	}

	return string(res)
}
