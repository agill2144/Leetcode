func largestPalindromic(num string) string {
    freq := make([]int, 11) // idx -> num mapping
    for i := 0; i < len(num); i++ {
        char := num[i]
        freq[int(char-'0')]++
    }
    out := new(strings.Builder)
    var middle string
    back := []string{}
    for i := len(freq)-1; i >= 0; i-- {
        if freq[i] <= 1 {
            continue
        }
        count := freq[i] / 2
        freq[i] = freq[i] % 2
        for k := 0; k < count; k++ {
            if out.String() == "" && i == 0 {continue}
            out.WriteString(fmt.Sprintf("%v", i))
            back = append(back, fmt.Sprintf("%v", i))
        }
    }
    for i := len(freq)-1; i >= 0; i-- {
        if freq[i] >= 1 {
            middle = fmt.Sprintf("%v", i)
            break
        }
    }
    out.WriteString(middle)
    for i := len(back)-1; i >= 0; i-- {
        out.WriteString(back[i])
    }
    if out.String() == "" {out.WriteString("0")}
    return out.String()
    
}
