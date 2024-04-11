func removeDuplicateLetters(s string) string {
    lastIdx := map[byte]int{}
    for i := 0; i < len(s); i++ {lastIdx[s[i]] = i}
    set := map[byte]struct{}{}
    st := []byte{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        _, ok := set[char]
        if !ok {
            // The result will be 0 if a == b, -1 if a < b, and +1 if a > b
            for len(st) != 0 && (strings.Compare(string(char),string(st[len(st)-1])) == 0 || strings.Compare(string(char),string(st[len(st)-1])) == -1) && lastIdx[st[len(st)-1]] > i {
                delete(set, st[len(st)-1])
                st = st[:len(st)-1]
            }

            set[char] = struct{}{}
            st = append(st, char)
        }
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}