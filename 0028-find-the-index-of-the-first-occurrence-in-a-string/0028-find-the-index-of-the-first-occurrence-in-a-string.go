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
    lps := findLPS(needle)
    i := 0
    j := 0
    for i < len(haystack) {
        if haystack[i] == needle[j] {
            j++
            i++
            if j == len(needle) {
                return i-len(needle)
            }
        } else if j > 0 {
            // move j back but match longest common prefix(using needle)
            // with longest common suffix(using haystack)
            j = lps[j-1]
        } else {
            // j has reched 0th idx,
            // start matching from start of needle again
            i++
        }
    }
    return -1
}

func findLPS(needle string) []int {
    out := make([]int, len(needle))
    i := 1
    j := 0
    for i < len(needle) {
        if needle[i] == needle[j] {
            j++
            out[i] = j
            i++
        } else if j > 0 {
            j = out[j-1]
        } else if j == 0 {
            i++
        }
    }
    return out
}

/*
    approach: brute force
    - form needle size substring at each char of haystack
    
    time = o(h) * o(n)
    space = o(n) ; when forming needle size substr
*/
// func strStr(haystack string, needle string) int {
//     for i := 0; i <= len(haystack)-len(needle); i++ {
//         subStr := haystack[i:i+len(needle)]
//         if subStr == needle {return i}
//     }
//     return -1
// }