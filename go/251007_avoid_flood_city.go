func avoidFlood(rains []int) []int {
	av := make([]int, len(rains))
	l := map[int]int{}
	valid := make([]int, 0)
	for i, v := range rains {
		if v > 0 {
			av[i] = -1
			if _, exist := l[v]; exist {
				if idx := bi_search(valid, l[v]); idx < len(valid) {
					av[valid[idx]] = v
					copy(valid[idx:], valid[idx+1:])
					valid = valid[:len(valid)-1]
				} else {
					return []int{}
				}
			}
			l[v] = i
		} else {
			valid = append(valid, i)
		}
	}
	for i, v := range av {
		if v == 0 {
			av[i] = 1
		}
	}
	return av
}

func bi_search(array []int, value int) int {
	l, r := 0, len(array)-1
	for l <= r {
		m := (l+r)>>1
		if array[m] < value {
			l = m+1
		} else {
			r = m-1
		}
	}
	return l
}
