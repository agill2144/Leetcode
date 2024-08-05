func vowelStrings(words []string, left int, right int) int {
    count := 0
    for i := left; i <= right; i++ {
        if isVowelWord(words[i]) {count++}
    }
    return count
}

func isVowelWord(word string) bool {
    vowels := map[byte]struct{}{
        'a': {}, 'e': {}, 'i':{}, 'o':{}, 'u':{},
    }
    _, ok := vowels[word[0]]
    _, ok2 := vowels[word[len(word)-1]]
    return ok && ok2
}