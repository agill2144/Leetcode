func minimumDeletions(s string) int {
    st := []byte{}
    count := 0
    for i := 0; i < len(s); i++ {
        if len(st) != 0 && st[len(st)-1] == 'b' && s[i] == 'a' {
            st = st[:len(st)-1]
            count++
        } else {
            st = append(st, s[i])
        }
    }
    return count
}