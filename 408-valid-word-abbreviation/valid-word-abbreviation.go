func validWordAbbreviation(word string, abbr string) bool {
    w := 0; a := 0
    num := 0
    for a < len(abbr) && w < len(word) {
        if abbr[a] >= '0' && abbr[a] <= '9' {
            if num == 0 && abbr[a] == '0' {return false}
            num = num * 10 + int(abbr[a]-'0')
            a++
            continue
        }
        for num > 0 {
            w++
            num--
        }
        if w == len(word){ break }
        
        if abbr[a] != word[w] { return false }
        a++
        w++
    }
    for num > 0 && w < len(word) {num--; w++}
    return a == len(abbr) && w == len(word) && num == 0
}