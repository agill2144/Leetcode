func maxDepth(s string) int {
    if len(s) <= 1 {return 0}
    count := 0
    ans := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
            ans = max(ans, count)
        }  else if s[i] == ')' {
            count--
        }
    }
    return ans
}