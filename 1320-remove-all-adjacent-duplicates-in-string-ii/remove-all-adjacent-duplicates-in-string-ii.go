func removeDuplicates(s string, k int) string {
    st := []*node{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        if len(st) != 0{
            top := st[len(st)-1]
            if char == top.val {
                if top.count + 1 == k {
                    st = st[:len(st)-1]
                } else {
                    st[len(st)-1].count++
                }
            } else {
                st = append(st, &node{val: char, count:1})
            }
        } else {
            st = append(st, &node{val: char, count:1})
        }
    }
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        char := st[i].val
        times := st[i].count
        for times != 0 {
            out.WriteByte(char)
            times--
        }
    }
    return out.String()
}

type node struct {
    val byte
    count int
}