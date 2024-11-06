
func findDisappearedNumbers(nums []int) []int {
    // 1  to n ( inclusize )
    // in a sorted array of values 1 to n ( inclusive )
    // without any missing number , idx for each val will be val-1
    // val 1 = idx 0 ( 1-1 )
    // val 4 = idx 3 ( 4-1 )
    // therefore we can mark idx positions of each val that we have seen so far
    // we will make the value at correct idx position ( val *= -1 )
    // the missing numbers will be idxs whose values are still positive 
    n := len(nums)
    for i := 0; i < n; i++ {
        absVal := abs(nums[i])
        idx := absVal - 1
        if nums[idx] < 0 {continue}
        nums[idx] = -nums[idx]
    }
    out := []int{}
    for i := 0; i < n; i++ {
        if nums[i] > 0 {out = append(out, i+1)}
    }
    return out
}

func abs (n int) int {
    if n < 0 {return -n}
    return n
}
// func findDisappearedNumbers(nums []int) []int {
//     set := map[int]bool{}
//     for i := 0; i < len(nums); i++ {set[nums[i]] = true}
//     exp := 1
//     out := []int{}
//     for i := exp; i <= len(nums); i++ {
//         if !set[i] {out = append(out, i)}
//     }
//     return out
// }