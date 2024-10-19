func countGoodSubstrings(s string) int {
    idx := map[byte]int{}
    count := 0
    left := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        val, ok := idx[char]
        if ok && left <= val {
            left = val+1
        }
        idx[char] = i
        if i-left+1 == 3 {
            count++
            leftChar := s[left]
            if val, ok := idx[leftChar]; ok && val <= left {
                delete(idx, leftChar)
            }
            left++
        }
    }
    return count
}