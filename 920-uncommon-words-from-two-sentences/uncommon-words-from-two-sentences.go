func uncommonFromSentences(s1 string, s2 string) []string {
    s1Freq := map[string]int{}
    s1Words := strings.Split(s1, " ")
    for i := 0; i < len(s1Words); i++ {
        s1Freq[s1Words[i]]++
    }
    s2Freq := map[string]int{}
    s2Words := strings.Split(s2, " ")
    for i := 0; i < len(s2Words); i++ {
        s2Freq[s2Words[i]]++
    }
    out := []string{}
    for k, v := range s1Freq {
        _, ok := s2Freq[k]
        if v > 1 || ok {continue}
        out = append(out, k)
    }
    for k, v := range s2Freq {
        _, ok := s1Freq[k]
        if v > 1 || ok {continue}
        out = append(out, k)
    }
    return out
}