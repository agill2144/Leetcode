func longestPalindrome(s string) int {
    set := map[byte]bool{}
    count := 0
    for i := 0; i < len(s); i++ {
        if set[s[i]] {
            count += 2
            delete(set, s[i])
        } else {
            set[s[i]] = true
        }
    }
    if len(set) != 0 {count++}
    return count
}