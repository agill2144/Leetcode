func numberOfSubstrings(s string) int {
    k := 3
    n := len(s)
    total := n*(n+1)/2
    return total - countLessThanOrEqualTo(s, k-1)
}

func countLessThanOrEqualTo(s string, k int) int {
    left := 0
    count := 0
    freq := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        if k != 3 {for len(freq) > k {
            leftChar := s[left]
            freq[leftChar]--
            if freq[leftChar] == 0 {delete(freq, leftChar)}
            left++
        }
        }
        count += (i-left+1)
    }
    return count
}

// func numberOfSubstrings(s string) int {
//     a := 0
//     b := 0
//     c := 0
//     left := 0
//     n := len(s)
//     count := 0
//     for i := 0; i < n; i++ {
//         if s[i] == 'a' {a++}
//         if s[i] == 'b' {b++}
//         if s[i] == 'c' {c++}
//         for a > 0 && b > 0 && c > 0 {
//             // ending at idx i , we have a valid substr
//             // it means, anything else we add from right of i, will not make this invalid
//             // therefore we find the end position of a valid window
//             // and then do a count number of substrings we 
//             count += (n-1-i+1)
//             if s[left] == 'a' {a--}
//             if s[left] == 'b' {b--}
//             if s[left] == 'c' {c--}
//             left++
//         }
//     }
//     return count
// }