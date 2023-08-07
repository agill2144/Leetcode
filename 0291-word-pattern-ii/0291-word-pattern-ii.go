func wordPatternMatch(pattern string, s string) bool {
    
    var dfs func(pPtr int, sPtr int, patternToS map[byte]string, sToPattern map[string]byte) bool
    dfs = func(pPtr, sPtr int, patternToS map[byte]string, sToPattern map[string]byte) bool {
        // base
        if pPtr == len(pattern) && sPtr == len(s) {return true}
        if pPtr == len(pattern) || sPtr == len(s) {return false}
                
        // logic
        pChar := pattern[pPtr]
        for i := sPtr; i < len(s); i++ {

            subStr := string(s[sPtr:i+1])
            val, ok := patternToS[pChar]
            if ok && val != subStr {continue}
            val2, ok2 := sToPattern[subStr]
            if ok2 && val2 != pChar {continue}
            
            
            if !ok {
                patternToS[pChar] = subStr
                sToPattern[subStr] = pChar
            }
            
            if dfs(pPtr+1, i+1, patternToS, sToPattern) {
                return true
            } 
            
            if !ok {
                delete(patternToS, pChar)
                delete(sToPattern, subStr)
            }
        }
        return false
    }
    return dfs(0, 0, map[byte]string{}, map[string]byte{})
}