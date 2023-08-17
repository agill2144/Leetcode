func threeSum(nums []int) [][]int {
    // To keep track on uniq triplets
    uniqTripSet := map[[3]int]struct{}{}
    // we are not using a idx map, becuase we have to form triplets using their values instead of idx
    ans := [][]int{}
    for i := 0; i < len(nums); i++ {
        ithVal := nums[i]
        target := 0-ithVal
        set := map[int]struct{}{}
        for j := i+1; j < len(nums); j++ {
            jthVal := nums[j]
            diff := target-jthVal
            if _, ok := set[diff]; ok {
                triplet := []int{ithVal,jthVal,diff}
                sort.Ints(triplet)
                tmp := [3]int{triplet[0], triplet[1], triplet[2]}
                if _, ok := uniqTripSet[tmp]; !ok {
                    ans = append(ans, triplet) 
                    uniqTripSet[tmp] = struct{}{}
                }
            }
            set[jthVal] = struct{}{}
        }
    }
    return ans
}