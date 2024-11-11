func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    words1 := strings.Split(sentence1, " ")
    words2 := strings.Split(sentence2, " ")
    w1L, w2L := 0,0
    for w1L < len(words1) && w2L < len(words2) && words1[w1L] == words2[w2L] {
        w1L++; w2L++
    }
    w1R, w2R := len(words1)-1,len(words2)-1
    for w1R >= 0 && w2R >= 0 && words1[w1R] == words2[w2R] {
        w1R--; w2R--
    }
    return w1L > w1R || w2L > w2R
}