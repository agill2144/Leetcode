func minLength(s string) int {
    st := []string{}
    for i := 0; i < len(s); i++ {
        char := string(s[i])
        if len(st) != 0 && (st[len(st)-1] + char == "AB" || st[len(st)-1] + char == "CD") {
            st = st[:len(st)-1]
        } else {
            st = append(st, char)
        }
    }
    return len(st)
}
/*

    CCCCDDDD

    
*/