func minWindow(s string, t string) string {
    freq := map[byte]int{}
    for i := 0; i < len(t); i++ {freq[t[i]]++}
    left := 0
    start, end := -1,-1
    fullMatch := 0
    for i := 0; i < len(s); i++ {
        _, ok := freq[s[i]]
        if ok {
            freq[s[i]]--
            if freq[s[i]] == 0 {fullMatch++}
        }
        for fullMatch == len(freq) {
            if start == -1 || i-left+1 < end-start+1 {
                start = left; end = i
            }
            _, ok = freq[s[left]]
            if ok {freq[s[left]]++}
            if freq[s[left]] == 1 {fullMatch--}
            left++
        }
    }
    if start == -1 {return ""}
    return s[start:end+1]
}

// brute force: o(t) + o(s^2 * t)
// sc - o(1)
// func minWindow(s string, t string) string {
//     tFreq := map[byte]int{}
//     for i := 0; i < len(t); i++ {tFreq[t[i]]++}
//     start, end := -1,-1
//     for i := 0; i < len(s); i++ {
//         if tFreq[s[i]] > 0 {
//             freq := map[byte]int{}
//             for k, v := range tFreq {freq[k] = v}
//             fullMatch := 0
//             ss, ee := -1,-1
//             for j := i; j < len(s); j++ {
//                 _, ok := freq[s[j]]
//                 if ok {
//                     freq[s[j]]--
//                     if freq[s[j]] == 0 {
//                         fullMatch++
//                     }
//                 }
//                 if fullMatch == len(freq) {
//                     ss, ee = i, j
//                     break
//                 }
//             }
//             if ss != -1 && (start == -1 || ee-ss+1 < end-start+1) {start, end = ss, ee}
//         }
//     }
//     if start == -1 {return ""}
//     return s[start:end+1]
// }