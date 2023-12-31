func makeEqual(words []string) bool {
    if len(words) == 1 {return true}
    charCounts := map[byte]int{}
    for i := 0; i < len(words); i++ {
        word := words[i]
        for k := 0; k < len(word); k++ {
            charCounts[word[k]]++
        }
    }
    for _, v := range charCounts {
        if v % len(words) != 0 {return false}
    }    
    return true
}