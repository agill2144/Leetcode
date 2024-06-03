func appendCharacters(s string, t string) int {
    sPtr, tPtr := 0,0

    for sPtr < len(s) && tPtr < len(t) {
        if s[sPtr] == t[tPtr] {
            sPtr++
            tPtr++
        } else {
            for sPtr < len(s) && s[sPtr] != t[tPtr] {
                sPtr++
            }
            if sPtr == len(s) {break}
        }
    }
    if tPtr >= len(t) {return 0}
    return len(t) - tPtr
    
}