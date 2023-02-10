/*
    We need to find the common elements between nums1 and nums2
    
    approach: freqMap
    - take smaller array and turn it into freqMap
    - Then loop over the larger array, and check if the each element exists in freqMap
    - if it does, then add this element to our result array 
    - decrement this element's frequency in freqMap by 1
    - if the frequency of this element reaches 0, then delete this element from map so we do not accidentally pick it again if this element exists more than N times in larger array compared to smaller array
    - Finally return the result array which will only have the common elements 
    
    m = larger array
    n = smaller array
    
    time: o(n) + o(m)
    space: o(n)
*/

// func intersect(nums1 []int, nums2 []int) []int {
    
//     // to make sure nums1 is always the larger array
//     if len(nums1) < len(nums2) {
//         return intersect(nums2, nums1)
//     }
    
    
//     freqMap := map[int]int{}
//     for _, num := range nums2 {
//         freqMap[num]++
//     }
    
//     result := []int{}
//     for i := 0; i < len(nums1); i++ {
//         val, ok := freqMap[nums1[i]]
//         if ok && val != 0 { // no need to delete, we can just check its frequency
//             result = append(result, nums1[i])
//             freqMap[nums1[i]]--
//         }
//     }
//     return result
// }


/*
    We need to find the common elements between nums1 and nums2.
    We can always sort both nums and deploy two pointers
    
    approach: 
    - sort(nums1), sort(nums2)
    - n1 and n2 ptrs starting at nums1 and nums2
    - if n1 == n2 then add this value to result and increment both ptrs by 1
    - if n1 > n2 - we need to move away from smaller number therefore increment n2
        - why?
        - because for n1 to become == n2, when n1 is larger than n2,
        - moving n1 will only make n1 larger than n2 because arrays are sorted
        - therefore moving n2 makes more sense as there is a possibility that if n2 is smaller than n1, then moving n2 may result into n2 == n1.
    - otherwise increment n1
    - what about out of bound check? when does our loop stop?
    - What if we do while n1 < len(nums1) && n2 < len(nums2)
    - whichever ptr goes out of bounds first, will stop our loop and help us avoid the out of bounds issue.
    
    m = larger array
    n = smaller array
    
    excluding sort time from below because I am implying the array will be sorted
    if not then add o(mlogm) + o(nlogn) to the below time 
    time: o(m) + o(n)
    space: o(1)
*/

// func intersect(nums1 []int, nums2 []int) []int {
//     // make sure nums1 is larger 
//     if len(nums1) < len(nums2) {
//         return intersect(nums2, nums1)
//     }
    
//     // sort ( implied )
//     sort.Ints(nums1)
//     sort.Ints(nums2)
    
    
//     n1 := 0
//     n2 := 0
//     result := []int{}
//     for n1 < len(nums1) && n2 < len(nums2) {
//         n1Val := nums1[n1]
//         n2Val := nums2[n2]
//         if n1Val == n2Val {
//             // we are looking for common elements
//             result = append(result, n1Val)
//             n1++
//             n2++
//         } else if n1Val > n2Val {
//             n2++
//         } else {
//             n1++
//         }
//     }
//     return result
// }



/*
    We need to find the common elements between nums1 and nums2.
    If we can imply that nums1 and nums2 are sorted, then
    another potential solution is to explore binary search
    
    approach: 
    - sort(nums1), sort(nums2)
    - binary search on the bigger array ( nums1 )
    - for every element in nums2
    - we will search for THE LEFTEST IDX of this number in nums1 using binary search
    - leftest ($x == target , not closest )
    - If we find it, then we will shrink our binary search window and exclude this idx we just added to our result array ( we cannot include the same number again in our next binary search )


    m = larger array
    n = smaller array
    
    excluding sort time from below because I am implying the array will be sorted
    if not then add o(mlogm) + o(nlogn) to the below time 
    time: o(nlogm)
    space: o(1)
*/
func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) < len(nums2) {
        return intersect(nums2, nums1)
    }
    sort.Ints(nums1)
    sort.Ints(nums2)
    result := []int{}
    leftIdx := 0
    for i := 0; i < len(nums2); i++ {
        idxFoundAt := binarySearch(leftIdx, nums2[i], nums1)
        if idxFoundAt != -1 {
            result = append(result, nums2[i])
            leftIdx = idxFoundAt+1
        }
    }
    return result
}

func binarySearch(left, target int, nums []int) int {
    idx := -1
    right  := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            idx = mid
            right = mid-1
        } else if nums[mid] < target {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx
}