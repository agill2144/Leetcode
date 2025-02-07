func removeDuplicates(s string) string {
    st := []byte{}
    for i := 0; i < len(s); i++ {
        if len(st) == 0 {st = append(st, s[i]); continue}
        if st[len(st)-1] == s[i] {
            st = st[:len(st)-1]
        } else {
            st = append(st, s[i])
        }
    }
    res := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        res.WriteByte(st[i])
    } 
    return res.String()
}