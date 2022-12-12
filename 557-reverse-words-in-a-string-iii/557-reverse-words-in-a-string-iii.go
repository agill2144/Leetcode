func reverseWords(s string) string {
    out := new(strings.Builder)
    start := 0
    for i := 0; i < len(s); i++ {
        strChar := string(s[i])
        if strChar == " " {
            for j := i-1; j >= start; j-- {
                out.WriteString(string(s[j])) 
            }
            out.WriteString(" ")
            start = i+1
        }
    }
    
    for j := len(s)-1; j >= start; j--{
        out.WriteString(string(s[j]))
    }
    return out.String()
}