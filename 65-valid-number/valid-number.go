func isNumber(s string) bool {
    seenDigit := false
    seenExp := false
    seenDec := false
    n := len(s)
    start := 0
    if s[start] == '+' || s[start] == '-' {start++}
    for i := start; i < n; i++ {
        char := s[i]
        if char == 'e' || char == 'E' {
            if !seenDigit || seenExp || i == n-1 {return false}
            seenExp = true
        } else if char == '+' || char == '-' {
            if (s[i-1] != 'e' && s[i-1] != 'E') || i == n-1 {return false}
        } else if char == '.' {
            if seenDec || seenExp {return false}
            seenDec = true
        } else if char >= '0' && char <= '9' {
            seenDigit = true
        } else {
            return false
        }
    }
    return seenDigit
}