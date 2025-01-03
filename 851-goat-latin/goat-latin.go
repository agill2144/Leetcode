func toGoatLatin(sentence string) string {
    /*
        vowel
        - append ma
        - append a's suffix based on idx

        consonant
        - take substr; word[1:]
        - add the 0th char to the end
        - add ma
        - append a's suffix based on idx
    */
    res := new(strings.Builder)
    suffix := "a"
    words := strings.Split(sentence, " ")
    for i := 0; i < len(words); i++ {
        word := words[i]
        if isVowel(word) {
            res.WriteString(word)
            res.WriteString("ma")
        } else {
            res.WriteString(word[1:])
            res.WriteByte(word[0])
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