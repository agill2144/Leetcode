func numberOfSpecialChars(word string) int {
    idxMap := map[byte][]int{}
    for i := 0; i < len(word); i++ {
        idxMap[word[i]] = append(idxMap[word[i]], i)
    }
    ansSet := map[string]struct{}{}
    for i := 0; i < len(word); i++ {
        char := word[i]
        if unicode.IsUpper(rune(char)) {continue}
        
        // char is lower, look for its upper pair
        lowerCase := char
        // lowercase "a" ascii value = 97
        // uppercase "A" ascii value = 65
        // the difference is 32
        // therefore if we have a lowercase
        // and we want its upper case ascii value
        // lowercaseAscII - 32
        upperCase := lowerCase-32
        if len(idxMap[upperCase]) == 0 {continue}
        lastLowerIdx := len(idxMap[lowerCase])-1
        
        // every lowercase occurrence appears before the first uppercase occurrence
        if idxMap[lowerCase][lastLowerIdx] > idxMap[upperCase][0] {continue}
        pair := fmt.Sprintf("%v-%v", lowerCase, upperCase)
        ansSet[pair] = struct{}{}
    }
    return len(ansSet)
}