// custom sort func
// we know the idx of each char ( idx map )
// then split the s string into a list of strings
// and apply sort.Slice ( where ith char is placed first if ith idx shows up before jth char idx )
// then re-stich the string list 
// time = o(m) + o(n) + o(nlogn) + o(n) = o(m) + o(nlogn) + o(n)
// space = o(n)
// func customSortString(order string, s string) string {
//     // space = o(1) because there are only 26 chars
//     orderIdx := map[string]int{}
//     // m = len(order)
//     // n = len(s)

//     // tc = o(m)
//     for i := 0; i < len(order); i++ {
//         orderIdx[string(order[i])] = i
//     }
//     // tc = o(n)
//     // sc = o(n)
//     sList := strings.Split(s, "")
//     // tc = o(nlogn)
//     // sc = o(n)
//     sort.Slice(sList, func(i, j int)bool{
//         iChar := sList[i]
//         jChar := sList[j]
//         return orderIdx[iChar] < orderIdx[jChar]
//     })
//     // tc = o(n)
//     return strings.Join(sList,"")
// }


// we know the order in order string
// put s chars in a freq map 
// take a ptr on order string
// and write the output string if order[i] exists in s ( freq num of times )
func customSortString(order string, s string) string {
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {freq[s[i]]++}
    out := new(strings.Builder)
    for i := 0; i < len(order); i++ {
        for freq[order[i]] > 0 {
            out.WriteByte(order[i])
            freq[order[i]]--
        }
    }
    for i := 0; i < len(s); i++ {
        for freq[s[i]] > 0 {out.WriteByte(s[i]); freq[s[i]]--}
    }
    
    return out.String()
}
