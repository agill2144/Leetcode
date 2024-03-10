func customSortString(order string, s string) string {
    sFreq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        sFreq[s[i]]++
    }
    ptr := 0
    out := new(strings.Builder)
    for ptr < len(order) {
        char := order[ptr]
        times := sFreq[char]
        for times > 0 {
            out.WriteByte(char)
            times--
        } 
        delete(sFreq, char)
        ptr++
    }
    for k, v := range sFreq {
        for v > 0 {
            out.WriteByte(k)
            v--
        }
    }
    return out.String()
}