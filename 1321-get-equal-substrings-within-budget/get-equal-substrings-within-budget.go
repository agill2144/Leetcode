func equalSubstring(s string, t string, maxCost int) int {
    maxWin := 0
    left := 0
    rCost := 0
    for i := 0; i < len(s); i++ {
        diff := abs(int(s[i]-'a') - int(t[i]-'a'))
        rCost += diff
        for rCost > maxCost && left <= i {
            leftCost := abs(int(s[left]-'a') - int(t[left]-'a'))
            rCost -= leftCost
            left++
        }
        maxWin = max(maxWin, i-left+1)
    } 
    return maxWin
}
func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}