func checkValidString(s string) bool {
    openParanSt := []int{} // idx of open parans
    asterikSt := []int{} // idx of asteriks
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '(' {
            openParanSt = append(openParanSt, i)
        } else if char == '*' {
            asterikSt = append(asterikSt, i)
        } else if char == ')' {
            if len(openParanSt) > 0 {
                openParanSt = openParanSt[:len(openParanSt)-1]
            } else if len(asterikSt) > 0 {
                // use last-seen / prev asterik to create a open paran for 
                // this closing paran
                // essentially, replace last seen asterik as a open-paran
                asterikSt = asterikSt[:len(asterikSt)-1]
            } else {
                return false
            }
        }
    }

    // its possible that we still have open params,
    // process those as much as possible
    // in this case, we are trying to close open parans
    // meaning, asteriks idx MUST be after open paran idx
    for len(openParanSt) > 0 && len(asterikSt) > 0 {
        if asterikSt[len(asterikSt)-1] > openParanSt[len(openParanSt)-1] {
            asterikSt = asterikSt[:len(asterikSt)-1]
            openParanSt = openParanSt[:len(openParanSt)-1]
            continue
        }
        return false
    }
    return len(openParanSt) == 0
}