func numberOfSpecialSubstrings(s string) int {
    count := 0
    win := map[byte]int{}
    left := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        _, ok := win[char]
        for ok {
            leftChar := s[left]
            win[leftChar]--
            if win[leftChar] == 0 {
                delete(win, leftChar)
            }
            left++
            _, ok = win[char]
        }
        win[char]++
        count += i-left+1
    }
    return count
}