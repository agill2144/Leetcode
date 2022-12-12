// worst problem desc ever...
// brute force
// doing this question brute force first makes me understand why binary search helps...
// time: o(c) where c is the number of citations
// func hIndex(citations []int) int { 
//     n := len(citations)
//     for i := 0; i < len(citations); i++ {
//         diff := n-i
//         if diff <= citations[i] {
//             return diff
//         }
//     }
//     return 0
// }


/*
    approach: binary search
    - basically search for the same EXACT thing we searched for in brute force
    - if n-idx <= citations[idx]
        - in brute force ( linear traverse, we returned right away because its the first element from left side )
        - but in binary search, we will save it as it may NOT be the first element and so we will save it and search left
            HOWEVER, if diff is the value at that idx, just return, we wont find a better ans
    - otherwise
        - search right
    
    - Also this sounds like find either idx whose value == n-idx || if closest to n-idx
    
time: o(logn)
space: o(1)

*/
func hIndex(citations []int) int {
    left := 0
    hIdx := 0
    n := len(citations)
    right := len(citations)-1
    for left <= right {
        mid := left + (right-left)/2
        if n-mid <= citations[mid] {
            // if diff is the value at that idx, just return, we wont find a better ans
            if n-mid == citations[mid] {
                return n-mid
            }
            hIdx = n-mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return hIdx
}