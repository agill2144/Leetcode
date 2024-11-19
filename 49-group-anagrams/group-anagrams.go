func groupAnagrams(strs []string) [][]string {
    groups := map[string][]string{}
    for i := 0; i < len(strs); i++ {
        word := strings.Split(strs[i], "")
        sort.Strings(word)
        sortStr := strings.Join(word,"")
        groups[sortStr] = append(groups[sortStr], strs[i])
    }
    out := [][]string{}
    for _, v := range groups {
        out = append(out, v)
    }
    return out
}