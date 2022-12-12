/*
    approach: BFS
    - We need to try forming every single substring after removing each char from string s
    - So we start with a original string s, and then we need to loop over this string and remove ith char and check if its valid
    - We can do this using a BFS
    - Enqueue our start node ( the original input string s)
    - Then process the queue
    - While processing the queue, we will take the size of the queue, we do want to do this level by level
    - Because at the very first level, where we attempt to form a substring after removing ith char
    - If we do find a valid substring, we will not process any more level below it
    - Why not? because we want the largest substr len, proccessing a child as a parent will only reduce the string len, therefore.
    - How do we make sure a child thats already in the queue should not be enqueud again?
        - Create a visited map/set to verify before enqueue-ing.
    - How do we check whether a string is valid ?
        - Use a separate function to validate a string of paranthesis is a valid string of paranthesis..
        - But how?
        - Since we only have 1 type of paran, () , then we can simply use a count to check if string of parans is balanced or not
        - loop over the string s
        - if char is a open paran
            - count++
        - if char is a close paran
            - count--
        - if count while looping goes negative, return false.
            - Why?
            - Example: ()())(
            - Accoording to above logic, count will be 0 at the end , claiming that the string of parans are valid.
            - However its not valid, because of the placements...
            - When this scenario happens, our count will go negative , which is a sign of this scenario, so just return false.
        - Finally if we looped over successfully, return whether count == 0 ( if yes, string of parans is balanced, if not, then string of parans is not balanced. )

    time: n^n
    space: n^n
    
*/

func removeInvalidParentheses(s string) []string {
    visited := map[string]bool{}
    q := []string{s}
    result := []string{}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if isBalanced(dq) {
                result = append(result, dq)
            } else {
                // enqueue all of this node's children - as long as they are not already visited
                for i := 0; i < len(dq); i++ {
                    if dq[i] == '(' || dq[i] == ')' {
                        child := string(dq[0:i]) + string(dq[i+1:])
                        if _, ok := visited[child]; !ok {
                            q = append(q, child)
                            visited[child] = true
                        }
                    }
                }
            }
            
            
            qSize--
        }
        // exit if we found our answer at a level 
        if len(result) > 0 {
            break
        }
    }
    return result
}


/*
    approach: DFS
    - The same logic is possible using DFS
    - The only difference is we potentially go down path that leads to the smallest substr and we would have to reset
    - However if we know the answers are at the deepest lowest level, then it makes sense to use DFS
    - Otherwise BFS
    
    Time: n^n
    space: visited map + recursion stack - n^n
*/
// func removeInvalidParentheses(str string) []string {
//     visited := map[string]bool{}
//     max := 0
//     result := []string{}
//     var dfs func(s string)
//     dfs = func(s string) {
//         // base
//         if len(s) < max || visited[s] {return}
//         visited[s] = true
//         if isBalanced(s) {
//             if len(s) > max {
//                 max = len(s)
//                 result = []string{s}
//             } else if len(s) == max {
//                 result = append(result, s)
//             }
//             return
//         }
        
//         for i := 0; i < len(s); i++ {
//             if s[i] == ')' || s[i] == '(' {
//                 child := string(s[0:i]) + string(s[i+1:])
//                 dfs(child)
//             }
//         }
//     }
//     dfs(str)
//     return result
// }


func isBalanced(s string) bool {
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
        } else if s[i] == ')' {
            count--
        }
        if count < 0 {return false}
    }
    return count == 0
}