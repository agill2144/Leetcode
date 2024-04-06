func reverseWords(s string) string {
    out := ""
    tmp := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            tmp.WriteByte(s[i])
        }
        if s[i] == ' ' || i == len(s)-1 {
            if tmp.Len() != 0 {
                if out == "" {
                    out = tmp.String()                    
                } else {
                    out = tmp.String() + " " + out                    
                }
                tmp.Reset()
            }
        }
    }
    return out
}