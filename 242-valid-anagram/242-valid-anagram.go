
// map solution
// time: o(s+t)
// space: o(1) -- there are only 26 alphabets
func isAnagram(s string, t string) bool {
    // make sure s is the smaller string
    if len(t) < len(s) {
        return isAnagram(t,s)
    }
    
    sMap := map[string]int{}
    for _, char := range s{
        sMap[string(char)]++
    }
    
    for _, char := range t {
        tChar := string(char)
        val, ok := sMap[tChar]
        if !ok {
            return false
        } else if val == 0 {
            return false
        }
        sMap[tChar]--
    }
    return true
}



// sort both strings and return equality check
// time: o(slogs) + o(tlogt) + split and join times for both s and t
// space: o(t) + o(s) for the split array 
// func isAnagram(s string, t string) bool {
//     news := strings.Split(s, "")
//     sort.Strings(news)
//     s = strings.Join(news, "")
//     newt := strings.Split(t, "")
//     sort.Strings(newt)
//     t = strings.Join(newt, "")
//     return s == t
// }


// prime product hash -- Does not work because the product becomes +Inf for larger words
// var primes = []float64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}
// func isAnagram(s string, t string) bool {
//     sHash := hash(s)
//     tHash := hash(t)
//     fmt.Println(sHash, tHash)
//     return sHash == tHash
// }

// func hash(word string) float64 {
//     var hash float64 = 1
//     for _, char := range word {
//         hash *= primes[char-'a']
//     }
//     return hash
// }

