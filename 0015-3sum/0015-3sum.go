

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        
        // skip ith value which is a dupe of previous
        if i != 0 && nums[i] == nums[i-1] {continue}
        
        // 2Sum two-ptr from here
        left := i+1
        right := len(nums)-1
        for left < right {
            sum := nums[left] + nums[right] + nums[i]
            if sum == 0 {
                out = append(out, []int{nums[i], nums[left], nums[right]})
                // we have used both values, move both ptrs
                left++
                right--
                // escape dupes
                for left < right && nums[left] == nums[left-1] {left++}
                for left < right && nums[right] == nums[right+1] {right--}
                
            } else if sum > 0 {
                // sum is too big, move away from bigger number ( i.e right ptr since arr is sorted )
                right--
            } else {
                // sum is too small, move away from smaller number
                left++
            }
        }
    }
    return out
}


// complement search
// time = o(n) * o(n) = o(n^2)
// space = o(k) for tracking uniq triplets * o(n) for 2sum complement search set = o(nk)
// func threeSum(nums []int) [][]int {
//     // To keep track of uniq triplets
//     // k uniq triplets
//     // o(k) space
//     uniqTripSet := map[[3]int]struct{}{}
//     // we are not using a idx map, becuase we have to form triplets using their values instead of idx
//     ans := [][]int{}
//     // o(n) time
//     for i := 0; i < len(nums); i++ {
//         ithVal := nums[i]
//         target := 0-ithVal
        
//         // why is this set inside ?
//         // this is 2sum from here
//         // in 2sum, we needed idx of pairs
//         // in this, we need values, therefore set
//         // space = o(n)
//         set := map[int]struct{}{}
        
//         // inner loop , time = o(n) * o(3log3) = o(n)
//         for j := i+1; j < len(nums); j++ {
//             jthVal := nums[j]
//             diff := target-jthVal
//             if _, ok := set[diff]; ok {
//                 triplet := []int{ithVal,jthVal,diff}
//                 sort.Ints(triplet)
//                 tmp := [3]int{triplet[0], triplet[1], triplet[2]}
//                 if _, ok := uniqTripSet[tmp]; !ok {
//                     ans = append(ans, triplet) 
//                     uniqTripSet[tmp] = struct{}{}
//                 }
//             }
//             set[jthVal] = struct{}{}
//         }
//     }
//     return ans
// }