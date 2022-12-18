// func decodeString(s string) string {
//     numSt := []int{}
//     strSt := []*strings.Builder{}
//     curr := new(strings.Builder)
//     num := 0
    
//     for i := 0; i < len(s); i++ {
//         if s[i] >= '0' && s[i] <= '9' {
//             num = num * 10 + int(s[i]-'0')
//         } else if s[i] == '[' {
//             strSt = append(strSt, curr)
//             numSt = append(numSt, num)
//             num = 0
//             curr = new(strings.Builder)
//         } else if s[i] == ']' {
//             topNum := numSt[len(numSt)-1]; numSt = numSt[:len(numSt)-1]
//             flattend := new(strings.Builder)
//             for i := 0; i < topNum; i++ {
//                 flattend.WriteString(curr.String())
//             }
//             curr = strSt[len(strSt)-1]; strSt = strSt[:len(strSt)-1]
//             curr.WriteString(flattend.String())
//         } else {
//             curr.WriteByte(s[i])
//         }
//     }
//     return curr.String()
// }


func decodeString(s string) string {
    i := 0
    var dfs func(str string) string
    dfs = func(str string) string {
        // base
        
        // logic
        num := 0
        curr := new(strings.Builder)
        for i < len(str) {
            if str[i] >= '0' && str[i] <= '9' {
                num = num * 10 + int(str[i]-'0')
                i++
            } else if str[i] == '[' {
                i++
                decoded := dfs(str)
                for k := 0; k < num; k++ {
                    curr.WriteString(decoded)
                }
                num = 0
            } else if str[i] == ']' {
                i++
                return curr.String()
            } else {
                curr.WriteByte(str[i])
                i++
            }
        }
        return curr.String()
    }
    return dfs(s)
}