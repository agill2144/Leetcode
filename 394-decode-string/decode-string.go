func decodeString(s string) string {
    ptr := 0
    var dfs func(str string) string
    dfs = func(str string) string{
        curr := new(strings.Builder)
        n := 0
        for ptr < len(str) {
            char := s[ptr]
            ptr++
            if char >= '0' && char <= '9' {
                n = n * 10 + int(char-'0')
            } else if char >= 'a' && char <= 'z' {
                curr.WriteByte(char)
            } else {
                if char == '[' {
                    res := dfs(s)
                    decoded := new(strings.Builder)
                    for n > 0 {
                        decoded.WriteString(res)
                        n--
                    }
                    curr.WriteString(decoded.String())
                    n = 0
                } else if char == ']' {

                    return curr.String()
                }
            }
        }
        return curr.String()

    }
    return dfs(s)
}

// func decodeString(s string) string {
//     strSt := []*strings.Builder{}
//     numSt := []int{}
//     curr := new(strings.Builder)
//     n := 0
//     for i := 0; i < len(s); i++ {
//         char := s[i]
//         if char >= '0' && char <= '9' {
//             n = n * 10 + int(char-'0')
//         } else if char >= 'a' && char <= 'z' {
//             curr.WriteByte(char)
//         } else {
//             if char == '[' {
//                 strSt = append(strSt, curr)
//                 numSt = append(numSt, n)
//                 curr = new(strings.Builder)
//                 n = 0
//             } else if char == ']' {
//                 // repeat curr string k times
//                 // k = top of num st
//                 k := numSt[len(numSt)-1]
//                 numSt = numSt[:len(numSt)-1]
//                 decoded := new(strings.Builder)
//                 for k > 0 {
//                     decoded.WriteString(curr.String())
//                     k--
//                 }
//                 // concatnate $parent + $decoded
//                 // parent = top of strSt
//                 parent := strSt[len(strSt)-1]
//                 strSt = strSt[:len(strSt)-1]
//                 parent.WriteString(decoded.String())
//                 curr = parent
//             }
//         }
//     }
//     return curr.String()
// }