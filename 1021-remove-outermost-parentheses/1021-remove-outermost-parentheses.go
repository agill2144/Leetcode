func removeOuterParentheses(s string) string {
    count := 0
    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            if count > 0 {
                out.WriteByte(s[i])
            }
            count++
        } else if s[i] == ')' {
            count--
            if count > 0 {
                out.WriteByte(s[i])
            }  

        }
        
    }
    
    return out.String()
}

// "(()())(())"
//  121210