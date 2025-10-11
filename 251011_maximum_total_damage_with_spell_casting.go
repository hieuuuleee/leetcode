func maximumTotalDamage(power []int) int64 {
    f := map[int]int{}
    for _, p := range power {
        f[p]++
    }
    nums := []int{}
    for k, _ := range f {
        nums = append(nums, k)
    }
    sort.Ints(nums)
    n := len(nums)
    dp := make([]int, n)
    dp[0] = nums[0] * f[nums[0]]
    for i := 1; i < n; i++ {
        num := nums[i]
        dp[i] = max(dp[i-1], num * f[num])
        for j := i-1; j >= 0 && j >= i-3; j-- {
            num2 := nums[j]
            if num - num2 > 2 {
                dp[i] = max(dp[i], num * f[num] + dp[j])
            }
        }
    }
    return int64(dp[n-1])
}
