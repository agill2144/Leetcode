func camelMatch(queries []string, pattern string) []bool {
    out := make([]bool, len(queries))
    for idx, word := range queries {
        w, p := 0,0
        // everything in pattern must match, we can only insert chars into pattern
        res := true
        for p < len(pattern) && w < len(word) && res {
            if word[w] <= 'Z' && pattern[p] <= 'Z' {
                if pattern[p] != word[w] {res = false}
                w++
                p++
            } else if word[w] != pattern[p] {
                w++
            } else {
                w++
                p++
            }
        }
        if !res { continue }
        
        // if word be bigger than pattern and pattern matched everything this far
        // ensure word does not have upper cases
        for w < len(word) && res {
            if word[w] <= 'Z' { res = false }
            w++
        }
        // found upper cases in above case
        if !res {continue}

        if p < len(pattern) {
            out[idx] = false
        } else {
            out[idx] = true
        }
    }
    return out
}