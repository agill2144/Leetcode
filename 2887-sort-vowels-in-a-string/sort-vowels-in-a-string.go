func sortVowels(s string) string {
    set := map[byte]bool{
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }
    vFreq := make([]int, 60)
    for i := 0; i < len(s); i++ {
        if set[s[i]] {vFreq[s[i]-'A']++}
    }
    sortedOrder := []byte{'A','E','I','O','U','a','e','i','o','u'}
    sPtr := 0
    out := new(strings.Builder)    
    for i := 0; i < len(s); i++ {
        if !set[s[i]] {
            out.WriteByte(s[i])
        } else {
            // we are at a vowel position in OG string
            // now we need the correct sorted vowel char
            // hence, based on the sorted order, 
            // find next available vowel char which has a freq > 0
            for sPtr < len(sortedOrder) && vFreq[int(sortedOrder[sPtr]-'A')] == 0 {
                sPtr++
            }

            // found a vowel char , write it to new string
            out.WriteByte(sortedOrder[sPtr])
            // decrement its count by 1
            vFreq[int(sortedOrder[sPtr]-'A')]--
        }
    }
    
    return out.String()
}
/*
    1. extract vowels
    2. sort vowels
    3. write back a new string
    tc : o(s) + o(slogs) + o(s)
    sc : o(s) for vs array
*/
// func sortVowels(s string) string {
//     set := map[byte]bool{
//         'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
//         'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
//     }
//     vs := []byte{}
//     for i := 0; i < len(s); i++ {
//         if set[s[i]] {vs = append(vs,s[i])}
//     }
//     sort.Slice(vs, func(i, j int)bool{
//         return vs[i] < vs[j]
//     })
//     vPtr := 0
//     out := new(strings.Builder)
//     for i := 0; i < len(s); i++ {
//         if set[s[i]] {
//             out.WriteByte(vs[vPtr])
//             vPtr++
//         } else {
//             out.WriteByte(s[i])
//         }
//     }
//     return out.String()
// }