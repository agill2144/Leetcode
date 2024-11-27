func minAddToMakeValid(s string) int {
    extraClosing := 0
    open := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            open++
        } else if s[i] == ')' {
            if open == 0 {extraClosing++; continue}
            open--
        }
    }
    return open + extraClosing
}