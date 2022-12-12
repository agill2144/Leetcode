/*
    appraoch: hashmap
    - we need to check whether the given words in a list are sorted
    - however with a custom order ( not using the traditional a..z order )
    - the order of characters are provided to us.
    - we can store the order of characters in a idx map
        - { $char: $idx }
    - then we can compare 2 words at a time and check whether
        - word1 should be placed before word2
        - how do we check that?
        - in a normal sort, how do we compare 2 strings?
        - word1 = work
        - word2 = word
        - we check character by character for both words
        - if characters match for both words, we have nothing to make a decision on
        - if characters do not match, then we have to make decision
            - should the character in word1 really come before character in word2 ?
        - likewise, we will do the same comparison; we will use the order idxMap to make our decision.
        - What if; we had
        - word1 = world
        - word2 = worn
        - as soon as our i pointer for comparing both characters goes out bound of bounds for either of the words, we are saying return true that they are sorted
        - but should world come before worm ? no, worn is smaller in length, so it must be placed before word1
        - even after comparing character by character
        - we need to check for len too.
        - if our character check came out true
        - then check if len of word1 is greater than word2
        - if it is, return false because for word1 to be placed before word2, it must be SMALLER in length
    
    time: 
    - o(order) + o(numberOfWords)xo(avgLenOfWords)
    - o(order) is constant - it will never exceed 26
    - so therefore o(nk) - n words, k avg len of each word
    
    space:
    - we allocated an extra idxMap, buts constant because there are only 26 characters that can be added as keys from order string

*/
func isAlienSorted(words []string, order string) bool {
    charIdxMap := map[byte]int{}
    for i := 0; i < len(order); i++ {
        charIdxMap[order[i]] = i
    }
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        if !isSorted(charIdxMap, word1, word2) {
            return false
        }
    }
    return true 
}

func isSorted(charIdxMap map[byte]int, word1, word2 string) bool {
    for i := 0; i < len(word1) && i < len(word2); i++ {
        word1Char := word1[i]
        word2Char := word2[i]
        if word1Char != word2Char {
            return charIdxMap[word1Char] < charIdxMap[word2Char]
        }
    }
    
    // if we got here, the characters are so far lexicographically sorted
    // but the smaller len must be word1 for word1 to be placed before word2
    if len(word1) > len(word2) {
        return false
    }
    
    return true
}