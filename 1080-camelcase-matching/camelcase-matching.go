func camelMatch(queries []string, pattern string) []bool {
    out := make([]bool, len(queries))
    for idx, word := range queries {
        w, p := 0,0
        // everything in pattern must match, we can only insert chars into pattern
        res := true
        for p < len(pattern) && w < len(word) && res {
            // if both chars are upper case, they both must match!
            if word[w] <= 'Z' && pattern[p] <= 'Z' {
                if pattern[p] != word[w] {res = false}
                w++
                p++
            } else if word[w] != pattern[p] {
                // if both are not uppercase, one of them is , or none of them are
                // we can only move w ptr forward, until word char matches pattern char
                // "everything in pattern must match, we can only insert chars into pattern"
                // therefore move only w ptr forward in hopes for a match
                w++
            } else {
                // the happy case, when both chars do match
                w++
                p++
            }
        }
        // "everything in pattern must match, we can only insert chars into pattern"
        // if pattern still has chars left, aint no way this is a match
        if p < len(pattern) && res { res = false }

        // if word be bigger than pattern and pattern matched everything this far
        // ensure word does not have upper cases
        // because we are only allowed to insert lower case chars into pattern
        for w < len(word) && res {
            if word[w] <= 'Z' { res = false }
            w++
        }

        out[idx] = res
    }
    return out
}