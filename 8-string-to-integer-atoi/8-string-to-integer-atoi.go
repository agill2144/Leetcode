func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if s == "" { return 0 }
    firstChar := s[0]
    if !(firstChar == '-' || firstChar == '+') && !isDigit(firstChar) {
        return 0
    }
    isNegative := false
    if firstChar == '-' { isNegative = true }
    n := 0
    limit := math.MaxInt32/10
    for i := 0; i < len(s); i++ {
        char := s[i]
        if isDigit(char) {
            charInt := int(char-'0')
            if !isNegative {
                if n == limit && charInt >= 7 { return math.MaxInt32 }
                if n > limit {return math.MaxInt32 }
            } else {
                if n == limit && charInt >= 8 { return math.MinInt32 }
                if n > limit {return math.MinInt32 }
            }
            n = n * 10 + charInt
        } else if i != 0 {
            break
        }
    }
    if isNegative { n *= -1 }
    return n
}

func isDigit(c byte) bool {
    return c >= '0' && c <= '9'
}