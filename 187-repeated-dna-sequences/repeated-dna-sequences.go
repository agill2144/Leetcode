func findRepeatedDnaSequences(s string) []string {
    freq := map[string]int{}
    left := 0
    for i := 0; i < len(s); i++ {
        if i-left+1 == 10 {
            freq[s[left:i+1]]++
            left++
        }
    }
    out := []string{}
    for k, v := range freq {
        if v >= 2 {out = append(out, k)}
    }
    return out
}