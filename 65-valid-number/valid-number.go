func isNumber(s string) bool {
    seenDigit := false
    seenDecimal := false
    seenExp := false
    start := 0
    n := len(s)
    if s[start] == '+' || s[start] == '-' {start++}
    for i := start; i < len(s); i++ {
        if s[i] == '+' || s[i] == '-' {
            if i == n-1 || (s[i-1] != 'e' && s[i-1] != 'E') {return false}             
        } else if s[i] == 'e' || s[i] == 'E' {
            if !seenDigit || seenExp || i == n-1 {return false}
            seenExp = true
        } else if s[i] == '.' {
            if seenExp || seenDecimal {return false}
            seenDecimal = true
        } else if s[i] >= '0' && s[i] <= '9' {
            seenDigit = true
        } else {
            return false
        }
    }
    return seenDigit
}