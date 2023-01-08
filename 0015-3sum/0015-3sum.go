// for each ith element, do twoSum using complement search
// what will be the target for the twoSum complement search; 
//  - what number do we need to add to ith element for it to equal to 0
//  - 0-ithNumber
// so our twoSum complement search will search for 0-ithNumber ( i being the outerloop idx )
// time : o(n) * o(n) = o(n^2)
func threeSum(nums []int) [][]int {
    out := [][]int{}
    set := map[int]struct{}{}
    outSeen := map[[3]int]struct{}{}
    
    for i := 0; i < len(nums); i++ { // o(n)
        outterNum := nums[i]
        if _, ok := set[outterNum]; ok {continue}
        set[outterNum] = struct{}{}
        target := 0-outterNum
        
        seen := map[int]struct{}{}
        for j := i+1; j < len(nums); j++ { // o(n)
            compl := target-nums[j]  
            if _, ok := seen[compl]; ok {
                triplet := []int{outterNum, nums[j], compl}
                sort.Ints(triplet) // (o3log3)
                tmp := [3]int{triplet[0], triplet[1], triplet[2]} 
                if _, ok := outSeen[tmp]; !ok {
                    outSeen[tmp] = struct{}{}
                    out = append(out, triplet)
                }
            }
            seen[nums[j]] = struct{}{}
        }
    } 
    // total time = o(n) x o(n) x o(3log3) = o(n^2)
    // space = o(n) + o(n) + o(n) ; 3 sets
    return out
}

