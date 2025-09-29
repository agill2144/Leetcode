func distributeCookies(cookies []int, k int) int {
    ans := math.MaxInt64
    childs := make([]int, k)
    var dfs func(ptr int, maxCookies int)
    dfs = func(ptr int, maxCookies int) {
        // base
        if ptr == len(cookies) {
            ans = min(ans, maxCookies)
            return
        }

        // logic
        for i := 0; i < len(childs); i++ {
            childs[i] += cookies[ptr]
            dfs(ptr+1, max(maxCookies, childs[i]))
            childs[i] -= cookies[ptr]
        }
    }
    dfs(0, 0)
    return ans
}