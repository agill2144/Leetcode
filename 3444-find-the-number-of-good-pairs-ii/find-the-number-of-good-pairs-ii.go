func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {

    maxN1 := math.MinInt64
    n1Freq := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        n1Freq[nums1[i]]++
        maxN1 = max(maxN1, nums1[i])
    }
    n2Freq := map[int]int{}
    for i := 0; i < len(nums2); i++ {
        n2Freq[nums2[i]*k]++
    }
    var count int64
    for ele, freq := range n2Freq {
        b := ele
        for b <= maxN1 {
            val, ok := n1Freq[b]
            if ok {
                count += (int64(val) * int64(freq))
            }
            b += ele
        }
    }
    return count

}

// func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
//     for i := 0; i < len(nums2); i++ {
//         nums2[i] *= k
//     }
//     sort.Ints(nums2)
//     sort.Ints(nums1)
//     var count int64
//     left := 0
//     for i := 0; i < len(nums2); i++ {
//         n2 := nums2[i]
//         left = findStartPos(nums1, left, n2)
//         if left >= 0 {
//             for j := left ; j < len(nums1); j++ {
//                 if nums1[j] % n2 == 0 {count++}
//             }
//         }
//     }
//     return count
// }

// target = 6
// nums = [1,2,3,4,5,6,7,8,9,10]
// we want to return the idx of 6
// nums = [1,2,3,4,5,8,9,10]
// we want to return the idx of 8
// so we want the leftmost on right side of target
// func findStartPos(nums []int, left int, target int) int {
//     right := len(nums)-1
//     ans := 0
//     for left <= right {
//         mid := left + (right-left)/2
//         if nums[mid] >= target {
//             ans = mid
//             right = mid-1
//         } else {
//             left = mid+1
//         }
//     }
//     return ans
// }