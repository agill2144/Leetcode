func decodeString(s string) string {
    i := 0
    var dfs func() string
    dfs = func() string {
        // base

        // logic
        currN := 0
        currS := new(strings.Builder)
        for i < len(s) {
            char := s[i]
            if char >= '0' && char <= '9' {
                currN = currN * 10 + int(char-'0')
                i++
            } else if char >= 'a' && char <= 'z' {
                currS.WriteByte(char)
                i++
            } else if char == '[' {
                i++
                res := dfs()
                tmp := new(strings.Builder)
                for k := 0; k < currN; k++ {
                    tmp.WriteString(res)
                }
                currN = 0
                currS.WriteString(tmp.String())
            } else if char == ']' {
                i++
                return currS.String()
            }
        }
        return currS.String()   
    }

    return dfs()


// n = len(string)
// k = numbers within string
// time = n * maxK
// space = o(k) + o(n)
// func decodeString(s string) string {
//     numSt := []int{}
//     strSt := []*strings.Builder{}
//     currN := 0
//     currS := new(strings.Builder)

//     for i := 0; i < len(s); i++ {
//         char := s[i]
//         if char >= '0' && char <= '9' {
//             currN =  currN * 10 + int(char-'0')
//         } else if char >= 'a' && char <= 'z' {
//             currS.WriteByte(char)
//         } else if char == '[' {
//             numSt = append(numSt, currN)
//             strSt = append(strSt, currS)
//             currN = 0
//             currS = new(strings.Builder)
//         } else if char == ']' {
//             times := numSt[len(numSt)-1]
//             numSt = numSt[:len(numSt)-1]
//             tmp := new(strings.Builder)
//             for k := 0; k < times; k++ {
//                 tmp.WriteString(currS.String())
//             }
//             parentStr := strSt[len(strSt)-1]
//             strSt = strSt[:len(strSt)-1]
//             parentStr.WriteString(tmp.String())
//             currS = parentStr
//         }
//     }
//     return currS.String()
}