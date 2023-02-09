func decodeString(s string) string {
    var dfs func() string
    i := 0
    dfs = func() string {
        // base 
        if i == len(s) {return ""}        
        curr := new(strings.Builder)
        num := 0

        // logic
        for i < len(s) {
            char := s[i]
            if char >= '0' && char <= '9' {
                num = num * 10 + int(char-'0')
                i++
            } else if char >= 'a' && char <= 'z' {
                curr.WriteByte(char)
                i++
            } else if char == '[' {
                i++
                innerStr := dfs()
                for k := 0; k < num; k++ {
                    curr.WriteString(innerStr)
                }
                num = 0
            } else if char == ']' {
                i++
                return curr.String()
            }
            
        }
        return curr.String()
    }
    return dfs()
}