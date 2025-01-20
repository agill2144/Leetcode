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
    length := 0
    for i < n {
        if word[i] == word[length] {
            length++
            lps[i] = length
            i++
        } else {
            if length != 0 {
                length = lps[length-1]
            } else {
                i++
            }
        }
    }
    return lps
}