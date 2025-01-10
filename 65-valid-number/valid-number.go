func isNumber(s string) bool {
    seenExp := false
    seenDec := false
    seenDig := false
    start := 0
    if s[start] == '+' || s[start] == '-' {start++}
    n := len(s)
    for i := start; i < n; i++ {
        if s[i] >= '0' && s[i] <= '9' {
            seenDig = true
        } else if s[i] == '.' {
            if seenExp || seenDec {return false}
            seenDec = true
        } else if s[i] == 'e' || s[i] == 'E' {
            if !seenDig || seenExp || i == n-1 {return false}
            seenExp = true
        } else if s[i] == '+' || s[i] == '-' {
            if i == n-1 || (i-1 >= 0 && (s[i-1] !='e' && s[i-1] != 'E')) {return false}
        } else {
            return false
        }
    }
    return seenDig
}