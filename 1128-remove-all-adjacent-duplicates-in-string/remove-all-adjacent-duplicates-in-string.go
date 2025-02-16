func removeDuplicates(s string) string {
    st := []byte{}
    for i := 0; i < len(s); i++ {
        pushCurr := true
        for len(st) > 0 && st[len(st)-1] == s[i] {
            st = st[:len(st)-1]
            pushCurr = false
        }
        if pushCurr {
            st = append(st, s[i])
        }
    }
    res := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        res.WriteByte(st[i])
    }
    return res.String()
}