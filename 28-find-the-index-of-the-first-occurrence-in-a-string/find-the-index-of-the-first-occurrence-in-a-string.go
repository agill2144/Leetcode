func strStr(src string, target string) int {
    if len(src) < len(target) {return -1}
    lps := calcLPS(target)
    i, j := 0, 0
    for i < len(src) && j < len(target) {
        if src[i] == target[j] {
            i++
            j++
        } else {
            // what is a longest COMMON suffix that is also a prefix
            // that we can skip checking because they are equal
            if j != 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    if j == len(target) {return i-len(target)}
    return -1
}

func calcLPS(word string) []int{
    n := len(word)
    lps := make([]int, n)
    i := 1
    // j does 2 things
    // 1. its position = len of prefix / suffix window
    // 2. its also the incoming char in prefix win that can be compared with incoming char (i) in suffix window
    j := 0
    for i < n {
        if word[i] == word[j] {
            j++
            lps[i] = j
            i++
        } else {
            if j != 0 {
                j = lps[j-1]
            } else {
                i++
            }
        }
    }
    return lps
    
}