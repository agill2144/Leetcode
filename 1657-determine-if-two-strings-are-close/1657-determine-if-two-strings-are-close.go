func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {return false}
    w1Total := 0
    w1Freq := make([]int, 26)
    w2Total := 0
    w2Freq := make([]int, 26)
    for i := 0; i < len(word1); i++ {w1Freq[int(word1[i]-'a')]++; w1Total++}
    for i := 0; i < len(word2); i++ {w2Freq[int(word2[i]-'a')]++; w2Total++}
    for i := 0; i < 26; i++ {
        if (w1Freq[i] == 0 && w2Freq[i] == 0) ||  (w1Freq[i] > 0 && w2Freq[i] > 0) {continue}
        return false
    }
    sort.Ints(w1Freq)
    sort.Ints(w2Freq)
    return reflect.DeepEqual(w1Freq,w2Freq)
}