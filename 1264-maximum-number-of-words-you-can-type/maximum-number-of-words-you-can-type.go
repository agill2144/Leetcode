func canBeTypedWords(text string, brokenLetters string) int {
    bkn := make([]bool, 26)
    for i := 0; i < len(brokenLetters); i++ {
        bkn[int(brokenLetters[i]-'a')] = true
    }
    count := 0
    skipCurr := false
    for i := 0; i < len(text); i++ {
        if text[i] == ' ' {
            if !skipCurr { count++ }
            skipCurr = false
        } else {
            if bkn[int(text[i]-'a')] {skipCurr = true}
        }
    }
    if !skipCurr {count++}
    return count
}