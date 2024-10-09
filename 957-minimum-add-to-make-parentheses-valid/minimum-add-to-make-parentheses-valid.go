func minAddToMakeValid(s string) int {
    count := 0 // open paran count ( like stack )
    closeCount := 0 // close paran showed up when there were no open ones
    for i := 0 ; i < len(s); i++ {
        if s[i] == '(' {
            count++
        } else {
            if count > 0 {count--} else {closeCount++}
        }
    }
    return count + closeCount
}