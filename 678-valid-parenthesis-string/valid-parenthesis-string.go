func checkValidString(s string) bool {
    asterik := []int{}
    open := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] == '*' {
            asterik = append(asterik, i)
        } else if s[i] == '(' {
            open = append(open, i)
        } else if s[i] == ')' {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(asterik) > 0 {
                asterik = asterik[:len(asterik)-1]
            } else {
                return false
            }
        }
    }
    for len(open) > 0 && len(asterik) > 0 && 
        asterik[len(asterik)-1] > open[len(open)-1] {
            open = open[:len(open)-1]
            asterik = asterik[:len(asterik)-1]
    }
    return len(open) == 0

}