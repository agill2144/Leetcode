func customSortString(order string, s string) string {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {freq[s[i]]++}
    out := new(strings.Builder)
    for i := 0; i < len(order); i++ {
        for freq[order[i]] > 0 {
            out.WriteByte(order[i])
            freq[order[i]]--
        }
    }
    for i := 0; i < len(s); i++ {
        for freq[s[i]] > 0 {
            out.WriteByte(s[i])
            freq[s[i]]--
        }        
    }
    return out.String()
}