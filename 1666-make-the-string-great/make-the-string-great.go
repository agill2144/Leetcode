func makeGood(s string) string {
    if len(s) <= 1 {return s}
    st := []byte{}
    n := len(s)
    for i := 0; i < n; i++ {
        char := s[i]
        if len(st) == 0 {
            st = append(st, char)
        } else {
            top := st[len(st)-1]
            if strings.ToLower(string(top)) == strings.ToLower(string(char)) {
                topIsUpper := unicode.IsUpper(rune(top))
                currIsUpper := unicode.IsUpper(rune(char))
                if (topIsUpper && !currIsUpper) || (!topIsUpper && currIsUpper) {
                    st = st[:len(st)-1]
                    continue
                }
            }
            st = append(st, char)
        }
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}