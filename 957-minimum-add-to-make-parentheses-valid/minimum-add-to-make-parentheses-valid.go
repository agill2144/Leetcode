func minAddToMakeValid(s string) int {
    countOpen := 0
    invalidClosing := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            countOpen++
        } else if s[i] == ')' {
            if countOpen > 0 {
                countOpen--
            } else {
                invalidClosing++
            }
        }
    }
    return countOpen + invalidClosing
}