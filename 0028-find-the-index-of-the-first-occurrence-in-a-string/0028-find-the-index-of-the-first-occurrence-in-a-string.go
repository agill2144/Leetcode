/*

Pattern Preprocessing (Building the LPS Array):
- KMP starts with a preprocessing step to construct the Longest Prefix Suffix (LPS) array. 
- The LPS array is key to understanding where to shift the pattern during the search phase.
- The LPS array for a pattern of length "m" is an array of length "m" where lps[i] stores the length of the longest common prefix of the pattern that is also the longest common suffix up to position i.

Search Phase:
- In the search phase, KMP starts comparing the characters of the pattern and the text from left to right.
- When a mismatch is found at position i in the pattern, 
    instead of simply shifting the pattern to the right by 1 character (as in brute force), 
    KMP uses the LPS array to determine how far it can safely shift the pattern.
- The LPS array tells us the length of the longest common prefix and suffix up to position i in the pattern. 
- This means that the first lps[i] characters in the pattern are already known to match with the text.
- By shifting the pattern to the right by i - lps[i] characters, KMP ensures that it continues the comparison at the next character in the text that hasn't been compared yet with the pattern.

*/
func strStr(haystack string, needle string) int {
    lps := getLPS(needle)
    i := 0
    j := 0
    for i < len(haystack) {
        if haystack[i] == needle[j] {
            i++
            j++
            if j == len(needle) {
                return i-len(needle)
            }
        } else if haystack[i] != needle[j] && j > 0 {
            j = lps[j-1]
        } else {
            i++
        }
    }
    return -1
}


func getLPS(pattern string) []int {
    lps := make([]int, len(pattern))
    i := 1
    j := 0
    for i < len(pattern) {
        if pattern[j] == pattern[i] {
            j++
            lps[i] = j
            i++
        } else if pattern[j] != pattern[i] && j > 0{
            j = lps[j-1]
        } else {
            i++
        }
    }
    return lps
}