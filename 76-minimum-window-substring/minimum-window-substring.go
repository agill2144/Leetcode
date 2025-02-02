type pair struct {
    char byte
    idx int
}
func minWindow(s string, t string) string {
    tFreq := map[byte]int{}    
    for i := 0; i < len(t); i++ {tFreq[t[i]]++}
    fs := []*pair{}
    for i := 0; i < len(s); i++ {
        _, ok := tFreq[s[i]]
        if ok {
            fs = append(fs, &pair{s[i],i})
        }
    }
    start, end := -1,-1
    left := 0
    fullMatch := 0
    for i := 0; i < len(fs); i++ {
        char, idx := fs[i].char,fs[i].idx
        tFreq[char]--
        if tFreq[char] == 0 {fullMatch++}

        for fullMatch == len(tFreq) {
            leftPair := fs[left]
            leftIdx := leftPair.idx
            leftChar := leftPair.char
            if start == -1 || (idx-leftIdx+1 < end-start+1) {
                start, end = leftIdx, idx
            }
            tFreq[leftChar]++
            if tFreq[leftChar] == 1 {fullMatch--}
            left++
        }
    }
    if start == -1 {return ""}
    return s[start:end+1]
}