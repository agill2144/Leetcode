func longestContinuousSubstring(s string) int {
    slow := 0
    fast := 1
    ans := 1
    for fast < len(s) {
        curr := s[fast]
        prev := s[fast-1]
        if curr-1 == prev {
            if fast-slow+1 > ans {ans = fast-slow+1}
        } else {
            slow = fast            
        }
        fast++
    }
    return ans
}