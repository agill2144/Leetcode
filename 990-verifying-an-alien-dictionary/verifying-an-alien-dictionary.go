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
        - but should world come before worn ? no, worn is smaller in length, so it must be placed before word1
        - even after comparing character by character
        - we need to check for len too.
        - if our character check came out true
        - then check if len of word1 is greater than word2
        - if it is, return false because for word1 to be placed before word2, it must be SMALLER in length

*/
func isAlienSorted(words []string, order string) bool {
    idx := map[byte]int{}
    for i := 0; i < len(order); i++ {idx[order[i]] = i}
    for i := 1; i < len(words); i++ {
        curr, prev := words[i], words[i-1]
        c, p := 0,0
        for c < len(curr) && p < len(prev) {
            if curr[c] == prev[c] {c++; p++; continue}
            cIdx := idx[curr[c]]
            pIdx := idx[prev[p]]
            if pIdx > cIdx {return false}
            break
        }

        // if p is still inbound and c went out of bounds
        // it means prev char is larger in size compared to curr
        // and curr should have been placed before prev ( smaller , then larger ) 
        if p < len(prev) && c == len(curr) {return false}
    }
    return true
}