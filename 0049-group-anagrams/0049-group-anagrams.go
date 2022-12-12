func groupAnagrams(strs []string) [][]string {
    m := map[float64][]string{}
    for i := 0; i < len(strs); i++ { // o(n)
        word := strs[i]
        primeProduct := hash(word) // o(k) , k is len of word
        m[primeProduct] = append(m[primeProduct], word)
    }
    // time so far: o(nk), space: o(n)
    
    out := [][]string{}
    for _, v := range m { // o(n) - worse case there are no anagrams
        out = append(out, v)
    }
    // time : o(nk) + o(n) = o(nk)
    // space: o(n)
    return out
}

var prime []float64 = []float64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func hash(word string) float64 {
    prod := 1.0
    for i := 0; i < len(word); i++ {
        // get idx of char given we have a char/byte = byte-'a'
        primeIdxOfThisChar := word[i]-'a'
        prod *= prime[primeIdxOfThisChar]
    }
    
    return prod
}

// func groupAnagrams(strs []string) [][]string {
    
//     if strs == nil || len(strs) == 0 {
//         return nil
//     }
    
//     m := map[string][]string{}
    
//     // o(n) -- n here is the number of words in strs array
//     for _, str := range strs {
//         strSlice := strings.Split(str,"")
//         // sort takes o(klogk) - k is the len of the current string being sorted
//         sort.Strings(strSlice)
//         sortedStr := strings.Join(strSlice, "")
//         m[sortedStr] = append(m[sortedStr], str)
//     }
//     // time of above: o(n) x o(klogk) = o(nklogk)

    
//     out := [][]string{}
//     // o(n) -- worse case if all words are distinct (i.e cannot be groupped into anagrams)
//     // then we will end up with our map being the same size as the input arr - therefore o(n) time
//     // space could also be o(n) if the above is true for worse case
//     for _, v := range m {
//         out = append(out, v)
//     }
    
//     // o(nklogk) + o(n)
//     // drop the non-dominant term
//     // time: o(nKlogK) -- where n is the array len and K is the len of each word in array
//     return out
// }


