// wtf is this question.... go learn some fucking english if you are parsing this line..
func diStringMatch(s string) []int {
    n := len(s)
    d := n
    i := 0
    out := []int{}
    for j := 0; j < len(s); j++ {
        if s[j] == 'I' {
            out = append(out,i)
            i++
        } else {
            out = append(out, d)
            d--
        }
    }
    if d >= 0 {out = append(out, d)} else { out = append(out, i)}
    return out
}