func vowelStrings(words []string, left int, right int) int {
    count := 0
    for i := left; i <= right; i++ {
        if isVowelWord(words[i]) {count++}
    }
    // time = o(right-left+1)
    return count
}

func isVowelWord(word string) bool {
    vowels := map[byte]bool{
        'a': true, 'e': true, 'i':true, 'o':true, 'u':true,
    }
    return vowels[word[0]] && vowels[word[len(word)-1]]
}