func minWindow(s string, t string) string {
    if len(t) > len(s) {return ""}
    tMap := map[byte]int{}
    for i := 0; i < len(t); i++ {tMap[t[i]]++}
    fm := 0
    minWin := len(s)+1
    minStr := ""
    left := 0
    for i := 0; i < len(s); i++ {
        _, ok := tMap[s[i]]
        if ok {
            tMap[s[i]]--
            if tMap[s[i]] == 0 {fm++}
        }
        for fm == len(tMap) {
            if i-left+1 < minWin {
                minWin = i-left+1
                minStr = s[left:i+1]
            }
            minWin = min(minWin, i-left+1)
            leftChar := s[left]
            if _, ok := tMap[leftChar]; ok {
                tMap[leftChar]++
                if tMap[leftChar] == 1 {fm--}
            }
            left++
        }

    }
    return minStr
}