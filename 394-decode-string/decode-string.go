func decodeString(s string) string {
    strs := []*strings.Builder{}
    nums := []int{}
    curr := new(strings.Builder)
    currNum := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        if char >= '0' && char <= '9' {
            currNum = currNum * 10 + int(char-'0')
        } else if char >= 'a' && char <= 'z' {
            curr.WriteByte(char)
        } else if char == '[' {
            nums = append(nums, currNum)
            strs = append(strs, curr)
            curr = new(strings.Builder)
            currNum = 0
        } else if char == ']' {
            // repeat curr string $top num times
            k := nums[len(nums)-1]
            nums = nums[:len(nums)-1]
            decoded := new(strings.Builder)
            for j := 0; j < k; j++ {
                decoded.WriteString(curr.String())
            }
            // concat $top-of-strs + decoded
            curr = new(strings.Builder)
            curr.WriteString(strs[len(strs)-1].String())
            curr.WriteString(decoded.String())
            strs = strs[:len(strs)-1]
        }
    }
    return curr.String()
}