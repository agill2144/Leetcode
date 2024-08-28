func equalSubstring(s string, t string, maxCost int) int {
    rCost := 0
    left := 0
    maxWin := 0
    n := len(s)
    for i := 0; i < n; i++ {
        rCost += abs(int(s[i]-'a') - int(t[i]-'a'))
        if rCost > maxCost {
            rCost -= abs(int(s[left]-'a') - int(t[left]-'a'))
            left++
        } else {
            maxWin = max(maxWin, i-left+1)
        }
    } 
    return maxWin
}
func abs(x int) int {
    if x < 0 {return x*-1 }
    return x
}