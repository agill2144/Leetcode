// n = num of words in strings
// k = max len of a word in strs
// tc = o(n * k)
// sc = o(n)
func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    for i := 0; i < len(strs); i++ {
        hashKey := hash(strs[i])
        m[hashKey] = append(m[hashKey], strs[i])
    }
    out := [][]string{}
    for _, v := range m {
        out = append(out, v)
    }
    return out
}
// k = max len of word
// tc = o(k) + o(26) = o(k)
// sc = o(26) + o(k) = o(k)
func hash(word string) string {
    freq := make([]int, 26)
    for i := 0; i < len(word); i++ {
        idx := int(word[i]-'a')
        freq[idx]++
    }
    out := new(strings.Builder)
    for i := 0; i < len(freq); i++ {
        out.WriteString(fmt.Sprintf("%v",freq[i]))
        if i != len(freq)-1 {
            out.WriteString("-")
        }
    }
    return out.String()
}
// n = num of words in strings
// k = max len of a word in strs
// tc = o(n * (k + klogk + k))
// sc = o(n)
// func groupAnagrams(strs []string) [][]string {
//     sortedMap := map[string][]string{}
//     for i := 0; i < len(strs); i++ {
//         chars := strings.Split(strs[i],"")
//         sort.Strings(chars)
//         sorted := strings.Join(chars, "")
//         sortedMap[sorted] = append(sortedMap[sorted], strs[i])
//     }
//     out := [][]string{}
//     for _, v := range sortedMap {
//         out = append(out, v)
//     }
//     return out
// }