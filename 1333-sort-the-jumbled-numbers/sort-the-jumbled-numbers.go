func sortJumbled(mapping []int, nums []int) []int {
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