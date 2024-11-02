func isCircularSentence(sentence string) bool {
    for i := 0; i < len(sentence); i++ {
        if sentence[i] == ' ' {
            if sentence[i-1] != sentence[i+1] {return false}
        }
    }
    return sentence[0] == sentence[len(sentence)-1]
}