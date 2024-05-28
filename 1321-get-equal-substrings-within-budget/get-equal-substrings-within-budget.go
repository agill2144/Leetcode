func equalSubstring(s string, t string, maxCost int) int {
    l1, l2 := 0, 0
    cost := 0
    ans := 0
    for i := 0; i < len(s); i++ {
        sChar := s[i]
        tChar := t[i]
        diff := abs(int(sChar-'a')-int(tChar-'a'))
        cost += diff
        for cost > maxCost {
            sLeft := s[l1]
            tLeft := t[l2]
            leftDiff := abs(int(sLeft-'a')-int(tLeft-'a'))
            cost -= leftDiff
            l1++
            l2++
        }
        ans = max(ans, i-l1+1)
    }
    return ans
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}