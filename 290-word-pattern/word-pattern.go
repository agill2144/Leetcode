func wordPattern(pattern string, s string) bool {
    pToS := map[byte]string{}
    sToP := map[string]byte{}
    words := strings.Split(s, " ")
    if len(words) != len(pattern) {return false}
    for i := 0; i < len(words); i++ {
        char := pattern[i]
        word := words[i]
        val, ok := pToS[char]
        if ok && val != word {return false}
        if !ok {pToS[char]=word}

        val2, ok2 := sToP[word]
        if ok2 && val2 != char {return false}
        if !ok2 {sToP[word]=char} 
    }
    return true
}