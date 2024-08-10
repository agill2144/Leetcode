func longestPalindrome(s string) string {
    n := len(s)
    memo := make([][]*bool, n)
    for i := range memo {
        memo[i] = make([]*bool, n)
    }

    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if i >= j {
            return true
        }

        if memo[i][j] != nil {
            return *memo[i][j]
        }

        if s[i] == s[j] && dfs(i+1, j-1) {
            result := true
            memo[i][j] = &result
        } else {
            result := false
            memo[i][j] = &result
        }

        return *memo[i][j]
    }

    maxSize := 0
    maxStr := ""
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if dfs(i, j) {
                if j-i+1 > maxSize {
                    maxSize = j-i+1
                    maxStr = s[i:j+1]
                }
            }
        }
    }

    return maxStr
}
