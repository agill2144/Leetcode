func customSortString(order string, s string) string {
    idx := map[string]int{}
    for i := 0; i < len(order); i++ {idx[string(order[i])] = i}
    strList := strings.Split(s, "")
    sort.SliceStable(strList, func(i, j int)bool{
        iIdx, ok := idx[strList[i]]
        jIdx, ok2 := idx[strList[j]]
        if ok && ok2 {
            return iIdx < jIdx
        }
        if ok {return true}
        if !ok && ok2 {return false}
        return false
    })
    return strings.Join(strList, "")
}
// func customSortString(order string, s string) string {
// // order and s consist of lowercase English letters.

//     // o(1) space
//     freq := map[byte]int{}
//     // o(s) time
//     for i := 0; i < len(s); i++ {freq[s[i]]++}

//     // space = o(s)
//     out := new(strings.Builder)

//     // time = o(1 * s) = o(s)
//     for i := 0; i < len(order); i++ {
//         for freq[order[i]] > 0 {
//             out.WriteByte(order[i])
//             freq[order[i]]--
//         }
//     }
//     // time = o(s)
//     for i := 0; i < len(s); i++ {
//         for freq[s[i]] > 0 {
//             out.WriteByte(s[i])
//             freq[s[i]]--
//         }        
//     }
//     return out.String()
// }