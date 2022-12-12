func wordPattern(pattern string, s string) bool {
    sSlice := strings.Split(s, " ")
    if len(sSlice) != len(pattern) {
        return false
    }
    
    wordToPattern := map[string]string{}
    patternToWord := map[string]string{}
    
    for i := 0; i < len(sSlice); i++ {
        word := sSlice[i]
        p := string(pattern[i])
        
        val, ok := wordToPattern[word]
        if ok && val != p {
            return false
        }
        wordToPattern[word] = p
        
        val, ok = patternToWord[p]
        if ok && val != word {
            return false
        }
        patternToWord[p] = word
    }
    return true
}