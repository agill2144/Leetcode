// for each ith element, do twoSum using complement search
// what will be the target for the twoSum complement search; 
//  - what number do we need to add to ith element for it to equal to 0
//  - 0-ithNumber
// so our twoSum complement search will search for 0-ithNumber ( i being the outerloop idx )
// time : o(n) * o(n) = o(n^2)
func threeSum(nums []int) [][]int {
    out := [][]int{}
    set := map[int]struct{}{}
    outSeen := map[string]struct{}{}
    
    for i := 0; i < len(nums); i++ {
        outterNum := nums[i]
        if _, ok := set[outterNum]; ok {continue}
        set[outterNum] = struct{}{}
        target := 0-outterNum
        
        seen := map[int]struct{}{}
        for j := i+1; j < len(nums); j++ {
            compl := target-nums[j]  
            if _, ok := seen[compl]; ok {
                triplet := []int{outterNum, nums[j], compl}
                sort.Ints(triplet)
                if _, ok := outSeen[fmt.Sprintf("%v", triplet)]; ok {continue}
                outSeen[fmt.Sprintf("%v", triplet)] = struct{}{}
                out = append(out, []int{outterNum, nums[j], compl})
            }
            seen[nums[j]] = struct{}{}
        }
    } 
    return out
}

