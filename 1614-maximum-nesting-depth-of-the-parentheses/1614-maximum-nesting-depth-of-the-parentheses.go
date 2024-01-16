/*
    approach: count
    - use count method used to validate paranthesis
    - keep a max ans saved after incrementing count
*/
func maxDepth(s string) int {
    max := 0
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
        } else if s[i] == ')' {
            count--
        }
        if count > max {max = count}
    }
    return max
}