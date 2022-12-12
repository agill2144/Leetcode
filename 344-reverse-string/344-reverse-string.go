// simple swap using 2 ptrs, o(n) time and o(1) space
func reverseString(s []byte)  {
    p, l := 0, len(s) - 1
    
    for p < l {
        s[p], s[l] = s[l], s[p]
        p++
        l--
    }
}