func distributeCookies(cookies []int, k int) int {
    ans := math.MaxInt64
    childs := make([]int, k)
    var dfs func(start int, cookieIdx int, maxi int)
    dfs = func(start int, cookieIdx int, maxi int) {
        // base
        if cookieIdx == len(cookies) {
            ans = min(ans, maxi)
            return
        }


        // logic
        for i := start; i < len(childs); i++ {
            childs[i] += cookies[cookieIdx]
            newM := max(maxi, childs[i])
            if childs[i] < ans {
                dfs(start, cookieIdx+1, newM)
            }
            childs[i] -= cookies[cookieIdx]
        }
    }
    dfs(0,0,math.MinInt64)
    return ans
}