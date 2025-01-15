func checkValidString(s string) bool {
    n := len(s)
    if n == 0 {return true}
    open := 0
    close := 0
    for i := 0; i < n; i++ {
        // handle using asterik as open
        if s[i] == '(' || s[i] == '*' {
            open++
        } else {
            open--
            if open < 0 {return false}
        }

        // simulatenously, traverse from righ to left
        // handle using asterik as close
        if s[n-1-i] == ')' || s[n-1-i] == '*' {
            close++
        } else {
            close--
            if close < 0 {return false}
        }

    }
    return true
}