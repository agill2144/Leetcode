func decodeString(s string) string {
    i := 0    
    var dfs func() string
    dfs = func() string {
        // base
        if i == len(s) {return ""}
        
        // logic
        curr := new(strings.Builder)
        n := 0
        for i < len(s) {
            char := s[i]
            if char >= 'a' && char <= 'z' {
                curr.WriteByte(char)
                i++
            } else if char >= '0' && char <= '9' {
                n = n * 10 + int(char-'0')
                i++
            } else if char == '[' {
                i++
                inner := dfs()
                for i := 0; i < n ; i++ {
                    curr.WriteString(inner)
                }
                n = 0
            } else if char == ']' {
                i++
                return curr.String()
            }
        }
        return curr.String()
    }
    
    return dfs()
}