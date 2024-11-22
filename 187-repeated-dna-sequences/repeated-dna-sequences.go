func findRepeatedDnaSequences(s string) []string {
    out := []string{}
    freq := map[string]int{}
    left := 0
    for i := 0; i < len(s); i++ {
        if i-left+1 == 10 {
            subStr := s[left:i+1]
            freq[subStr]++
            count := freq[subStr]
            if count == 2 {
                out = append(out, subStr)
            }
            left++
        }
    }
    return out
}