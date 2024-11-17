func minRemoveToMakeValid(s string) string {
    // a left ( paran is balanced with a right )
    // so if there is a extra right ) when there was no open left (
    // this means, this extra ) is invalid, and can be removed

    // 1. 1st pass to remove all invalid closing parans
    open := 0
    tmp := new(strings.Builder)
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            open++
            tmp.WriteByte(s[i])
        } else if s[i] == ')' {
            // is this an invalid closing?
            // it would be if we dont have any openings
            // but if we had openings, this is a valid char
            if open > 0 {
                open--
                tmp.WriteByte(s[i])
            }
        }  else {
            // regular char from a to z
            tmp.WriteByte(s[i])
        }
    }

    // a right closing paran is balanced by its left opening
    // so if there was a extra ( open paran, when there was nothing closing on right
    // that means this extra ( open paran is invalid and can be removed
    // we need to traverse from right to left because now we are trying to balance from the closing to opening
    // and closing paran would show up after opening if traversing from left to right
    // therefore traversing from right to left
    tmpStr := tmp.String()
    tmp.Reset()
    close := 0
    for i := len(tmpStr)-1; i >= 0; i-- {
        if tmpStr[i] == ')' {
            close++
            tmp.WriteByte(tmpStr[i])
        } else if tmpStr[i] == '(' {
            if close > 0 {
                close--
                tmp.WriteByte(tmpStr[i])
            }
        } else {
            tmp.WriteByte(tmpStr[i])
        }
    }
    tmpStr = tmp.String()
    tmp.Reset()
    for i := len(tmpStr)-1; i >= 0; i-- {
        tmp.WriteByte(tmpStr[i])
    }
    return tmp.String()
}