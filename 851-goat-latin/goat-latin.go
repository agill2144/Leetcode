func toGoatLatin(sentence string) string {
    words := strings.Split(sentence, " ")
    res := new(strings.Builder)
    suffix := "a"
    for i := 0; i < len(words); i++ {
        if isVowel(words[i]) {
            res.WriteString(words[i])
            res.WriteString("ma")
        } else {
            res.WriteString(words[i][1:])
            res.WriteByte(words[i][0])
            res.WriteString("ma")
        }
        res.WriteString(suffix)
        suffix += "a"
        if i != len(words)-1 {res.WriteString(" ")}
    }
    return res.String()

}
func isVowel(s string) bool {
    set := map[byte]bool{
        'a':true,'e':true,'i':true,'o':true,'u':true,
        'A':true,'E':true,'I':true,'O':true,'U':true,
    }
    return set[s[0]]
}
