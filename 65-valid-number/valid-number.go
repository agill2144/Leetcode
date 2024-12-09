/*
    rules tldr
    - we must have seen atleast 1 digit
    - "+" or "-" are only allowed 2 times, and never allowed at the last position
        - at the beginning
        - right after "e" || "E"
    - "e" and "E" is not allowed
        - when we have not seen digit yet, cannot have a exponent without a base number
        - this is at the last position
        - we have already seen an e || E before


*/
func isNumber(s string) bool {
    n := len(s)
    // we can only have 1 decimal ever
    seenDecimal := false
    // we must have seen 1 digit
    seenDigit := false
    // we can only have 1 exponent
    seenExp := false
    start := 0
    // we can have a operator at the very beginning AND right after e||E AND we cannot end number with operator
    if s[start] == '+' || s[start] == '-' {start = 1}
    for i := start; i < n; i++ {
        char := s[i]
        if char == '+' || char == '-' {
            // cannot have a operator at the end
            if i == n-1 || (i-1 >= 0 && s[i-1] != 'e' && s[i-1] != 'E') {return false}
        } else if char == 'E' || char == 'e' {
            // we have seen an exp already or
            // we have not seen a digit yet, cannot have a e without a digit first, or
            // or this e is at the, invalid!
            if seenExp || !seenDigit || i == n-1 {return false}
            seenExp = true
        } else if char >= '0' && char <= '9' {
            seenDigit = true
        } else if char == '.' {
            // cannot have more than 1 decimal OR
            // cannot have a decimal be part of the exponent
            if seenDecimal || seenExp {return false}
            seenDecimal = true
        } else {
            // random char that is not valid "~", "a", "b", "?" etc..
            return false
        }
    }
    return seenDigit
}