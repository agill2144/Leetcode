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
            // cannot have a operator at the end
            if i == len(s)-1 {return false}
            if i-1 >= 0 && s[i-1] != 'e' && s[i-1] != 'E' {return false}
        } else if char == 'E' || char == 'e' {
            // we have seen an exp already or
            // we have not seen a digit yet, cannot have a e without a digit first, or
            // or this e is at the, invalid!
            if seenExp || !seenDigit || i == len(s)-1 {return false}
            seenExp = true
        } else if char >= '0' && char <= '9' {
            seenDigit = true
        } else if char == '.' {
            // cannot have more than 1 decimal OR
            // cannot have a decimal be part of the exponent
            if seenDecimal || seenExp {return false}
            seenDecimal = true
        } else {
            return false
        }
    }
    return seenDigit
}