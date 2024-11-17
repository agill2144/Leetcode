func minRemoveToMakeValid(s string) string {
    n := len(s)
    var open byte = '('
    var close byte = ')'
    count := 0
    tmp := new(strings.Builder)
    for i := 0; i < n; i++ {
        char := s[i]
        if char == open {
            count++
            tmp.WriteByte(char)
        } else if char == close {
            if count > 0 {count--; tmp.WriteByte(char)}
        } else {
            tmp.WriteByte(char)
        }
    }
    tmpStr := tmp.String()
    open, close = close, open
    count = 0
    tmp.Reset()
    for i := len(tmpStr)-1; i >= 0; i-- {
        char := tmpStr[i]
        if char == open {
            count++
            tmp.WriteByte(char)
        } else if char == close {
            if count > 0 {count--; tmp.WriteByte(char)}
        } else {
            tmp.WriteByte(char)
        }
    }
    tmpStr = tmp.String()
    tmp.Reset()
    for i := len(tmpStr)-1; i >= 0; i-- {
        tmp.WriteByte(tmpStr[i])
    }
    return tmp.String()
}