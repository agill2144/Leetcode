// two pointers
// once we visited all characters of s from left to right, we have a valid subsequence s in t
func isSubsequence(s string, t string) bool {
    sPtr := 0
    tPtr := 0
    for sPtr < len(s) && tPtr < len(t) {
        if s[sPtr] == t[tPtr] {
            sPtr++
        }
        tPtr++
    }
    return sPtr == len(s)
}