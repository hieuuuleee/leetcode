func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    n := len(potions)
    res := make([]int, len(spells))
    for i, sp := range spells {
        t := float64(success) / float64(sp)
        pos := bisect(potions, t)
        res[i] = n - pos
    }
    return res
}

func bisect(arr []int, t float64) int {
    l, r := 0, len(arr)
    for l < r {
        m := l + (r - l) / 2
        if float64(arr[m]) < t {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}
