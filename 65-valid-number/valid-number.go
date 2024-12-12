func isNumber(s string) bool {
    seenDecimal := false
    seenDigit := false
    seenExp := false
    start := 0
    if s[start] == '+' || s[start] == '-' {
        start++
    }
    n := len(s)
    for i := start; i < n; i++ {
        char := s[i]
        if char == '+' || char == '-' {
            // 12e+2
            if (s[i-1] != 'e' && s[i-1] != 'E') || i == n-1 {return false}            
        } else if char == 'e' || char == 'E' {
            if !seenDigit || seenExp || i==n-1 {return false}
            seenExp = true
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