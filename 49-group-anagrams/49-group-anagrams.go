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




func groupAnagrams(strs []string) [][]string {
    m := map[float64][]string{}
    for _, s := range strs {
        k := hash(s)
        m[k] = append(m[k], s)
    }
    out := [][]string{}
    for _, v := range m {
        out = append(out, v)
    }
    return out
}

var primes = []float64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}

// o(k) -- where k is the len of the string word
func hash(word string) float64 {
    var hash float64 = 1
    for _, char := range word {
        hash *= primes[char-'a']
    }
    return hash
}