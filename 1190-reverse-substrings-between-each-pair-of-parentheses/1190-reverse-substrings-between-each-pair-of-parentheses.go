func reverseParentheses(s string) string {
    st := []string{}
    curr := ""
    for i := 0; i < len(s); i++ {
        char := string(s[i])
        if char == "(" {
            // push parent to stack
            st = append(st, curr)
            curr = ""
        } else if char == ")" {
            // reverse curr
            currRev := reverse(curr)
            // combine with parent
            if len(st) != 0 {
                parent := st[len(st)-1]
                st = st[:len(st)-1]
                currRev = parent+currRev
            }
            curr = currRev
        } else {
            curr += char
        }
    }
    return curr
}

func reverse(s string) string {
    newS := ""
    for i := len(s)-1; i >= 0 ; i-- {
        newS += string(s[i])
    }
    return newS
}