func twoEditWords(queries []string, dictionary []string) []string {
    out := []string{}
    for i := 0; i < len(queries); i++ {
        query := queries[i]
        for j := 0; j < len(dictionary); j++ {
            dict := dictionary[j]
            if query == dict {out = append(out, query); break}

            q := 0
            d := 0
            countDiff := 0
            for q < len(query) && d < len(dict) {
                if query[q] != dict[d] {countDiff++}
                q++
                d++
            }
            if countDiff <= 2 {out = append(out, query); break}
        }
    }
    return out
}