
// merge sort ; recursively
func largestNumber(nums []int) string {
    var dfs func(left, right int)
    dfs = func(left, right int) {
        // base
        if left >= right {return}

        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)

        merged := []int{}
        i, j := left, mid+1
        for i <= mid && j <= right {
            ij, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[i], nums[j]))
            ji, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[j], nums[i]))
            if ij > ji {
                merged = append(merged, nums[i])
                i++
            } else {
                merged = append(merged, nums[j])
                j++
            }
        }
        for i <= mid {merged = append(merged, nums[i]);i++}
        for j <= right {merged = append(merged, nums[j]);j++}
        ptr := 0
        for i := left; i <= right; i++ {
            nums[i] = merged[ptr]
            ptr++
        }
    }
    dfs(0, len(nums)-1)
    res := new(strings.Builder)
    for i := 0; i < len(nums); i++ {
        // edge case when constructing numbers: leading zeros
        // in this case, if nums[i] is 0 and we do not have anything in our result
        // yet, drop this 0 because it creates an invalid number
        if nums[i] == 0 && len(res.String()) == 0 {continue}
        res.WriteString(fmt.Sprintf("%v", nums[i]))
    }
    if res.String() == "" {return "0"}
    return res.String()
}

// sort on ;
// whatever concatnation results into a bigger num
// its either i+j or j+i
// func largestNumber(nums []int) string {
//     sort.Slice(nums, func(i, j int)bool{
//         ij, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[i], nums[j]))
//         ji, _ := strconv.Atoi(fmt.Sprintf("%v%v", nums[j], nums[i]))
//         return ij > ji
//     })
//     res := new(strings.Builder)
//     for i := 0; i < len(nums); i++ {
//         if nums[i] == 0 && len(res.String()) == 0 {continue}
//         res.WriteString(fmt.Sprintf("%v", nums[i]))
//     }
//     if res.String() == "" {return "0"}
//     return res.String()
// }