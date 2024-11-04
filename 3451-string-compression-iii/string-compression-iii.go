func compressedString(word string) string {
    if len(word) == 0 {return ""}    
    res := &strings.Builder{}
    i := 0
    for i < len(word) {
        count := 0
        startChar := word[i]
        for i == 0 || (i < len(word) && word[i] == startChar && count < 9) {
            i++
            count++
        }
        res.WriteString(fmt.Sprintf("%v%v", count, string(startChar)))
        
    }
    return res.String()
}

/*
    char = 'a'
    count = 1

    abcde
      i
*/