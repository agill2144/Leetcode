func minWindow(s string, t string) string {
    type pair struct {
        char byte
        idx int
    }
    filtered := []pair{}    
    freq := map[byte]int{}
    for i := 0; i < len(t); i++ {freq[t[i]]++}
    for i := 0; i < len(s); i++ {
        if _, ok := freq[s[i]]; ok {
            filtered = append(filtered, pair{s[i], i})
        }
    }
    fullMatch := 0
    left := 0
    start, end := -1,-1
    for i := 0; i < len(filtered); i++ {
        char := filtered[i].char
        idx := filtered[i].idx
        freq[char]--
        if freq[char] == 0 {fullMatch++}
        for fullMatch == len(freq) {
            leftChar := filtered[left].char
            leftIdx := filtered[left].idx
            if start == -1 || idx-leftIdx+1 < end-start+1 {
                start, end = leftIdx, idx
            }
            freq[leftChar]++
            if freq[leftChar] == 1 {fullMatch--}
            left++
        }
    }
    if start == -1 {return ""}
    return s[start:end+1]
}

// func minWindow(s string, t string) string {
//     freq := map[byte]int{}
//     for i := 0; i < len(t); i++ {freq[t[i]]++}
//     left := 0
//     res := ""
//     fullMatch := 0
//     for i := 0; i < len(s); i++ {
//         _, ok := freq[s[i]]
//         if ok {
//             freq[s[i]]--
//             if freq[s[i]] == 0 {fullMatch++}
//         }
//         for fullMatch == len(freq) {
//             subStr := s[left:i+1]
//             if res == "" || len(subStr) < len(res) {
//                 res = subStr
//             }
//             _, ok = freq[s[left]]
//             if ok {freq[s[left]]++}
//             if freq[s[left]] == 1 {fullMatch--}
//             left++
//         }
//     }
//     return res
// }