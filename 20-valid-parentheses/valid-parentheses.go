func isValid(s string) bool {
    validOpens := map[byte]byte{
        ')': '(', ']': '[','}':'{',
    }
    st := []byte{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '(' || char == '[' || char == '{' {
            st = append(st, char)
        } else {
            if len(st) == 0 {
                return false
            }
            if validOpens[char] != st[len(st)-1] {
                return false
            }
            st = st[:len(st)-1]
        }
    }
    if len(st) != 0 {return false}
    return true
}