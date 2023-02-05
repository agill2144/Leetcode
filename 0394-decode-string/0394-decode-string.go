func decodeString(s string) string {
    ptr := 0
    var dfs func() string
    dfs = func() string {
        // base
        // if start == len(s) {return ""}
        
        // logic
        curr := new(strings.Builder)
        n := 0
        for ptr < len(s) {
            char := s[ptr]
            if char >= '0' && char <= '9' {
                n = n * 10 + int(char-'0')
                ptr++
            } else if char == '[' {
                ptr++
                inner := dfs()
                decoded := new(strings.Builder)
                for k := 0; k < n; k++ {
                    decoded.WriteString(inner)
                }
                n = 0
                curr.WriteString(decoded.String())
            } else if char == ']' {
                ptr++
                return curr.String()
            } else if char >= 'a' && char <= 'z' {
                curr.WriteByte(char)
                ptr++
            }
        }
        return curr.String()
    }
    return dfs()
}