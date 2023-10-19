func backspaceCompare(s string, t string) bool {
    newS := ""
    skipCount := 0
    
    for i := len(s)-1; i>=0 ; i-- {
        char := string(s[i])
        if char == "#" {
            skipCount++
        } else if skipCount != 0 {
            skipCount--
        } else {
            newS += char
        }
    }
    
    newT := ""
    skipCountT := 0
    
    for i := len(t)-1; i>=0 ; i-- {
        char := string(t[i])
        if char == "#" {
            skipCountT++
        } else if skipCountT != 0 {
            skipCountT--
        } else {
            newT += char
        }
    }
    return newS == newT
}