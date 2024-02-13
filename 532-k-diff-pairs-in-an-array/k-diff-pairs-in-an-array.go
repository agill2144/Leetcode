/*
    if nums are sorted, we can also use two ptrs
    - putting two ptrs ( one on each side of the array ), will result into missing some results
    - do a dry on this ^ to see proof
    - hence using slow and fast ptrs that start from left side of the array and both move towards the right side of the array
*/
func findPairs(nums []int, k int) int {
    sort.Ints(nums)
    slow := 0
    fast := 1
    count := 0
    
    for fast < len(nums) {
        if slow == fast {
            fast++
            // why continue here ? to re-val the main loop break condition 
            continue
        }
        diff := nums[fast]-nums[slow]
        if diff == k {
            count++
            // since we have used both values, it makes sense to escape them both
            slow++
            fast++
            // but since the array could have dupes, keep escaping them until we do not have dupes
            for slow < fast && nums[slow] == nums[slow-1] {slow++}
            for fast < len(nums) && nums[fast] == nums[fast-1] {fast++}
        } else if diff > k {
            // we are using slow and fast ptrs
            // slow will be behind fast . i.e value at slow will be smaller than the value at fast
            // and when diff is too big, we have a sorted array
            // move slow ahead to reduce the bigger diff, otherwise moving fast forward will only make the diff even bigger
            slow++
        } else {
            fast++
        }
    }
    return count
}


/*
    searching for a pair ( two sum ), whats another possible solution?
    if nums are sorted, we can use binary search ( sorted ? think binary search right away )
    sort and binary search for the complement
    time = o(nlogn) for sort + o(nlogn) for binary search
    space = o(n) but this depends on the sorting algo that was used ( eg; if merge sort was used )
*/
// func findPairs(nums []int, k int) int {
//     sort.Ints(nums)
//     count := 0
//     for i := 0; i < len(nums); i++ {
//         // skip this number if its the same as previous
//         // if prev is same and had an ans, we already added it when i was at prev pos
//         // repeating the same element again as prev, will result into dupes
//         if i != 0 && nums[i] == nums[i-1] {continue}
        
//         // why do we not need the 1st complement search here?
//         // nums are sorted
//         // if at an ith number, we remove k, then c1 will be less than the ith value right ?        
//         // we have already searched from behind this ith value, thats why we are at this ith value
//         // if there was such pair before, it was already discovered before
//         // what about when i = 0
//         // when i = 0, nums[0]-k is , will be a number less than the num at idx 0 right
//         // is there a number before idx 0 ? no . there is not
//         // therefore in binary search we only need to search forward
//         // c1 := nums[i]-k
//         // idx1 := binarySearch(i+1, c1, nums)
//         // if idx1 != -1 {
//         //     count++
//         // }
        
//         c2 := nums[i]+k
//         idx2 := binarySearch(i+1, c2, nums)
//         if idx2 != -1 {
//             count++
//         }
//     }
//     return count
// }

// func binarySearch(left int, target int, nums []int) int {
//     right := len(nums)-1
//     for left <= right {
//         mid := left+(right-left)/2
//         if nums[mid] == target {
//             return mid
//         } else if target > nums[mid] {
//             left = mid+1
//         } else {
//             right = mid-1
//         }
//     }
//     return -1
// } 

/*
    approach: complement search with a hashset
    - if we have a number in hand; eg 5
    - and k is 2
    - what number would we have to remove from 5 to make it equal to k ?
    - that would be 3, and the way we get that is 5-2 = 3 ( nums[i]- k ),
    - another way to make 5 equal to k is,
    - 7-5 = k
    - how do we get a 7 ? , nums[i]+k
    - so we have two complements to search for ( nums[i]-k and nums[i]+k )
    - inorder to only store unique pairs, we will use a hashset and store each sorted pair
    
    time = o(n)
    space = o(n) for numsSet and uniquePairSet
*/

// func findPairs(nums []int, k int) int {
//     uniquePairSet := map[[2]int]struct{}{}
//     numsSet := map[int]struct{}{}
//     for i := 0; i < len(nums); i++ {
//         c1 := nums[i]-k
//         if _, ok := numsSet[c1]; ok {
//             tmp := [2]int{c1, nums[i]}
//             if nums[i] < c1 {tmp = [2]int{nums[i], c1}}
//             uniquePairSet[tmp] = struct{}{}
//         }
//         c2 := nums[i]+k
//         if _, ok := numsSet[c2]; ok {
//             tmp := [2]int{c2, nums[i]}
//             if nums[i] < c2 {tmp = [2]int{nums[i], c2}}
//             uniquePairSet[tmp] = struct{}{}
//         }
        
//         numsSet[nums[i]] = struct{}{}
//     }
//     return len(uniquePairSet)
// }

/*
    approach: brute force
    - create all possible pair
    - and see if their abs diff == k
    - if yes, save the sorted pair in a hashset to ensure uniqueness
        - because we are not counting duplicate pairs
    - finally, return the size of hashset
    
    time = o(n^2)
    space = o(n/2) ; o(n)
*/
// func findPairs(nums []int, k int) int {
//     set := map[[2]int]struct{}{}
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             if abs(nums[i]-nums[j]) == k {
//                 tmp := [2]int{nums[i], nums[j]}
//                 if nums[j] < nums[i] {
//                     tmp = [2]int{nums[j], nums[i]}
//                 }
//                 if _, ok := set[tmp]; !ok {
//                     set[tmp] = struct{}{}
//                 }
//             }
//         }
//     }
//     return len(set)
// }

func abs(x int) int {
    if x < 0 {return x * -1}
    return x
}