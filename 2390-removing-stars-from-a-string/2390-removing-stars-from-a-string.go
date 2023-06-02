func removeStars(s string) string {
    st := []byte{}
    for i := 0; i < len(s); i++ {
        // *
        if len(st) != 0 && s[i] == '*' {
            st = st[:len(st)-1]
        } else {
            st = append(st, s[i])
        }
    }
    
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}