/*
    approach: dfs with backtracking
    - Since we have to generate every single valid combination of parans, dfs + backtracking makes sense
    - We are given n - which represents the number of parans pairs we are given
    - for example: n = 3 - we have 3 open and 3 close parans to work with
    - for example: n = 1 - we have 1 open and 1 close paran to work with
    
    - inorder to generate any valid parans... 
        - we have to start with an open paran right?
        - so we will keep track of number of opens we have used so far 
        - while openUsed < n , append an open paran to our current path ( ACTION )
        - we can potentially have more open parans, even tho we just used an open paran - so recurse ( RECURSE )
        - when our openUsed is no longer less than n, then we can no longer add more open parans - this means we have used up all our open parans allowed..
        - Now inorder to make path valid, we have to start closing our parans...
        - So as long as closeUsed < openUsed , add, close paran ( ACTION )
        - we may have more close parans to use so recurse ( RECURSE )
        - When we have used all open and all closed, our base case will add the path to a result array and return
        - Then once the call goes back to the close parans block, we will undo our action ( BACKTRACK )
        - Which will then end as there is nothing to do after backtracking our closUsed action
        - So recursion goes back to a previous parent call and at somepoint goes back to openUsed block where we have not used all open parans but goes to closeUsed call to generate other combinations with closed parans.
    
    Time: exponential - 2^n
    Space: ???
    
*/
func generateParenthesis(n int) []string {
    result := []string{}
    var backtrack func(openN, closeN int, path string)
    backtrack = func(openN, closeN int, path string) {
        // base
        if openN == n && closeN == n {
            result = append(result, path)
            return
        }
        
        // logic
        if openN < n {
            // action
            path += "("
            // recurse
            backtrack(openN+1, closeN, path)
            // backtrack
            path = path[:len(path)-1]
        }
        if closeN < openN {
            // action
            path += ")"
            // recurse
            backtrack(openN, closeN+1, path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    backtrack(0,0,"")
    return result
}




// /*
//     approach: dfs and backtracking to generate all permutations of a string s
//     then checking at the end if its valid, only then adding to result if it is valid
    
//     Result: Fail, TLE
// */
// func generateParenthesis(n int) []string {
   
//     // generate initial string, it may be invalid
//     s := []string{}
//     for i := 0; i < n; i++ {
//         s = append(s, ")")
//         s = append(s, "(")
//     }
    
//     result := map[string]bool{}
//     var backtrack func(start int)
//     // then generate all permutations and check if they are balanced
//     backtrack = func(start int) {
//         // base
//         if start == len(s) {
//             joined := strings.Join(s, "")
//             if isBalanced(joined) {
//                 result[joined] = true
//             }
//             return
//         }
        
//         // logic
//         for i := start; i < len(s); i++ {
//             s[i],s[start] = s[start],s[i]
//             backtrack(start+1)
//             s[i],s[start] = s[start],s[i]
//         }
//     }
//     backtrack(0)
//     r := []string{}
//     for k, _ := range result {
//         r = append(r, k)
//     }
//     return r
// }

// func isBalanced(s string) bool {
//     count := 0
//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' {count++}
//         if s[i] == ')' {count--}
//         if count < 0 {return false}
//     }
//     return count == 0
// }


