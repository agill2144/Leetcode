func checkValidString(s string) bool {
    asteriks := []int{}
    open := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            asteriks = append(asteriks, i)
        } else if s[i] == '(' {
            open = append(open, i)
        } else {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(asteriks) > 0 {
                // use asteriks to place a missing open param
                asteriks = asteriks[:len(asteriks)-1]
            } else {
                return false
            }
        }
    }

    // if we still have open parans that did not close
    // we need asteriks to appear after open paran idx
    for len(open) != 0 && len(asteriks) != 0 {
        if asteriks[len(asteriks)-1] > open[len(open)-1] {
            // asterik appears after open paran
            // therefore this asterik can be used as a closing paran
            asteriks = asteriks[:len(asteriks)-1]
            open = open[:len(open)-1]
        } else {
            return false
        }
    }
    return len(open) == 0
}