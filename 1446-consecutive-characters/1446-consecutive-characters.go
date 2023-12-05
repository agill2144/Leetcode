func maxPower(s string) int {
    idx := map[byte]int{}
    left := 0
    max := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        idx[char] = i
        if len(idx) == 1 {
            size := i-left+1
            if size > max {max = size}
        } else {
            leftChar := s[left]
            leftCharLastIdx := idx[leftChar]
            left = leftCharLastIdx + 1
            delete(idx, leftChar)
        }
    } 
    return max
}