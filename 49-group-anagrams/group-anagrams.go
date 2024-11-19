
func groupAnagrams(strs []string) [][]string {
    grps := map[string][]string{}
    for i := 0; i < len(strs); i++ {
        key := naiveHash(strs[i])
        grps[key] = append(grps[key], strs[i])
    }
    out := [][]string{}
    for _, v := range grps {out = append(out, v)}
    return out
}

func naiveHash(key string) string {
    freq := make([]int, 26)
    for i := 0; i < len(key); i++ {
        idx := int(key[i]-'a')
        freq[idx]++
    }
    res := new(strings.Builder)
    for i := 0; i < len(freq); i++ {
        res.WriteString(fmt.Sprintf("%v",freq[i]))
        if i != len(freq)-1 {res.WriteByte('-')}
    }
    return res.String()
}

var primes []float64 = []float64{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107}
func hash(key string) float64 {
    prod := 1.0
    for i := 0; i < len(key); i++ {
        prod *= (primes[int(key[i]-'a')])
    }
    return prod
}
// func groupAnagrams(strs []string) [][]string {
//     groups := map[string][]string{}
//     // n = len strs
//     // k = avg len of each word
//     // (n * klogk) + n
//     // space = o(k + n)
//     for i := 0; i < len(strs); i++ {
//         word := strings.Split(strs[i], "")
//         sort.Strings(word)
//         sortStr := strings.Join(word,"")
//         groups[sortStr] = append(groups[sortStr], strs[i])
//     }
//     out := [][]string{}
//     for _, v := range groups {
//         out = append(out, v)
//     }
//     return out
// }