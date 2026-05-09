// n = num of words in strings
// k = max len of a word in strs
// tc = o(n * (k + klogk + k))
// sc = o(n)
func groupAnagrams(strs []string) [][]string {
    sortedMap := map[string][]string{}
    for i := 0; i < len(strs); i++ {
        chars := strings.Split(strs[i],"")
        sort.Strings(chars)
        sorted := strings.Join(chars, "")
        sortedMap[sorted] = append(sortedMap[sorted], strs[i])
    }
    out := [][]string{}
    for _, v := range sortedMap {
        out = append(out, v)
    }
    return out
}