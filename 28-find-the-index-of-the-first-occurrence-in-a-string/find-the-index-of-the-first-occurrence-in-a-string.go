/*
    approach: brute force
    - try evaluating each position in haystack string whether its the first index or not
    - for each char in haystack
        - check if we can form needle from here
    - m = len(haystack)
    - n = len(needle)
    tc = o(m*n)
    sc = o(1)

    approach: kmp
    - kmp allows us to skip checking some chars over and over again in target string
    - in brute force, we kept restarting target ptr from idx 0 whenever a mismatch occurs
    - kmp says, is a suffix in target which also exists as a prefix, then we can skip checking the prefix
        - but be greedy and try to find the largest such suffix in target string
        - that also exists as a prefix string in target string
        - if such common prefix + suffix exists, we can skip checking the prefix 
        - why?
        - because prefix == suffix
        - and suffix was a substr we JUST matched in src string
        - hence we can skip checking the prefix
        - and start our target ptr after the size of the longest common prefix window
        - This is called LPS; "longest common prefix suffix"
    - to calc LPS of a string
        - we will always try to find the largest one from the get-go
        - NOTE: LPS of a substr from idx 0-3 includes chars from 0 to 3(inclusive)
                   012345
            - eg: "aabcse"
            - lps of idx 2 means, lps of substr "aab"
        - LPS of a single char is always 0
        - we will use 2 ptrs to creating a GROWING window
        - i ptr starting at idx 1
        - j ptr starting at idx 0
        - i ptr tells us the incoming char in suffix window
        - j ptr tells us 2 things
            1. incoming char in suffix window
            2. size of the prefix window
        - if the 2 chars at both ptrs match, we move j forward
               j
            - "aa"
                i
            - here incoming char in suffix window(i) is "a"
            - and incoming char in prefix window(j) is also "a"
            - they both match, then we move j ptr forward
                j
            - "aa"
                i
            - now j ptr tells us the LONGEST PREFIX WINDOW SIZE
            - meaning the longest prefix/suffix at idx 1 is of size 1
            - therefore lps[i] = j
        - what if they dont match?
            - we shrink the prefix window
            - so move the "j" ptr back by 1 ?
            - no, instead of moving j ptr back 1 by 1,
            - we will do an optimal shrink
            - we will look at the longest prefix of prev char
            - and make j directly jump there instead.
            - eg:
                      j
                0 1 2 3 4 5 6
                a b c a b c y
                            i
               [0,0,0,1,2,3,]
            - when i is at idx 6 and j is at idx 3
            - we are trying to grow prefix/suffix window to size 4
            - by adding "a" in prefix and "y" and in suffix window
            - but the two chars do not match
            - instead of moving j back 1-by-1, trying to find a smaller window
            - we can tell j to move back to next optimal smaller window
            - which is, j will look at its prev char's lps value and go to that idx
            - lps of idx 2 = 0, so move j back to idx 0
                j
                0 1 2 3 4 5 6
                a b c a b c y
                            i
               [0,0,0,1,2,3,]
            - now do the chars match? nope
            - we cannot move j back any further, therefore lps of i(6) = 0
            - and we can move i forward , and j stays at idx 0
            - waiting to be grown again
                - when incoming char in suffix matches incoming char in prefix

    m = len(haystack)
    n = len(needle)
    tc = o(n) + o(m+n)
    sc = o(n) because of lps array
*/
func strStr(haystack string, needle string) int {
    if len(haystack) < len(needle) {return -1}
    lps := calcLPS(needle)
    i := 0
    j := 0
    for i < len(haystack) && j < len(needle) {
        if haystack[i] == needle[j] {
            i++
            j++
        } else {
            // what is the longest suffix that also exists as a prefix
            // that can be skipped checking again from the needle string
            if j > 0 {
                j = lps[j-1]
            } else {
                // no common prefix suffix exists, because j is already at idx 0
                i++
            }
        }
    }
    if j == len(needle) {return i-len(needle)}
    return -1
}

func calcLPS(word string) []int {
    n := len(word)
    lps := make([]int, n)
    i := 1 // incoming char in suffix window
    j := 0 // incoming char in prefix window, and its also telling us the size of prefix window
    for i < n {
        if word[i] == word[j] {
            j++ // prefix window size increased (i.e prefix window == suffix window )
            lps[i] = j
            i++ // evaluate the next suffix char (i ptr) with the next prefix char (j ptr)
        } else {
            // do optimal squeeze of j
            if j > 0 {
                j = lps[j-1]
            } else {
                // if j is already at 0, we do not have a common prefix/suffix
                // i.e lps[i] is 0
                // and start forming a new window by moving i forward (in hopes that the 2 ptr chars will match)
                i++
            }
        }
    }
    return lps
}


// func strStr(haystack string, needle string) int {
//     for i := 0; i < len(haystack); i++ {
//         h := i
//         n := 0
//         for h < len(haystack) && n < len(needle) {
//             if haystack[h] != needle[n] {break}
//             h++
//             n++
//         }
//         if n == len(needle) {return h-len(needle)}
//     }
//     return -1
// }