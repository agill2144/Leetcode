func replaceWords(dictionary []string, sentence string) string {
    set := map[string]struct{}{}
    for i := 0; i < len(dictionary); i++ {set[dictionary[i]] = struct{}{}}
    out := new(strings.Builder)
    words := strings.Split(sentence, " ")
    for i := 0; i < len(words); i++ {
        word := words[i]
        found := false
        for j := 0; j < len(word); j++ {
            subStr := word[:j+1]
            if _, ok := set[subStr]; ok {
                found = true
                out.WriteString(subStr)
                break
            }
        }
        if !found {out.WriteString(word)}
        if i != len(words)-1 {out.WriteString(" ")}
    }
    return out.String()
}