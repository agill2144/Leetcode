func minDistance(word1 string, word2 string) int {
    m := len(word1)+1
    n := len(word2)+1
    dp := make([][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([]int, n)
        dp[i][0] = i
    }
    for i := 0; i < n; i++ {
        dp[0][i] = i
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            iChar := word1[i-1]
            jChar := word2[j-1]
            if iChar == jChar {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j]))+1
            }
        }
    }
    return dp[m-1][n-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

/*
    approach: DFS, try all 3 choices at each ith char
    - using a for loop recursion with pivot idx on word1
    - recurse on the bigger word
        - why?
    - then apply all 3 choices on each ith char ( delete, insert and replace )
    - delete does not word2
    - insert and replace; depends on word2
        - i.e we only want to use chars from word2 and not all 26 characters.
        - for each char, we want to try insert/replace with all characters from word2 ( therefore another nested loop within. )
        - however insert only makes sense if the len of word1 < len word2
        - once we have used a char from word2, we need to tell our next recursion to skip that char
        - therefore this nested for loop will have its own pivot idx
    
    time:
    - 3 choices
    - m = len(word1)
    - n = len(word2)
    3^mn
    
    space: mn for recursion stack

*/
// func minDistance(word1 string, word2 string) int {
//     if len(word2) > len(word1) {return minDistance(word2, word1)}
//     minOp := math.MaxInt64
//     var dfs func(word string, start int, start2 int, count int)
//     dfs = func(word string, start int, start2 int, count int) {
//         // base
//         if word == word2 {
//             if count < minOp {minOp = count}
//             return
//         }
        
        
//         // logic
//         for i := start; i < len(word); i++ {
//             // delete
//             dfs(word[:i]+word[i+1:], i, start2, count+1)
//             // insert and replace
//             for j := start2; j < len(word2); j++ {
//                 if len(word) < len(word2) {
//                     // insert after ith position
//                     dfs(word[:i+1] + string(word2[j]) + word[i+1:], i+2, j+1, count+1)
//                 }
//                 // replace
//                 dfs(word[:i] + string(word2[j]) + word[i+1:], i+1, j+1, count+1)
//             }
//         }
//     }
//     dfs(word1, 0, 0, 0)
//     return minOp
// }