/*
    approach: 2 stacks
    - flatten / expand a bracket and combine with the parent
        - to expand a bracket, we need previous number
        - to combine a expanded bracket, we need previous parent.
    - previous / looking back for a char we last saw = stack is possible.
    - lets start with a numStack and strStack.
    - we will also need to maintain a current running string and current running num.
    - if we run into a char thats either a char or a number, we can concatnate that with their appropriate vars.
    - if we run into a open bracket '['
        - we know that the inside string of this opening bracket would need access to curr number and current parent we have in hand
        - because the curr number and current parent in hand is the previous digit and char for the inner string.
        - therefore if we pushed both number and parent str in their stacks, we can have access to them in o(1) time because of stack (they would be at the top)
    - after we go thru chars inside [..] - we will have a curr string in hand, curr number may be 0 if there were no nesting.
    - when we run into a closing bracket, its time to expand the curr string
        - which means we need previous number ( top of number stack ), lets call this k
        - repeat the curr string k times
        - now we need to combine with parent string ( top of parent stack ), lets call this parent
        - now to parent, append the expanded str
        - finally re-assign curr str to new parent str because its expanded and combined with it parent while curr was just whatever we had [...] inside a closing bracket
    
    Tldr: decode a bracket and combine it with its parent
        - decoding a bracket requires access to prev num and prev parent ( stack )
    
    We can do this with 1 stack too. Space would be the same regardless.
    Instead of having 1 dedicated stack for each use case, we can push a pair to 1 stack (previousChar, previousNum)
    
    
    time: o(s*k^countNumberOfNesting)
        s is the len of input string
        k is the digit in the input string
    or time can be;
        o(lenOfOutputStr)
    
    space: o(numSt) + o(strSt)
    
*/
// func decodeString(s string) string {
//     numSt := []int{}
//     strSt := []*strings.Builder{}
    
//     curr := new(strings.Builder)
//     num := 0
    
//     for i := 0; i < len(s); i++ {
//         char := s[i] // byte, i.e the ascii value
        
//         // if ascii value is >= ascii of 0 and <= ascii of 9, then this is a digit
//         if char >= '0' && char <= '9' {
//             num = num * 10 + int(char-'0')
//         } else if char >= 'a' && char <= 'z' {
//             curr.WriteByte(char)
//         } else if char == '[' {
//             numSt = append(numSt, num)
//             strSt = append(strSt, curr)
//             curr = new(strings.Builder)
//             num = 0
//         } else if char == ']' {
//             k := numSt[len(numSt)-1]
//             numSt = numSt[:len(numSt)-1]
//             expandedCurr := new(strings.Builder)
//             for i := 0; i < k; i++ {
//                 expandedCurr.WriteString(curr.String())
//             }
//             parent := strSt[len(strSt)-1]
//             strSt = strSt[:len(strSt)-1]
//             parent.WriteString(expandedCurr.String())
//             curr = parent
//         }
//     }
//     return curr.String()
// }



/*
    approach: recursion
    - instead of explicit stack, we can store the stack level vars in recursion stack
    - return curr string on ']'
    - return curr string on ']', the expand the returned string k time ( curr num in that recursion stack )
    - and parent will be curr str in that recursion stack.
    - Note: we wont make the "i" pointer as part of recursion ( even tho it feels like we should, or else how would a recursive call know where to start the for loop from ? )
        - if we do use i ptr as part of the recursion, then when we go back to a paused recursion call, the i will get backtracked automatically to where the i was where the recursive call started from. and we do not want this. DRAW THIS OUT ON YOUR BOARD TO UNDERSTAND BETTER.
*/
func decodeString(s string) string {
    ptr := 0
    var dfs func() string
    dfs = func() string {
        // base
        
        // logic
        curr := new(strings.Builder)
        num := 0
        for ptr < len(s) {
            char := s[ptr]
            if char >= '0' && char <= '9' {
                num = num * 10 + int(char-'0')
                ptr++
            } else if char >= 'a' && char <= 'z' {
                curr.WriteByte(char)
                ptr++
            } else if char == '[' {
                // ptr++ before because we need to escape this char and have the next recursive call start from next char, therefore.
                ptr++
                innerStr := dfs()
                expanded := new(strings.Builder)
                for i := 0; i < num; i++ {
                    expanded.WriteString(innerStr)
                }
                // reset, as we have just done this work ^
                num = 0
                // combine with parent, parent in recursion is curr var
                curr.WriteString(expanded.String())
                
            } else if char == ']' {
                ptr++
                return curr.String()
            }
        }
        return curr.String()
    }
    return dfs()
}














