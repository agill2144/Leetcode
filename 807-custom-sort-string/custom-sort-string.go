/*
    we want to sort s based on a custom "order"
    therefore we should collect all chars from order string that also exist in "s"
    because order matters, chars in $order must be placed first
    
    appraoch: using a freq map
    - store chars of s in a freq map
    - loop thru the order string to collect chars that should be placed first
    - if the order[i] does exist in s map, take it and add it to our output
        - freq val times
        - delete this key from s map
    - then loop over s string one last time, to collect remaining chars that did not exist in $order
    - there is a possibility that not all chars of $s exist in $order,
    - therefore we must also preserve the order of chars in $s that do not exist in $order
    
    s = len(s)
    t = len(order)
    
    time = o(s) + o(t) + o(s)
    space = o(s)
*/
// func customSortString(order string, s string) string {
    
//     // make the chars of $s searchable 
//     freqMap := map[byte]int{}
//     for i := 0; i < len(s); i++ {
//         freqMap[s[i]]++
//     }
    
//     out := new(strings.Builder)
//     // collect the chars from order that also exist in $s
//     for i := 0; i < len(order); i++ {
//         char := order[i]
//         count := freqMap[char]
//         for count > 0 {
//             out.WriteByte(char)
//             count--
//         }
//         delete(freqMap, char)
//     }
    
//     // collect the remaining left over chars from $s that did not exist in $order
//     for i := 0; i < len(s); i++ {
//         count := freqMap[s[i]]
//         if count > 0 {
//             out.WriteByte(s[i])
//         }
//     }
//     return out.String()
// }


func customSortString(order string, s string) string {
    orderIdx := map[string]int{}
    for i := 0; i < len(order); i++ {
        orderIdx[string(order[i])] = i
    }
    sList := strings.Split(s, "")
    sort.Slice(sList, func(i, j int)bool{
        iChar := sList[i]
        jChar := sList[j]
        return orderIdx[iChar] < orderIdx[jChar]
    })
    return strings.Join(sList,"")
}