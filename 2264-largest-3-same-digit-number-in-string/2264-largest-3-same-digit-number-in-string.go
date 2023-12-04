// sliding window ; form all 3 size number, save the max
func largestGoodInteger(num string) string {
    maxNum := ""
    curr := []string{}
    left := 0
    window := map[string]int{}
    for i := 0; i < len(num); i++ {
        
        curr = append(curr, string(num[i]))
        window[string(num[i])]++

        if i - left + 1 == 3 {
            leftChar := string(num[left])
            if len(window) == 1 {
                currStr := strings.Join(curr, "")
                if maxNum == "" || strings.Compare(currStr,maxNum) == 1 {
                    maxNum = currStr
                }
            }
            window[leftChar]--
            curr = curr[1:]
            if val := window[leftChar]; val == 0 {delete(window, leftChar)}
            left++
        }
    }
    return maxNum
}