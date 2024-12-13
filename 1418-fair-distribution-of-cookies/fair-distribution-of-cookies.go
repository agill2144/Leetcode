func distributeCookies(cookies []int, k int) int {
    minVal := math.MaxInt64
    childs := make([]int, k)
    var dfs func(start ,cookiePtr, maxCookies int)
    dfs = func(start , cookiePtr, maxCookies int) {
        // base
        if cookiePtr == len(cookies) {
            minVal = min(minVal, maxCookies)
            return
        }

        // logic
        for i := start; i < len(childs); i++ {
            if childs[i] + cookies[cookiePtr] > minVal {continue}
            // action
            childs[i] += cookies[cookiePtr]

            dfs(start, cookiePtr+1, max(maxCookies, childs[i]))
            childs[i] -= cookies[cookiePtr]
        }
    }
    dfs(0, 0, 0)
    return minVal
}