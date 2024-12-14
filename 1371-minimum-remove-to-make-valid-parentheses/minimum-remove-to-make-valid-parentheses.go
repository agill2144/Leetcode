func minRemoveToMakeValid(s string) string {
    out := new(strings.Builder)
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
            out.WriteByte(s[i])
        } else if s[i] == ')' {
            if count > 0 {
                count--
                out.WriteByte(s[i])
            }
        } else {
            out.WriteByte(s[i])
        }
    }
    outStr := out.String()
    out.Reset()
    count = 0
    for i := len(outStr)-1; i >= 0; i-- {
        if outStr[i] == '(' {
            if count > 0 {
                count--
                out.WriteByte(outStr[i])
            }
        } else if outStr[i] == ')' {
            count++ 
            out.WriteByte(outStr[i])
        } else {
            out.WriteByte(outStr[i])
        }
    }
    outStr = out.String()
    out.Reset()
    for i := len(outStr)-1; i >= 0; i-- {
        out.WriteByte(outStr[i])
    } 
    return out.String()
}