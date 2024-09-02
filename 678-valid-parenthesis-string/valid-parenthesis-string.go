func checkValidString(s string) bool {
    open := []int{} // idx of open parans
    asterik := []int{} // idx of asteriks
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char == '(' {
            open = append(open, i)
        } else if char == '*' {
            asterik = append(asterik, i)
        } else if char == ')' {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(asterik) > 0 {
                // use last-seen / prev asterik to create a open paran for 
                // this closing paran
                // essentially, replace last seen asterik as a open-paran
                asterik = asterik[:len(asterik)-1]
            } else {
                // too many closing parans without an asterik or a open paran
                // therefore never possible to make this closing paran balanced
                // therefore return false
                return false
            }
        }
    }

    // its possible that we still have open params,
    // process those as much as possible
    // in this case, we are trying to close open parans
    // meaning, asteriks idx MUST be after open paran idx
    for len(open) > 0 && len(asterik) > 0 {
        if asterik[len(asterik)-1] > open[len(open)-1] {
            asterik = asterik[:len(asterik)-1]
            open = open[:len(open)-1]
            continue
        }
        return false
    }
    return len(open) == 0
}