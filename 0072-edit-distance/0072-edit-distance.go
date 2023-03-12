/*
    We are given word1 and word2
    We have to transform word1 to word2
    For each char in word1 we have 3 choices/decisions = insert, add , delete
        - given choices for us to pick and figure out what makes sense - DP hint
    We also have to do this in minimum number of operations.
        - asking for optimal answers, also another DP hint
    
    For each character in word1, we will explore all 3 choices in a DFS manner
    - update: only use chars from word2
    - delete: delete this char
    - add: only use chars from word1
    We will do this in a DFS manner.
    
    After drawing the recursion tree, we see repeated subproblems
    - This is indeed a DP problem
    - We can solve the smallest subproblem and work ourself up to the bigger problem (bottom up)
    - For each subproblem, we have 3 choices to make
    - Once we know the answer from all 3 choices, we will pick the smallest one, as we are trying to minimize the num of operations needed
    
    approach: bottom up dp
    - create a dp matrix
    - m = word2 ( target word )
    - n = word1 ( source word )
    - we will take a dummy row and dummy col
    - Each row in dp matrix represent each character in word2 ( target char )
    - Each col in dp matrix represent each character in word1 ( source char )
    - Then we will fill out the dp matrix - by simplying drawing the recursion tree for each subproblem
    - Initially we will start with empty char (0,0) - this requires 0 operations to transform - therefore dp[0][0] = 0
    - For all cols in first row;
        - target word = "" ( dp[0][0] ) the dummy char
        - inorder to form "" ( as target ), we have to delete all characters we have in each subrpoblem
        - if the src word is horse
        - then each char in src word is a col
        - then the smallest subproblem would be given word1 "h", transform into word2 "" in min number of operations - which can be done by deleting the "h" - so 1 delete operation
        - then the next subproblem would be be given word1 "ho", transform into word2 "" in min number of operations - which can be done by deleting the "ho" - so 2 delete operation
        - this way the entire first row for each col, the answer will the len of src word (subproblem word to transform to ) with target word2 as "" - can only be done by deleting
            all N characters in current subproblem.
    - For 0th col of all rows;
        - given word2 = "ros" (target word)
        - src word1 = "" - the dummy char we took
        - inorder to make word1 "" into word2 "ros" - the only way to do this is to insert 3 characters, so 3 min operations needed
        - Now fill out the 0th col for each row
    
    - bottom up, start from the back of word2 and look back
    - if the last char of word2 == last char of current word1, then it means our subproblem is remove the matching chars and solve the remaining subproblem
    - for each char of word2, we have to make 3 decisions
        - add : right above current cell in dp matrix
        - delete : 1 step back in same row
        - update : diagonal up left 
    - once you make the dp matrix on board, you will notice the above pattern ^
    
    time : o(mn)
    space: o(mn)
*/
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