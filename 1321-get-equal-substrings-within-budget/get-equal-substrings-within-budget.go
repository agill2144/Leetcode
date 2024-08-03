func equalSubstring(s string, t string, maxCost int) int {
    maxWin := 0
    left := 0
    cost := 0
    for i := 0; i < len(s); i++ {
        diff := abs(int(s[i]-'a')-int(t[i]-'a'))
        cost += diff
        for cost > maxCost {
            leftCost := abs(int(s[left]-'a')-int(t[left]-'a'))
            cost -= leftCost
            left++
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}