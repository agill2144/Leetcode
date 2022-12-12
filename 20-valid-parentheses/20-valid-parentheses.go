func isValid(s string) bool {
    st := []byte{}
    m := map[byte]byte{
        '}':'{',
        ']':'[',
        ')':'(',
    }
    
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '(' || char == '[' || char == '{' {
            st = append(st, char)
        } else {
            if len(st) == 0 {return false}
            top := st[len(st)-1]
            st = st[:len(st)-1]
            if m[char] != top {
                return false
            }
        }
    }
    return len(st) == 0
}