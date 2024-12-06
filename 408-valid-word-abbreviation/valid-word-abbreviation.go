func validWordAbbreviation(word string, abbr string) bool {
    w := 0
    a := 0
    n := 0
    for w < len(word) && a < len(abbr) {
        for a < len(abbr) && abbr[a] >= '0' && abbr[a] <= '9' {
            if abbr[a] == '0' && n == 0 {return false}
            n = n * 10 + int(abbr[a]-'0')
            a++
        }
        for n != 0 && w < len(word) {
            n--
            w++
        }
        if a == len(abbr) || w == len(word) {break} 
        if abbr[a] != word[w] {return false}
        a++; w++
    }
    for n != 0 && w < len(word) {n--; w++}
    return n == 0 && a == len(abbr) && w == len(word)
}