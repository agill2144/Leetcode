func doesAliceWin(s string) bool {
    set := map[byte]bool{'a':true, 'e':true, 'i':true, 'o':true, 'u': true}
    count := 0
    for i := 0; i < len(s); i++ {
        if set[s[i]] {count++}
    }
    return count > 0
}