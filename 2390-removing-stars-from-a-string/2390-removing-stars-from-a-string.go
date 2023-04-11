func removeStars(s string) string {
    st := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            // pop top
            if len(st) != 0 {st = st[:len(st)-1]}
        } else {
            st = append(st, s[i])
        }
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {out.WriteByte(st[i])}
    return out.String()
}
