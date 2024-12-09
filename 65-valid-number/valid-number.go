func isNumber(s string) bool {
    // we can only have 1 decimal ever
    seenDecimal := false
    // we must have seen 1 digit
    seenDigit := false
    // we can only have 1 exponent
    seenExp := false
    start := 0
    // we can have a operator at the very beginning AND right after e||E AND we cannot end number with operator
    if s[start] == '+' || s[start] == '-' {start = 1}
    for i := start; i < len(s); i++ {
        char := s[i]
        if char == '+' || char == '-' {
            if i == len(s)-1 {return false}
            if i-1 >= 0 && (s[i-1] == 'e' || s[i-1] == 'E') && i != len(s)-1 {continue}
            return false
        } else if char == 'E' || char == 'e' {
            if !seenDigit || i == len(s)-1 {return false}
            if seenExp {return false}
            seenExp = true
        } else if char >= '0' && char <= '9' {
            seenDigit = true
        } else if char == '.' {
            if seenDecimal || seenExp {return false}
            seenDecimal = true
        } else {
            return false
        }
    }
    return seenDigit
}