// n = len(strs)
// k = avg len of each str
// tc = o(n * (k+klogk+k)) + o(n)
// sc = o(n)
func groupAnagrams(strs []string) [][]string {
    // sc = o(n)
    grps := map[string][]string{}
    for i := 0; i < len(strs); i++ { // n
        chars := strings.Split(strs[i],"") // k 
        sort.Strings(chars) // klogk
        sorted := strings.Join(chars, "") // k
        grps[sorted] = append(grps[sorted], strs[i]) // k
    } // n * (k + klogk + k)
    out := [][]string{}
    // n
    for _, v := range grps {
        out = append(out, v)
    }
    return out
}