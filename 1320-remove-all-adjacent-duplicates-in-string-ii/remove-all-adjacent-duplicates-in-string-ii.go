func removeDuplicates(s string, k int) string {
    type stNode struct {
        char byte
        freq int
    }
    st := []*stNode{}
    for i := 0; i < len(s); i++ {
        curr := s[i]
        if len(st) != 0 && curr == st[len(st)-1].char {
            st[len(st)-1].freq++
            if st[len(st)-1].freq == k {
                st = st[:len(st)-1]
            }
            continue
        }

        st = append(st, &stNode{s[i], 1})
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        curr := st[i]
        for t := 0; t < curr.freq; t++{
            out.WriteByte(curr.char)
        }
    }
    return out.String()
}