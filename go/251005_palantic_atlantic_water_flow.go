func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 {
        return nil
    }
    m, n := len(heights), len(heights[0])
    p := make([][]bool, m)
    a := make([][]bool, m)
    for i := 0; i < m; i++ {
        p[i] = make([]bool, n)
        a[i] = make([]bool, n)
    }
    for j := 0; j < n; j++ {
        dfs(heights, 0, j, 0, p)
        dfs(heights, m-1, j, 0, a)
    }
    for i := 0; i < m; i++ {
        dfs(heights, i, 0, 0, p)
        dfs(heights, i, n-1, 0, a)
    }
    var ans [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if p[i][j] && a[i][j] {
                ans = append(ans, []int{ i, j })
            }
        }
    }
    return ans
}

func dfs(heights [][]int, i, j int, h int, check [][]bool) {
    if i < 0 || len(heights)-1 < i || j < 0 || len(heights[0])-1 < j {
        return
    }
    if check[i][j] || heights[i][j] < h {
        return
    }
    check[i][j] = true
    dfs(heights, i-1, j, heights[i][j], check)
    dfs(heights, i+1, j, heights[i][j], check)
    dfs(heights, i, j-1, heights[i][j], check)
    dfs(heights, i, j+1, heights[i][j], check)
}
