func canBeValid(s string, locked string) bool {
    if len(s) % 2 != 0 {return false}
    flexible := []int{}
    open := []int{}
    for i := 0; i < len(s); i++ {
        if locked[i] == '0' {
            flexible = append(flexible, i)
        } else if s[i] == '(' {
            open = append(open, i)
        } else if s[i] == ')' {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(flexible) > 0 {
                flexible = flexible[:len(flexible)-1]
            } else {
                return false
            }
        }
    }
    for len(open) > 0 && len(flexible) > 0 {
        if flexible[len(flexible)-1] > open[len(open)-1] {
            flexible = flexible[:len(flexible)-1]
            open = open[:len(open)-1]
            continue
        }
        break
    }
    return len(open) == 0 && len(flexible) % 2 == 0
}