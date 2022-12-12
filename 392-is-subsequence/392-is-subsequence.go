func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {return false}
    left := 0
    right := 0
    
    for left != len(s) && right != len(t) {
        leftChar := string(s[left])
        rightChar := string(t[right])
        
        if leftChar == rightChar {
            left++
            right++
        } else {
            right++
        }
    }
    return left >= len(s)
}