func minAddToMakeValid(s string) int {
    count := 0 // open paran count ( like stack )
    closeCount := 0
    for i := 0 ; i < len(s); i++ {
        if s[i] == '(' {
            count++
        } else {
            if count > 0 {count--} else {closeCount++}
        }
    }
    return count + closeCount
}

func abs(x int) int {
    if x < 0 {return x *-1}
    return x
}