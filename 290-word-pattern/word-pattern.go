func wordPattern(pattern string, s string) bool {
    ps := map[string]string{}
    sp := map[string]string{}
    sList := strings.Split(s, " ")
    if len(sList) != len(pattern) {return false}
    for i := 0; i < len(pattern); i++ {
        p := string(pattern[i])
        sWord := sList[i]
        
        psVal, ok := ps[p]
        if !ok {
            ps[p] = sWord
        } else if psVal != sWord {return false}

        spVal, ok2 := sp[sWord]
        if !ok2 {
            sp[sWord] = p
        } else if spVal != p {return false}
    }
    return true
}