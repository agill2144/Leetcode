func maximizeSweetness(sweetness []int, k int) int {
    // min ans is maximized
    left := 1
    right := 0
    for i := 0; i < len(sweetness); i++ {
        right += sweetness[i]
    }
    ans := -1
    for left <= right {
        // if a chunk at-min is allowed $mid sweetness
        // how many chunks can we get ( all chunks must satisfy atMin sweetness level )
        mid := left + (right-left)/2

        subArrCount := 0
        sum := 0
        for i := 0; i < len(sweetness); i++ {
            sum += sweetness[i]
            // if atMin is met, cut the bar here
            if sum >= mid {
                subArrCount++
                sum = 0
            }
        }

        // did $mid work?
        // $mid will only work if we had atleast k+1 chunks with $atMin sweetness level
        // why k+1 ? because we have to include ourselves, i.e there are k+1 total people and not k..
        if subArrCount >= k+1 {
            // we have atleast k+1 chunks that meet $atMin sweetness level
            // save this ans as potentianl ans
            ans = mid
            // search right, because we want largest such ans
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans

}
// func maximizeSweetness(sweetness []int, k int) int {
//     // min ans is maximized
//     start := 1
//     end := 0
//     for i := 0; i < len(sweetness); i++ {
//         end += sweetness[i]
//     }
//     ans := -1
//     for i := start; i <= end; i++ {
//         atMin := i
//         subarrayCount := 0
//         sum := 0
//         for j := 0; j < len(sweetness); j++ {
//             sum += sweetness[j]
//             if sum >= atMin {
//                 subarrayCount++
//                 sum = 0
//             }
//         }

//         // its okay to have more than k+1 chunks, we can apparently throw away extra left overs :shrug
//         // thats why >= k+1 and not == k+1
//         if subarrayCount >= k+1 {
//             ans = atMin
//         } else {
//             break
//         }
//     }
//     return ans
// }