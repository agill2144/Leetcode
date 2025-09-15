func canBeTypedWords(text string, brokenLetters string) int {
    set := make([]bool, 26)
    for i := 0; i < len(brokenLetters); i++ {
        set[int(brokenLetters[i]-'a')] = true
    }
    count := 0
    i := 0
    for i < len(text) {
        ptr := i
        skip := false
        for ptr < len(text) && text[ptr] != ' ' {
            if set[int(text[ptr]-'a')] {skip = true}
            ptr++
        }
        if !skip {count++}
        i = ptr+1
    }
    return count
}