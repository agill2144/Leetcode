func longestPalindrome(s string) string {
    memo := make([][]*bool, 1001) // 0 = false, 1 = true
    for i := 0; i < len(memo); i++ {
        memo[i] = make([]*bool, 1001)
    }
    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if i >= j {return true}
        
        if memo[i][j] != nil {
            return *memo[i][j]
        }

        if s[i] == s[j] {
            memo[i][j] = toPtrBool(dfs(i+1, j-1))
        } else {
            memo[i][j] = toPtrBool(false)
        }
        return *memo[i][j]
    }
    maxSize := 0
    maxStr := ""
    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            if dfs(i, j) {
                if j-i+1 > maxSize {maxSize = j-i+1; maxStr = s[i:j+1]}
            }
        }
    }
    return maxStr
}

func toPtrBool(x bool) *bool {return &x}