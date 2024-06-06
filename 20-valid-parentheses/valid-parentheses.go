func isValid(s string) bool {
    valid := map[byte]byte{
        '}':'{',']':'[',')':'(',
    }
    st := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            st = append(st, s[i])
        } else {
            if len(st) == 0 {return false}
            validOpen := valid[s[i]]
            if st[len(st)-1] != validOpen {return false}
            st = st[:len(st)-1]
        }
    }
    return len(st) == 0 
}