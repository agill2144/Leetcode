// sort on ;
// whatever concatnation results into a bigger num
// its either i+j or j+i
func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int)bool{
        ij, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[i], nums[j]))
        ji, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[j], nums[i]))
        return ij > ji
    })
    res := new(strings.Builder)
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 && len(res.String()) == 0 {continue}
        res.WriteString(fmt.Sprintf("%v", nums[i]))
    }
    if res.String() == "" {return "0"}
    return res.String()
}