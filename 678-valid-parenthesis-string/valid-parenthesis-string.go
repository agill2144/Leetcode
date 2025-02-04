func checkValidString(s string) bool {
    open := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '*' || s[i] == '(' {
            open++
        } else if s[i] == ')' {
            open--
            if open < 0 {return false}
        }
    }
    close := 0
    for i := len(s)-1; i >= 0; i-- {
        if s[i] == ')' || s[i] == '*' {
            close++
        } else if s[i] == '(' {
            close--
            if close < 0 {return false}
        }
    }
    return true
}