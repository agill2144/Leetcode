func longestPalindrome(s string) int {
    set := map[byte]bool{}
    size := 0
    for i := 0; i < len(s); i++ {
        if set[s[i]] {
            delete(set, s[i])
            size += 2
        } else {
            set[s[i]] = true
        }
    }
    if len(set) != 0 {size++}
    return size
}