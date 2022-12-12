// func backspaceCompare(s string, t string) bool {
//     newS := ""
//     skipCount := 0
    
//     for i := len(s)-1; i>=0 ; i-- {
//         char := string(s[i])
//         if char == "#" {
//             skipCount++
//         } else if skipCount != 0 {
//             skipCount--
//         } else {
//             newS += char
//         }
//     }
    
//     newT := ""
//     skipCountT := 0
    
//     for i := len(t)-1; i>=0 ; i-- {
//         char := string(t[i])
//         if char == "#" {
//             skipCountT++
//         } else if skipCountT != 0 {
//             skipCountT--
//         } else {
//             newT += char
//         }
//     }
//     return newS == newT
// }


func backspaceCompare(s string, t string) bool {
    newS := ""
    newT := ""
    skipSCount := 0
    skipTCount := 0
    sPtr := len(s)-1
    tPtr := len(t)-1
    
    for sPtr >= 0 || tPtr >= 0 {
        
        // check s        
        if sPtr >= 0 {
            sChar := string(s[sPtr])
            if sChar == "#" {
                skipSCount++
            } else if skipSCount != 0 {
                skipSCount--
            } else {
                newS += sChar
            }
            sPtr--
        }
        
        // check t
        if tPtr >= 0 {
            tChar := string(t[tPtr])
            if tChar == "#" {
                skipTCount++
            } else if skipTCount != 0 {
                skipTCount--
            } else {
                newT += tChar
            }
            tPtr--
        }
    }
    
    
    return newS == newT
}