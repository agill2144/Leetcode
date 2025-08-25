// hash it based on freq of each char in each word
// hash key = [26]int{} which has freq of each char
// this will be in sorted form by default
// so if 2 words are anagrams (aba vs aab ) -> their freq arr will lool identical
// aba = [2,1, 0,0,....]
// aab = [2,1, 0,0,....]
// hence proving anagrams
// tc = o(n * k) + o(n)
// sc = o(n)
func groupAnagrams(strs []string) [][]string {
    grps := map[[26]int][]string{}    
    for i := 0; i < len(strs); i++ {
        key := hash(strs[i])
        grps[key] = append(grps[key], strs[i])
    }
    out := [][]string{}
    for _, v := range grps {
        out = append(out, v)
    }
    return out
}
// tc = o(k)
// sc = o(26)
func hash(str string) [26]int{
    out := [26]int{}
    for i := 0; i < len(str); i++ {
        charIdx := str[i]-'a'
        out[charIdx]++
    }
    return out
}
// if there are multiple words that are anagrams of each other
// then a sorted form of that word will all map to the same word
// using this we can build a map with sorted str as the key 
// and values will be a list of strs that when sorted map to the same exact key
// essentially making them collide because 2 words that are anagrams will map to the same sorted word
// n = len(strs)
// k = avg len of each str
// tc = o(n * (k+klogk+k)) + o(n)
// sc = o(n)
// func groupAnagrams(strs []string) [][]string {
//     // sc = o(n)
//     grps := map[string][]string{}
//     for i := 0; i < len(strs); i++ { // n
//         chars := strings.Split(strs[i],"") // k 
//         sort.Strings(chars) // klogk
//         sorted := strings.Join(chars, "") // k
//         grps[sorted] = append(grps[sorted], strs[i]) // k
//     } // n * (k + klogk + k)
//     out := [][]string{}
//     // n
//     for _, v := range grps {
//         out = append(out, v)
//     }
//     return out
// }