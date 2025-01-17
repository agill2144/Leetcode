func customSortString(order string, s string) string {
    freq := make([]int, 26)
    for i := 0; i < len(s); i++ {
        idx := int(s[i]-'a')
        freq[idx]++
    }
    res := new(strings.Builder)
    for i := 0; i < len(order); i++ {
        for freq[int(order[i]-'a')] > 0 {
            res.WriteByte(order[i])
            freq[int(order[i]-'a')]--
        }
    }
    for i := 0; i < len(s); i++ {
        if freq[int(s[i]-'a')] > 0 {
            res.WriteByte(s[i])
            freq[int(s[i]-'a')]--
        }
    }
    return res.String()
}
// m = len(order)
// n = len(s)
// tc = o(m) + o(nlogn) 
// sc = o(n)
// func customSortString(order string, s string) string {
//     idx := map[string]int{}
//     for i := 0; i < len(order); i++ {
//         idx[string(order[i])] = i
//     }
//     chars := strings.Split(s, "")
//     sort.Slice(chars, func(i, j int)bool {
//         iidx, iok := idx[chars[i]]
//         jidx, jok := idx[chars[j]]
//         if iok && jok {
//             return iidx < jidx
//         } else if iok {
//             return true
//         } else if jok {
//             return false
//         }
//         return true
//     })
//     return strings.Join(chars, "")
// }