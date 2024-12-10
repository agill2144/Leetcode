func isNumber(s string) bool {
    n := len(s)
    seenExp := false
    seenDigit := false
    seenDecimal := false
    start := 0
    if s[start] == '+' || s[start] == '-' {start++}
    for i := start; i < len(s); i++ {
        char := s[i]
        if char == 'e' || char == 'E' {
            if !seenDigit || i == n-1 || seenExp {return false}
            seenExp = true
        } else if char == '+' || char == '-' {
            if !seenDigit || i == n-1 || (s[i-1] != 'e' && s[i-1] != 'E'){return false}
        } else if char == '.' {
            if seenDecimal || seenExp {return false}
            seenDecimal = true
        } else if char >= '0' && char <= '9' {
            seenDigit = true
        } else {
            return false
        }
    }
    return seenDigit
}