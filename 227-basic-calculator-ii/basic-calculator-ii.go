// maintain last contribution to help undo math when 
// operator being processed in * or /
func calculate(s string) int {
    lastContr := 0
    curr := 0
    total := 0
    var operator byte = '+'
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            curr = curr * 10 + int(char-'0')
        }
        if char == '+' || char == '-' || char == '*' || char == '/' || i == len(s)-1 {
            if operator == '+' {
                total += curr
                lastContr = +curr
            } else if operator == '-' {
                total -= curr
                lastContr = -curr
            } else if operator == '*' {
                total = total - lastContr + (lastContr*curr)
                lastContr = lastContr*curr
            } else if operator == '/' {
                total = total - lastContr + (lastContr/curr)
                lastContr = lastContr/curr
            }
            curr = 0
            operator = char
        }
    }
    return total
}
// func calculate(s string) int {
//     st := []int{}
//     digit := 0
//     var operator byte = '+'
//     for i := 0; i < len(s); i++ {
//         char := s[i]
//         if s[i] >= '0' && s[i] <= '9' {
//             charInt := int(char-'0')
//             digit = digit * 10 + charInt
//         }
//         if char == '+' || char == '-' || char == '*' || char == '/' || i == len(s)-1 {
//             if operator == '+' {
//                 st = append(st, digit)
//             } else if operator == '-' {
//                 st = append(st, -digit)
//             } else if operator == '*' {
//                 top := st[len(st)-1]
//                 res := top*digit
//                 st[len(st)-1] = res
//             } else if operator == '/' {
//                 top := st[len(st)-1]
//                 res := top/digit
//                 st[len(st)-1] = res
//             }
//             digit = 0
//             operator = char
//         }
//     }
//     res := 0
//     for len(st) != 0 {
//         top := st[len(st)-1]
//         st = st[:len(st)-1]
//         res += top
//     }
//     return res
// }