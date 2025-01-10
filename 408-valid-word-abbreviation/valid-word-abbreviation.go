func validWordAbbreviation(word string, abbr string) bool {
    n := 0
    w, a := 0, 0
    for w < len(word) && a < len(abbr) {
        for a < len(abbr) && abbr[a] >= '0' && abbr[a] <= '9' {
            if abbr[a] == '0' && n == 0 {return false}
            n = n*10+int(abbr[a]-'0')
            a++
        }
        if n != 0 {
            w += n
            n = 0
        }
        if a >= len(abbr) || w >= len(word) {
            break
        }
        if abbr[a] != word[w] {return false}
        a++
        w++
    }
    if n != 0 {
        w += n
        n = 0
    }
    return n == 0 && w == len(word) && a == len(abbr)
}