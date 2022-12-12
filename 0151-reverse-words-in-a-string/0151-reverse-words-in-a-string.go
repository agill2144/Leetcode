func reverseWords(s string) string {
    trimmedS := strings.TrimSpace(s)
    outBldr := new(strings.Builder)
    split := strings.Split(trimmedS, " ")
    for i := len(split)-1; i>=0 ;i-- {
        word := split[i]
        if len(word) == 0 {continue}
        outBldr.WriteString(word)
        if i != 0 {outBldr.WriteString(" ")}
    }
    return outBldr.String()
}
