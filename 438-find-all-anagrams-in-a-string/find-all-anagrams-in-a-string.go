func findAnagrams(s string, p string) []int {
    if len(p) > len(s) {return nil}
    pHash := hash(p)
    win := make([]int, 26)
    left := 0
    out := []int{}
    for i := 0; i < len(s); i++ {
        idx := int(s[i]-'a')
        win[idx]++
        if i-left+1 == len(p) {
            // convert win to hash string
            winSb := new(strings.Builder)
            for i := 0; i < len(win); i++ {
                winSb.WriteString(fmt.Sprintf("%v", win[i]))
                if i != len(win)-1 {winSb.WriteString("-")}
            }
            if winSb.String() == pHash {
                out = append(out, left)
            }
            
            // move left ptr out
            leftCharIdx := int(s[left]-'a')
            win[leftCharIdx]--
            left++
        }
    }
    return out
}

func hash(val string) string {
    out := make([]int, 26)
    for i := 0; i < len(val); i++ {
        idx := int(val[i]-'a')
        out[idx]++
    }
    outSb := new(strings.Builder)
    for i := 0; i < len(out); i++ {
        outSb.WriteString(fmt.Sprintf("%v", out[i]))
        if i != len(out)-1 {outSb.WriteString("-")}
    }
    return outSb.String()
}