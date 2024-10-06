func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    sent1 := strings.Split(sentence1, " ")
    sent2 := strings.Split(sentence2, " ")
    s1, s2 := 0,0
    for s1 < len(sent1) && s2 < len(sent2) {
        s1Word := sent1[s1]
        s2Word := sent2[s2]
        if s1Word == s2Word {
            s1++
            s2++
            continue
        }
        break
    }
    e1, e2 := len(sent1)-1, len(sent2)-1
    for e1 >= 0 && e2 >= 0 {
        s1Word := sent1[e1]
        s2Word := sent2[e2]
        if s1Word == s2Word {
            e1--
            e2--
            continue
        }
        break

    }
    return e1 < s2 || e2 < s1
}