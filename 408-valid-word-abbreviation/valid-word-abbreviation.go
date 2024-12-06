func validWordAbbreviation(word string, abbr string) bool {
    n := 0
    a := 0
    w := 0
    for a < len(abbr) && w < len(word) {
        if abbr[a] >= '0' && abbr[a] <= '9' {
            if abbr[a] == '0' && n == 0 {return false}
            n = n * 10 + int(abbr[a]-'0')
            a++
            continue
        }
        for n != 0 && w < len(word) {
            w++
            n--
        }
        if w == len(word) {
            break
        }
        if abbr[a] != word[w] {return false}
        w++
        a++
    }
    for n != 0 && w < len(word) {n--; w++}
    return w == len(word) && a == len(abbr) && n == 0
}