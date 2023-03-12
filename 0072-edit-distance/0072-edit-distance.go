func minDistance(word1 string, word2 string) int {
    m := len(word1)+1
    n := len(word2)+1
    dp := make([][]int, m)
    for i := 0; i < len(dp); i++ {dp[i] = make([]int, n)}
    
    for i := 0; i < m; i++ {dp[i][0] = i}
    for j := 0; j < n; j++ {dp[0][j] = j}
    
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

// func minDistance(word1 string, word2 string) int {
//     if word1 == word2 {return 0}
//     if len(word2) > len(word1) {return minDistance(word2, word1)}
//     minOp := math.MaxInt64
//     var dfs func(word string, pivot, pivot2, count int)
//     dfs = func(word string, pivot,  pivot2, count int) {
//         // base
//         if word == word2 {
//             // fmt.Println(word, word2, count, minOp)
//             if count < minOp {minOp = count}
            
//         }
        
        
//         // logic
//         for i := pivot; i < len(word); i++ {

//             // delete does not depend on word2, we can do this outside.
//             dfs(word[0:i]+word[i+1:], i, pivot2, count+1)

//             // for insert/replace, we want to use our options from word2
//             for j := pivot2; j < len(word2); j++ {

//                 // insert, only makes sense if len of our word < len of target word.
//                 if len(word) < len(word2) {
//                     // insert after ith position
//                     newWord := word[:i+1] + string(word2[j]) + word[i+1:]
//                     // we have use the jth char, move to next, therefore j+1
//                     dfs(newWord, i+2, j+1, count+1)
//                 }

                
//                 // replace, only makes sense if our ith char != jth char
//                 if word[i] != word2[j] {
//                     // we have use the jth char, move to next, therefore j+1
//                     dfs(word[0:i]+string(word2[j])+word[i+1:], i+1, j+1, count+1)
//                 }
//             }

//         }
        
//     }
//     dfs(word1, 0, 0, 0)
//     return minOp
// }