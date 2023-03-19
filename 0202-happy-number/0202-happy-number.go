func isHappy(n int) bool {
    seen := map[int]bool{}
    var dfs func(n int) bool
    dfs = func(n int) bool {
        // base
        if n == 1 {
            return true
        }
        total := 0
        for n > 0 {
            total += (n%10) * (n%10)
            n /= 10
        }
        _, ok := seen[total]
        if ok { return false }
        seen[total] = false
        return dfs(total)
    }
    return dfs(n)
}