/*
    
*/
func minimizeMax(nums []int, p int) int {
    // max diff is minimized
    sort.Ints(nums)
    left := 0
    right := nums[len(nums)-1]-nums[0]
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        
        count := 0
        i := 0
        for i < len(nums)-1 {
            curr := nums[i]
            next := nums[i+1]
            diff := next-curr
            if diff <= mid {  // <= atMax allowed diff
                count++
                i+=2 // move 2 positions, since this idx and next idx cannot be used again
            } else {
                // otherwise move only 1 position
                i++
            }
        }
        
        if count >= p {
            // if we were able to create atleast p pairs ( can be more )
            // it means, the next set of numbers will work too
            // therefore no need to search right, save this ans as potential ans
            // and continue searching left
            // since we want smallest such ans
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

/*
    "such that the maximum difference amongst all the pairs is minimized"
    - binary search on answers hint
    - max ans is minimized
   
    approach: brute force
    - we want the smallest difference between 2 nums
    - the smallest difference between any 2 nums will always be the numbers that are sequentially next to each other
    - therefore lets sort the numbers first
   
    - now try all possible differences start from 0
    - but end at what?
        - whats the max difference possible ?
        - difference between largest-smallest number
    - so a range loop evaluating how many pairs would we have if 
        - atMax the allowed difference is $i
        - can be <= , but never more
    
    - if we get back more than p pairs
        - return $i and exit early
        - why exit early?
        - because we wanted the smallest such ans 
        - and we greedily started from smallest such ans
    
    time = o(nlogn) + o(log$maxDiff * n)
    space = o(1)

*/
// func minimizeMax(nums []int, p int) int {
//     sort.Ints(nums)
//     start := 0
//     end := nums[len(nums)-1]
//     for i := start; i <= end; i++ {
//         atMax := i
//         count := 0
//         j := 0
//         for j < len(nums)-1 {
//             curr := nums[j]
//             next := nums[j+1]
//             diff := next-curr
//             if diff <= atMax {
//                 count++
//                 j += 2
//             } else {
//                 j++
//             }
//         }
        
//         if count >= p {
//             return i
//         }
//     }
//     return -1
// }