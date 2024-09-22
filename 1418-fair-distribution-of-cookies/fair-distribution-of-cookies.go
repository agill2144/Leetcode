func distributeCookies(cookies []int, k int) int {
    ans := math.MaxInt64
    var dfs func(start int, childs []int, cookieIdx int, maxi int)
    dfs = func(start int, childs []int, cookieIdx int, maxi int) {
        // base
        if cookieIdx == len(cookies) {
            ans = min(ans, maxi)
            return
        }

        // logic
        for i := start; i < len(childs); i++ {
            childs[i] += cookies[cookieIdx]
            newM := max(maxi, childs[i])
            dfs(start, childs, cookieIdx+1, newM)
            childs[i] -= cookies[cookieIdx]
        }
    }
    dfs(0, make([]int, k),0, 0)
    return ans
}