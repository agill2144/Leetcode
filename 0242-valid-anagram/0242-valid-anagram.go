// char freq count
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {return false}
    sCount := make([]int, 26)
    totalCount := 0
    for i := 0; i < len(s); i++ {
        sCount[int(s[i]-'a')]++
        totalCount++
    }
    for i := 0; i < len(t); i++ {
        tCharIdx := int(t[i]-'a')
        if sCount[tCharIdx] == 0 {return false}
        sCount[tCharIdx]--
        totalCount--
    }
    return totalCount == 0
}

// sort and do equality check
// func isAnagram(s string, t string) bool {
//     // o(s)
//     sList := strings.Split(s,"")
//     // o(s)
//     tList := strings.Split(t,"")
//     // o(sLogs)
//     sort.Strings(sList)
//     // o(tLogt)
//     sort.Strings(tList)
//     // o(s) + o(t)
//     return strings.Join(sList,"") == strings.Join(tList, "")
// }