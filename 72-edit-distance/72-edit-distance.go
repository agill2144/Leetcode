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
    if word1 == word2 {
        return 0
    }
    m := len(word2)
    n := len(word1)
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    
    for j := 0; j < len(dp[0]); j++ { dp[0][j] = j }
    for i := 0; i < len(dp); i++ { dp[i][0] = i }
    
    for i := 1; i < len(dp); i++ { // row, word2 ( target char )
        for j := 1; j < len(dp[0]); j++ {  // word1 ( src char )
            word1Char := word1[j-1]
            word2Char := word2[i-1]
            if word1Char == word2Char {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1 
            }
        }
    }
    return dp[len(dp)-1][len(dp[0])-1]
}

func min(x, y int) int {
    if x < y {return x}
    return y
}