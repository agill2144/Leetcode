func wordPattern(pattern string, s string) bool {
    sList := strings.Split(s, " ")
    if len(sList) != len(pattern) {return false}
    pToS := map[byte]string{}
    sToP := map[string]byte{}
    pPtr := 0; sPtr := 0
    for pPtr < len(pattern) && sPtr < len(s) {
        pchar := pattern[pPtr]
        sWord := sList[sPtr]
        val, ok := pToS[pchar]
        if ok && val != sWord {return false}
        pToS[pchar] = sWord

        val2, ok2 := sToP[sWord]
        if ok2 && val2 != pchar {return false}
        sToP[sWord]=pchar
        sPtr++; pPtr++
    }
    return true
}