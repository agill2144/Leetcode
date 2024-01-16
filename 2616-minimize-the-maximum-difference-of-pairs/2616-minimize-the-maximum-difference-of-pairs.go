func minimizeMax(nums []int, p int) int {
    // such that the maximum difference amongst all the pairs is minimized
    // max ans is minimized
    sort.Ints(nums)
    ans := 0
    left := 0
    right := nums[len(nums)-1]
    for left <= right {
        
        // i.e atMax diff 
        mid := left + (right-left)/2

        // evaluate mid
        // count pairs whose diff <= $mid
        count := 0
        for i := 0; i < len(nums)-1; i++ {
            curr := nums[i]
            next := nums[i+1]
            diff := next-curr
            if diff <= mid {
                count++
                i++
            }
        }
        
        // does mid work?
        // mid works if we were able to make atleast p pair ( or more ) but atleast p
        if count >= p {
            ans = mid
            // if worked, save it, search left -- everything to the right of this will work
            // we want smallest such ans
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}


// func minimizeMax(nums []int, p int) int {
//     // such that the maximum difference amongst all the pairs is minimized
//     // max ans is minimized
//     sort.Ints(nums)
//     ans := 0
//     for i := 0; i <= nums[len(nums)-1]; i++ {

//         count := 0
//         atMax := i
//         j := 0
//         for j < len(nums)-1 {
//             curr := nums[j]
//             next := nums[j+1]
//             diff := next-curr
//             if diff <= atMax { 
//                 count++
//                 j+=2
//             } else {
//                 j++
//             }
//         }
//         if count < p {continue}
//         return i
//     }
//     return ans
// }