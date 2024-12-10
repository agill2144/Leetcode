func maximumLength(s string) int {
    left := 1
    right := len(s)-2
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        // mid is our len
        // evaluate each substring of size mid
        found := false
        freq := map[string]int{}
        for i := 0; i+mid < len(s)+1; i++ {
            subStr := s[i:i+mid]
            expected := ""
            for k := i; k < i+mid; k++ {expected += string(s[i])}
            if subStr == expected {
                freq[subStr]++
                if freq[subStr] == 3 {found = true; break}
            }
        }
        if found {
            ans = mid
            left = mid+1   
        } else {
            right = mid-1
        }
    }
    return ans
}