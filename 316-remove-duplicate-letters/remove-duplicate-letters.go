func removeDuplicateLetters(s string) string {
    lastIdx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        lastIdx[s[i]] = i
    }
    inStack := map[byte]struct{}{}
    st := []byte{}
    for i := 0; i < len(s); i++ {
        curr := s[i]
        _, ok := inStack[curr]
        for len(st) != 0 && curr <= st[len(st)-1] && lastIdx[st[len(st)-1]] > i && !ok {
            delete(inStack, st[len(st)-1])
            st = st[:len(st)-1]
        }
        if !ok {
            st = append(st, curr)
            inStack[curr] = struct{}{}
        }
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}