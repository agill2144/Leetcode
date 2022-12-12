func wordPattern(pattern string, s string) bool {
    pMap := map[byte]string{}
    sMap := map[string]byte{}
    sSplit := strings.Split(s, " ")
    if len(sSplit) != len(pattern) {return false} 
    for i := 0; i < len(pattern); i++ {
        pChar := pattern[i]
        sWord :=  sSplit[i]
        
        val, ok := pMap[pChar]
        if !ok {
            pMap[pChar] = sWord
        } else if val != sWord {
            return false
        }
        
        v, ok := sMap[sWord]
        if !ok {
            sMap[sWord] = pChar
        } else if v != pChar {
            return false
        }
    }
    return true
}