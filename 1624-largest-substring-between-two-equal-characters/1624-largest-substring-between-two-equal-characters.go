func maxLengthBetweenEqualCharacters(s string) int {
    longest := math.MinInt64
    charIdx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        idx, ok := charIdx[char]
        if ok {
            if (i-idx+1)-2 > longest {
                longest = (i-idx+1)-2
            }
        } else {
            charIdx[char] = i
        }
    }
    if longest == math.MinInt64 {return -1}
    return longest
}