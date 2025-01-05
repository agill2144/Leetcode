func minWindow(s string, t string) string {
    tFreq := map[byte]int{}
    for i := 0; i < len(t); i++ {tFreq[t[i]]++}
    type char struct {
        c byte
        idx int
    }
    filteredS := []*char{}
    for i := 0; i < len(s); i++ {
        if _, ok := tFreq[s[i]]; ok {
            filteredS = append(filteredS, &char{s[i], i})
        }
    }
    left := 0
    start, end := -1,-1
    fullMatch := 0
    for i := 0; i < len(filteredS); i++ {
        c,idx := filteredS[i].c, filteredS[i].idx
        tFreq[c]--
        if tFreq[c] == 0 {
            fullMatch++
        }
        for fullMatch == len(tFreq) {
            leftChar := filteredS[left].c
            leftIdx := filteredS[left].idx
            if start == -1 || idx-leftIdx+1 < end-start+1 {
                start = leftIdx
                end = idx
            }
            if _, ok := tFreq[leftChar]; ok {
                tFreq[leftChar]++
                if tFreq[leftChar] == 1 {
                    fullMatch--
                }
            }
            left++
        }
    }
    if start == -1 {return ""}
    return s[start:end+1]
}