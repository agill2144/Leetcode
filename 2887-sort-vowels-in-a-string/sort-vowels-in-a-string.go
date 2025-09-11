func sortVowels(s string) string {
    set := map[byte]bool{
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }
    vs := []byte{}
    for i := 0; i < len(s); i++ {
        if set[s[i]] {vs = append(vs,s[i])}
    }
    sort.Slice(vs, func(i, j int)bool{
        return vs[i] < vs[j]
    })
    vPtr := 0
    out := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if set[s[i]] {
            out.WriteByte(vs[vPtr])
            vPtr++
        } else {
            out.WriteByte(s[i])
        }
    }
    return out.String()
}