func isValid(s string) bool {
    m := map[byte]byte{
        '}':'{',
        ')':'(',
        ']':'[',
    }
    st := []byte{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '[' || char == '{' || char == '(' {
            st = append(st, char)
        } else {
            if len(st) == 0 {return false}
            matchingPair := m[char]
            top := st[len(st)-1]
            st = st[:len(st)-1]
            if top != matchingPair {return false}
        }
    }
    if len(st) != 0 {return false}
    return true
}