func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    d := make([]int, n)
    res := math.MinInt
    for i := n-1; i >= 0; i-- {
        if i + k < n {
            d[i] = d[i+k]
        }
        d[i] += energy[i]
        res = max(res, d[i])
    }
    return res
}
