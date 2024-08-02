func characterReplacement(s string, k int) int {
    freq := map[byte]int{}
    maxRep := 0
    left := 0
    maxWin := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        maxRep = max(maxRep, freq[s[i]])
        nonRepeating := (i-left+1) - maxRep
        if nonRepeating <= k {
            maxWin = max(maxWin, i-left+1)
        } else {
            freq[s[left]]--
            left++
        }
    }
    return maxWin

}

// func characterReplacement(s string, k int) int {
//     freq := map[byte]int{}
//     maxRep := 0
//     left := 0
//     maxWin := 0
//     for i := 0; i < len(s); i++ {
//         freq[s[i]]++
//         maxRep = max(maxRep, freq[s[i]])
//         nonRepeating := (i-left+1) - maxRep
//         for nonRepeating > k {
//             freq[s[left]]--
//             if freq[s[left]] == 0 {delete(freq, s[left])}
//             maxRep = 0
//             for _, v := range freq {maxRep = max(maxRep, v)}
//             left++
//             nonRepeating = (i-left+1) - maxRep            
//         }
//         maxWin = max(maxWin, i-left+1)
//     }
//     return maxWin
// }