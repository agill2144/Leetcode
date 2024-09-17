func uncommonFromSentences(s1 string, s2 string) []string {
    freq := map[string]int{}
    s1Words := strings.Split(s1, " ")
    for i := 0; i < len(s1Words); i++ {
        freq[s1Words[i]]++
    }
    s2Words := strings.Split(s2, " ")
    for i := 0; i < len(s2Words); i++ {
        freq[s2Words[i]]++
    }
    out := []string{}
    for k, v := range freq {
       if v > 1 {continue}
       out = append(out, k)
    }
    return out
}