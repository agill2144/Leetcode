func checkValidString(s string) bool {
    open := 0
    close := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '*' {
            open++
        } else if s[i] == ')' {
            open--
            if open < 0 {return false}
        }

        if s[len(s)-1-i] == ')' || s[len(s)-1-i] == '*' {
            close++
        } else if s[len(s)-1-i] == '(' {
            close--
            if close < 0 {return false}
        }
    }
    return true
}