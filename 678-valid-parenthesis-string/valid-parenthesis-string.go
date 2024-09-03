func checkValidString(s string) bool {
    open := []int{}
    asteriks := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            open = append(open, i)
        } else if s[i] == '*' {
            asteriks = append(asteriks, i)
        } else if s[i] == ')' {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(asteriks) > 0 {
                // use a previously seen asterik as a open paran
                asteriks = asteriks[:len(asteriks)-1]
            } else {
                // we are at a closing paran
                // without a openining paran and
                // without a asterik in the past
                // therefore this string can never be valid
                return false
            }
        }
    }

    // its possible that we have open parans left
    // now we will use remaining asteriks as the closing paran.
    // for the asteriks to be used as a closing paran, its idx must be
    // AFTER opening paran idx. because a paran pair is valid if we have a open and its corresponding
    // closer is after.
    for len(open) > 0 && len(asteriks) > 0 {
        if asteriks[len(asteriks)-1] > open[len(open)-1] {
            open = open[:len(open)-1]
            asteriks = asteriks[:len(asteriks)-1]
            continue
        }
        break
    }
    return len(open) == 0
}