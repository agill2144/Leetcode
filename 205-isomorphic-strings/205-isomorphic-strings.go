func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap := map[byte]byte{}
    tMap := map[byte]byte{}
    
    for i := 0; i < len(s); i++ {
        sChar := s[i]
        tChar := t[i]
        
        val, ok := sMap[sChar]
        if !ok {
            sMap[sChar] = tChar
        } else if val != tChar {
            return false
        }
        
        
        val, ok = tMap[tChar]
        if !ok {
            tMap[tChar] = sChar
        } else if val != sChar {
            return false
        }
        
    }
    return true
}