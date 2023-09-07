func decodeString(s string) string {
    i := 0
    var dfs func() string
    dfs = func() string {
        // base
        // logic
        curr := new(strings.Builder)
        num := 0
        for i < len(s) {
            char := s[i]
            if char == '[' {
                i++
                child := dfs()
                tmp := new(strings.Builder)
                for num != 0 {
                    tmp.WriteString(child)
                    num--
                }
                curr.WriteString(tmp.String())
                num = 0
            } else if char == ']' {
                i++
                return curr.String()
            } else if char >= '0' && char <= '9' {
                num = num * 10 + int(char-'0')
                i++
            } else {
                curr.WriteByte(char)
                i++
            }
        }
        return curr.String()
    }
    return dfs()   
}


// func decodeString(s string) string {
//     numSt := []int{}
//     parentSt := []*strings.Builder{}
    
//     curr := new(strings.Builder)
//     num := 0
    
//     for i := 0; i < len(s); i++ {
//         char := s[i]
//         if char == '[' {
//             numSt = append(numSt, num)
//             parentSt = append(parentSt, curr)
//             curr = new(strings.Builder)
//             num = 0
//         } else if char == ']' {
//             k := numSt[len(numSt)-1]
//             numSt = numSt[:len(numSt)-1]
//             tmp := new(strings.Builder)
//             for k != 0 {
//                 tmp.WriteString(curr.String())
//                 k--
//             }
//             parent := parentSt[len(parentSt)-1]
//             parentSt = parentSt[:len(parentSt)-1]
//             parent.WriteString(tmp.String())
//             curr = parent
            
//         } else if char >= '0' && char <= '9' {
//             num = num * 10 + int(char-'0')
//         } else {
//             curr.WriteByte(char)
//         }
//     }
//     return curr.String()
// }