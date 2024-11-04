func makeFancyString(s string) string {
    res := &strings.Builder{}
    count := 0
    for i := 0; i < len(s); i++ {
        prev := s[i]
        curr := s[i]
        if i > 0 { prev = s[i-1]}
        if prev == curr {
            count++
        } else {
            count = 1
        }
        if count < 3 {
            res.WriteByte(curr)
        }
    }
    return res.String()
}