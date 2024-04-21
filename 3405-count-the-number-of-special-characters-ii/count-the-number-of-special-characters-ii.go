// time = o(2n)  = o(n)
// space = o(n/2) for ansSet to de-dupe = o(n)
func numberOfSpecialChars(word string) int {
    idxMap := map[byte][]int{}  // first and last idx of each char
    for i := 0; i < len(word); i++ {
        _, ok := idxMap[word[i]]
        if !ok {
            idxMap[word[i]] = []int{i,i}
        } else {
            idxMap[word[i]][1] = i
        }
    }
    count := 0
    lowers := make([]bool, 26)
    uppers := make([]bool, 26)
    for i := 0; i < len(word); i++ {
        char := word[i]
        if char < 97 {continue}
        
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
        lastLowerIdx := idxMap[lowerCase][1]
        firstUpperIdx := idxMap[upperCase][0]
        
        // every lowercase occurrence appears before the first uppercase occurrence
        if lastLowerIdx > firstUpperIdx {continue}
        // inorder to not count dupe
        if !lowers[int(lowerCase-'a')] && !uppers[int(upperCase-'A')] {
            lowers[int(lowerCase-'a')] = true
            uppers[int(upperCase-'A')] = true
            count++
        }
    }
    return count
}