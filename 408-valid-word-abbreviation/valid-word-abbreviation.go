func validWordAbbreviation(word string, abbr string) bool {
    w, a := 0, 0
    n := 0
    for w < len(word) && a < len(abbr) {
        if abbr[a] >= '0' && abbr[a] <= '9' {
            if n == 0 && abbr[a] == '0' {return false}
            n = n * 10 + int(abbr[a]-'0')
            a++
            continue
        }
        for n > 0 {
            w++
            n--
        }
        if w == len(word) {break}

        if abbr[a] != word[w] {return false}
        a++; w++
    }
    for n > 0 && w < len(word) {
        n--
        w++
    }
    return n == 0 && w == len(word) && a == len(abbr)
}

// i12izatio1
//.         a
// internationalization
//                    w