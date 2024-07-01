func isValid(s string) bool {
    valids := map[byte]byte{
        '}':'{',']':'[',')':'(',
    }
    st := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            st = append(st, s[i])
        } else {
            if len(st) == 0 {return false}
            top := st[len(st)-1]
            st = st[:len(st)-1]
            val := valids[s[i]]
            if val != top {return false}
        }
    }
    return len(st) == 0
}