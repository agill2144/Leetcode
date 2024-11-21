func minRemoveToMakeValid(s string) string {
    open := 0
    res := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            open++
            res.WriteByte(s[i])
        } else if s[i] == ')' {
            if open == 0 {continue}
            open--
            res.WriteByte(s[i])
        } else {
            res.WriteByte(s[i])
        }
    }
    resStr := res.String()
    close := 0
    res.Reset()
    for i := len(resStr)-1; i >= 0; i-- {
        if resStr[i] == ')' {
            close++
            res.WriteByte(resStr[i])
        } else if resStr[i] == '(' {
            if close == 0 {continue}
            close--
            res.WriteByte(resStr[i])        
        } else {
            res.WriteByte(resStr[i])
        }
    }
    resStr = res.String()
    res.Reset()
    for i := len(resStr)-1; i >= 0; i-- {
        res.WriteByte(resStr[i])
    }
    return res.String()

}
// func minRemoveToMakeValid(s string) string {
//     invalidIdxs := map[int]bool{}
//     st := []int{}
//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' {
//             st = append(st, i)
//         } else if s[i] == ')' {
//             if len(st) == 0 {
//                 invalidIdxs[i] = true
//                 continue
//             }
//             st = st[:len(st)-1]
//         }
//     }
//     for len(st) != 0 {
//         top := st[len(st)-1]
//         st = st[:len(st)-1]
//         invalidIdxs[top] = true 
//     }
//     res := new(strings.Builder)
//     for i := 0; i < len(s); i++ {
//         if invalidIdxs[i] {continue}
//         res.WriteByte(s[i])
//     }
//     return res.String()
// }