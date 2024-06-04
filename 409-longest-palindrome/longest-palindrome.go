func longestPalindrome(s string) int {
    chars := map[byte]struct{}{}
    total := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        _, ok := chars[char]
        if ok {
            total += 2
            delete(chars, char)
        } else {
            chars[char] = struct{}{}
        }
    }
    if len(chars) != 0 {
        total++
    }
    return total
}