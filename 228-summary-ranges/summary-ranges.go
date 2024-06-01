func summaryRanges(nums []int) []string {
    
    out := []string{}
    i := 0
    for i < len(nums) {
        tmp := new(strings.Builder)
        start := i
        for i < len(nums) && i+1 < len(nums) && nums[i]+1 == nums[i+1] {
            i++
        }
        startN := nums[start]
        endN := nums[i]
        if startN != endN {
            tmp.WriteString(fmt.Sprintf("%v->%v", startN, endN))
        } else {
            tmp.WriteString(fmt.Sprintf("%v", endN))
        }
        out = append(out, tmp.String())
        i++
    }
    return out
}