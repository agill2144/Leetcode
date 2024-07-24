func sortJumbled(mapping []int, nums []int) []int {
    // map each num in nums to a new mapped val
    mappedNums := map[int]int{}
    for i := 0; i < len(nums); i++ {
        nStr := fmt.Sprintf("%v",nums[i])
        newMappedVal := 0
        for j := 0; j < len(nStr); j++ {
            dig := nStr[j]
            digInt := int(dig-'0')
            newMappedVal = newMappedVal * 10 + mapping[digInt]
        }
        mappedNums[nums[i]] = newMappedVal
    }
    // sort original nums based on its mapped value
    // if 2 mapped values are same, keep the original order
    // hence using "sort.SliceStable" over "sort.Slice"
    sort.SliceStable(nums, func(i, j int)bool{
        iVal := nums[i]
        iMappedVal := mappedNums[iVal]
        jVal := nums[j]
        jMappedVal := mappedNums[jVal]
        
        if iMappedVal == jMappedVal {
            return i < j
        }

        return iMappedVal < jMappedVal
    })
    return nums
}