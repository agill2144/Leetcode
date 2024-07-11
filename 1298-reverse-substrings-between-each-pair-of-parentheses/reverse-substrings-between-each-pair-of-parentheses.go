func reverseParentheses(s string) string {
    st := []int{}
    pairs := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            pairs[i] = top
            pairs[top] = i
        }
    }

    res := new(strings.Builder)
    i := 0
    dir := 1
    for i < len(s) {
        if s[i] == '(' || s[i] == ')' {
            i = pairs[i]
            dir *= -1
        } else {
            res.WriteByte(s[i])
        }
        i += dir
    }
    return res.String()
}

// func reverseParentheses(s string) string {
//     st := []*strings.Builder{}
//     curr := new(strings.Builder)
//     for i := 0; i < len(s); i++ {
//         if s[i] == '(' {
//             st = append(st, curr)
//             curr = new(strings.Builder)
//         } else if s[i] == ')' {
//             currRev := reverse(curr)
//             if len(st) != 0 {
//                 parent := st[len(st)-1]
//                 st = st[:len(st)-1]
//                 parent.WriteString(currRev.String())
//                 curr = parent
//             } else {
//                 curr = currRev
//             }
//         } else {
//             curr.WriteByte(s[i])
//         }
//     }
//     return curr.String()
// }


// func reverse(s *strings.Builder) *strings.Builder {
//     newS := new(strings.Builder)
//     sStr := s.String()
//     for i := len(sStr)-1; i >= 0 ; i-- {
//         newS.WriteByte(sStr[i])
//     }
//     return newS
// }
