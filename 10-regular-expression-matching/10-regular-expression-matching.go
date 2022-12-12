// time : o(mn)
// space: o(mn)
func isMatch(s string, p string) bool {
    if s == p {
        return true
    }
    m := len(s)
    n := len(p)
    dp := make([][]bool, m+1)
    for idx, _ := range dp { dp[idx] = make([]bool, n+1)}
    dp[0][0] = true
    
    for j := 1; j < len(dp[0]); j++ {
        if p[j-1] == '*' {
            // we only have a 0 case when trying to match subproblem to ""
            // because the 1 case means, choose the preceding char to match to ""
            // we have preceding chars but nothing to match against, therefore
            // only 0 case...
            dp[0][j] = dp[0][j-2]
        }
    }
    
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            if p[j-1] != '*' {
                // when chars match, the remaining subproblem is up-left
                if p[j-1] == s[i-1] || p[j-1] == '.' {
                    dp[i][j] = dp[i-1][j-1]
                }
            } else {
                // zero case - 2 steps back ( i.e if we did not choose preceding char )
                // the remaining subproblem is by removing this char * and its preceding char
                // which is 2 steps back ( since we removed 2 chars )
                dp[i][j] = dp[i][j-2]
                
                // one case - only possible if preceding char matches word char
                if p[j-2] == s[i-1] || p[j-2] == '.' {
                    // then the remaining subproblem is above
                    // however we have choose either or, it could be true in zero case or vice versa, therefore its between 2 steps back || above
                    dp[i][j] = dp[i-1][j] || dp[i][j] 
                }
            }
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}
