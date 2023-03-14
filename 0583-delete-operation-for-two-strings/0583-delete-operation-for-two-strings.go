/*
    same question as edit distance
    except only 1 choice given in this and this is delete a char
    
    when two characters match, nothing to delete
    when two characters do not match;
    - try deleting from word1 - you will see repeated subproblem
    - try deleting from word2 - you will see repeated subproblem
    - pick min of above two repeating problem answers + 1(for deleting either of the char)
    
    time: o(mn)
    space: o(mn)
    
*/
func minDistance(word1 string, word2 string) int {
    m := len(word1)+1
    n := len(word2)+1
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        dp[i][0] = i
    }
    for j := 0; j < n; j++ {
        dp[0][j] = j
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1] 
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1])+1
            }
        }
    }
    return dp[m-1][n-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}