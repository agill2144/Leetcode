/*
    rules tldr
    - we must have seen atleast 1 digit
    - "+" or "-" are only allowed 2 times, and never allowed at the last position
        - at the beginning
        - right after "e" || "E"
        - so its not allowed when i == n-1 || prev char is not a e and not a E
    - "e" and "E" is not allowed
        - when we have not seen digit yet, cannot have a exponent without a base number
        - this is at the last position
        - we have already seen an e || E before
    - "." is not allowed;
        - when we have already seen a dot before
        - when we have seen an exp already (dots / decimals are not allowed in exponents )

*/
func isNumber(s string) bool {
    n := len(s)
    seenDigit := false
    seenExp := false
    seenDecimal := false
    start := 0
    if s[start] == '+' || s[start] == '-' {start++}
    for i := start ; i < n; i++ {
        char := s[i]
        if char == '+' || char == '-' {
            if i == n-1 ||
                (i-1 >= 0 && s[i-1] != 'e' && s[i-1] != 'E') {
                    return false
                }
        } else if char == 'e' || char == 'E' {
            if i == n-1 || seenExp || !seenDigit {return false}
            seenExp = true
        } else if char == '.' {
            if seenExp || seenDecimal {return false}
            seenDecimal = true
        } else if char >= '0' && char <= '9' {
            seenDigit = true
        } else {
            return false
        }
    }
    return seenDigit
}