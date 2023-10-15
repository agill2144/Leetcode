// The most magical fucking algo : KMP
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
        } else if haystack[i] != needle[j] && j > 0 {
            /*
                When you encounter a mismatch at haystack[i] and needle[j],
                We have already matched ALL prev chars upto this idx ( both in haystack and needle )
                
                We need to find out whether there are chars that appear both in the front and back of the prev substr
                Does the substr have a prefix thats also a suffix 
                why?
                *** because we will use prefix from pattern and suffix from haystack ****
                the LPS array is telling us the LPS itself
                if all chars upto this idx have matched, 
                    we need to know which chars from pattern can be re-mapped to just immediate chars of src string
                    i.e what is the longest prefix / suffix of prev char
                    therefore lps[j-1]
                
            */
            j = lps[j-1]
        } else if haystack[i] != needle[j] && j == 0 {
            // j is already at 0th idx, cant move that back anymore
            i++
        }
    }
    return -1
}

func findLPS(needle string) []int {
    out := make([]int, len(needle))
    j := 0
    i := 1
    for i < len(needle) {
        if needle[i] == needle[j] {
            j++
            out[i] = j
            i++
        } else if needle[j] != needle[i] && j > 0 {
            j = out[j-1]
        } else if needle[j] != needle[i] && j == 0 {
            i++
        }
    }
    return out
}