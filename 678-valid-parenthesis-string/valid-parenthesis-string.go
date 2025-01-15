func checkValidString(s string) bool {
    n := len(s)
    if n == 0 {return true}
    open := 0
    for i := 0; i < n; i++ {
        if s[i] == '(' || s[i] == '*' {
            open++
        } else {
            open--
            if open < 0 {return false}
        }
    }
    close := 0
    for i := n-1; i >= 0; i-- {
        if s[i] == ')' || s[i] == '*' {
            close++
        } else {
            close--
            if close < 0 {return false}
        }
    }
    return true
}