/*
    approach: counting
    - use the counter method that we use to validate paranthesis
    - dry-run to see when to add open-paran to output ( count > 1 )
    - dry-run to see when to add close-paran to output (count >= 1 )
    
    time = o(s)
    space = o(s)
*/
func removeOuterParentheses(s string) string {
    count := 0
    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
            if count > 1 {
                out.WriteByte(s[i])
            } 
        } else if s[i] == ')' {
            count--
            if count >= 1 {
                out.WriteByte(s[i])
            }
        }
    }
    return out.String()
}