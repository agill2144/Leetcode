func removeOuterParentheses(s string) string {
    dq := []byte{}
    count := 0
    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        dq = append(dq, s[i])
        if s[i] == '(' {
            count++
        } else if s[i] == ')' {
            count--
        }
        
        if count == 0 {
            dq = dq[1:]
            dq = dq[:len(dq)-1]
            for len(dq) != 0 {
                out.WriteByte(dq[0])
                dq = dq[1:]
            }            
        }
    }
    
    return out.String()
}

// "(()())(())"
//  121210