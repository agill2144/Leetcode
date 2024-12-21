func validWordAbbreviation(word string, abbr string) bool {
    w, a := 0, 0
    n := 0
    for w < len(word) && a < len(abbr) {
        for a < len(abbr) && abbr[a] >= '0' && abbr[a] <= '9' {
            if abbr[a] == '0' && n == 0 {return false}
            n = n * 10 + int(abbr[a]-'0')
            a++
        }

        w += n
        n = 0
        if a == len(abbr) || w == len(word) {break}

        if word[w] != abbr[a] {return false}
        a++
        w++
    }

    w += n
    n = 0
    return a == len(abbr) && w == len(word)
}