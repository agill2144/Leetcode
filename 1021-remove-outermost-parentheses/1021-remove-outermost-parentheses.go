// use balancing counter thats used to validate paranthesis
// then draw it out to see when does it make sense to add '(' ; count >= 2
// and when does it make sense to add ')' ; count >= 1
func removeOuterParentheses(s string) string {
    if len(s) == 0 {return ""}
    out := new(strings.Builder)
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
            if count >= 2 {
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

/*

(()())(())
12
(()


*/