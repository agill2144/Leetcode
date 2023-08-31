func maxDepth(s string) int {
    st := []byte{}
    count := 0
    maxCount := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '('{
            st = append(st, s[i])
            count++
        } else if s[i] == ')' {
            st = st[:len(st)-1]
            maxCount = max(count, maxCount)
            count--
        }
    }
    return maxCount
}
func max(x, y int) int {
    if x > y {return x }
    return y
}