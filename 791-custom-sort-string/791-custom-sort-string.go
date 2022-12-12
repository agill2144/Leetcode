/*
    s must be in the same form as order
    - FreqMap of s
    - loop over order
    - For each char in order, get all chars of this char and append to output string
    - while decrementing the count of values in freqMap as they are done
    - finally append whatever remaining is left in the freqMap to output string

    time: o(len(s)) + o(len(order))
    space: o(len(s))
*/

// func customSortString(order string, s string) string {
//     freqMap := map[string]int{}
//     for i := 0; i < len(s); i++ {
//         freqMap[string(s[i])]++
//     }
//     out := new(strings.Builder)
//     for i := 0; i < len(order); i++ {
//         char := string(order[i])
//         numTimes := freqMap[char] // if key is not found, the value will be 0 because default int value is 0
//         for j := 0; j < numTimes; j++{
//             out.WriteString(char)
//         }
//         delete(freqMap, char)
//     }
    
//     // whatever is left over, append
//     for key, val := range freqMap {
//         for i := 0; i < val; i++ {
//             out.WriteString(key)
//         }
//     }
//     return out.String()
// }


// use char array
func customSortString(order string, s string) string { 
    charFreq := [26]int{}
    for i := 0; i < len(s); i++ {
        charFreq[s[i]-'a']++
    }
    out := new(strings.Builder)
    for i := 0; i < len(order); i++ {
        idxOfChar := order[i]-'a'
        numTimes := charFreq[idxOfChar]
        char := string('a'+idxOfChar)
        for j := 0; j < numTimes;j++ {
            out.WriteString(char)
            charFreq[idxOfChar]--
        }
    }
        
    for i := 0; i < 26; i++ {
        char := string('a'+i)
        for j := 0; j < charFreq[i];j++ {
            out.WriteString(char)
        }
    }
    
    return out.String()
}