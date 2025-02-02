func minAddToMakeValid(s string) int {
    open := 0
    extraClosing := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            open++
        } else if s[i] == ')' {
            if open > 0 {
                open--
            } else {
                extraClosing++
            }
        }
    }
    return open + extraClosing
}