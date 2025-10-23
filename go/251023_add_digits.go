func hasSameDigits(s string) bool {
	d := getd(s)
	for len(d) > 2 {
		new_d := []int{}
		for i := 0; i < len(d)-1; i++ {
			new_d = append(new_d, (d[i]+d[i+1])%10)
		}
		d = new_d
	}
	return d[0] == d[1]
}
func getd(s string) []int {
    d := []int{}
    for _, c := range s {
        di := int(c - '0')
        d = append([]int{di}, d...)
    }
    return d
}
