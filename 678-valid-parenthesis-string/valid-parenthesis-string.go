// bhai... matlab... who tf comes up with this SHIT ???? 
// FUCK YOU 
func checkValidString(s string) bool {
    open := 0
    close := 0
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == '(' || s[i] == '*' {
            open++
        } else if s[i] == ')' {
            open--
            if open < 0 {return false}
        }

        if s[n-1-i] == ')' || s[n-1-i] == '*' {
            close++
        } else if s[n-1-i] == '(' {
            close--
            if close < 0 {return false}
        }
    }
    return true
}
// func checkValidString(s string) bool {
//     asterik := []int{}
//     open := []int{}
//     for i := 0; i < len(s); i++ {
//         if s[i] == '*' {
//             asterik = append(asterik, i)
//         } else if s[i] == '(' {
//             open = append(open, i)
//         } else if s[i] == ')' {
//             if len(open) > 0 {
//                 open = open[:len(open)-1]
//             } else if len(asterik) > 0 {
//                 asterik = asterik[:len(asterik)-1]
//             } else {
//                 return false
//             }
//         }
//     }
//     for len(open) > 0 && len(asterik) > 0 && 
//         asterik[len(asterik)-1] > open[len(open)-1] {
//             open = open[:len(open)-1]
//             asterik = asterik[:len(asterik)-1]
//     }
//     return len(open) == 0

// }