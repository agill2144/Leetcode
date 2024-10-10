func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int) bool {
        iNum := nums[i]
        jNum := nums[j]
        ij := fmt.Sprintf("%v%v",iNum, jNum)
        ji := fmt.Sprintf("%v%v", jNum, iNum)
        ijInt, _ := strconv.Atoi(ij)
        jiInt, _ := strconv.Atoi(ji)
        // ith num should be placed first
        // if i+j > j+i
        return ijInt > jiInt
    })
    out := new(strings.Builder)
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 && len(out.String()) == 0 {continue}
        out.WriteString(fmt.Sprintf("%v", nums[i]))
    }
    if out.String() == "" {return "0"}
    return out.String()
}