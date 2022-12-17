// two pointers
// once we visited all characters of s from left to right, we have a valid subsequence s in t
// time: o(t) if t is the bigger string
// space: o(1)
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