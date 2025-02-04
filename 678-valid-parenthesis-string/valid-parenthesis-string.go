func checkValidString(s string) bool {
    open := 0
    close := 0
    n := len(s)
    for i := 0; i < len(s); i++ {
        if s[i] == '*' || s[i] == '(' {
            open++
        } else if s[i] == ')' {
            open--
            if open < 0 {return false}
        }

        if s[n-1-i] == '*' || s[n-1-i] == ')' {
            close++
        } else if s[n-1-i] == '(' {
            close--
            if close < 0 { return false }
        }
    }
    return true
}