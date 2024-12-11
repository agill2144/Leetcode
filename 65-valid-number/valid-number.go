/*
    not allowed rules tldr
    - we must have seen a digit
    - "+" or "-" are only allowed 2 times
        - at the beginning
        - right after "e" || "E"
        - not allowed when at n-1 idx 
        - not allowe when prev char is not "e" or "E"
    - "e" and "E":
        - not allowed when we haven't seen a digit yet
        - not allowed when at n-1 idx 
        - not we have already seen an e || E before
    - "." is not allowed;
        - not allowed when we have already seen one before
        - not allowed after exp, ie; we have seen an exp before this alreay
    
    tc = o(s)
    sc = o(1)
*/
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