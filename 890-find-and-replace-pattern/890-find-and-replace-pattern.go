func findAndReplacePattern(words []string, pattern string) []string {
    result := []string{}
    
    for i := 0; i < len(words); i++ {
        word := words[i]
        wp := map[byte]byte{}
        pw := map[byte]byte{}
        
        reachedEnd := true
        for j := 0; j < len(word); j++ {
            wChar := word[j]
            pChar := pattern[j]
            
            val, ok := wp[wChar]
            if !ok {
                wp[wChar] = pChar
            } else if val != pChar {
                reachedEnd = false
                break
            }
            
            val, ok = pw[pChar]
            if !ok {
                pw[pChar] = wChar
            } else if val != wChar {
                reachedEnd = false
                break
            }
        }
        
        if reachedEnd { result = append(result, string(word))}
    }
    return result
}