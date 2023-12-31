func maxLengthBetweenEqualCharacters(s string) int {
    longest := math.MinInt64
    for i := 0; i < len(s); i++ {
        for j := i+1; j < len(s); j++ {
            if s[j] == s[i] {
                if (j-i+1)-2 >= longest {
                    longest = (j-i+1)-2
                    
                }
            }
        }
    }
    if longest == math.MinInt64 {
        longest = -1
    }
    return longest
}