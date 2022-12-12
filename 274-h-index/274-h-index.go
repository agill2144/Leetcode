
// brute force - sort + binary search
// hIdx is n-idx <= valAtThatIdx
// if valAtThatIdx == n-idx == then ans is n-idx
// func hIndex(citations []int) int {
//     sort.Ints(citations)
//     left := 0
//     right := len(citations)-1
//     n := len(citations)
//     hIdx := 0
//     for left <= right {
//         mid := left + (right-left)/2
//         diff := n-mid
//         if diff <= citations[mid] {
//             if diff == citations[mid] {
//                 return diff
//             }
//             hIdx = diff
//             right = mid-1
//         } else {
//             left = mid+1
//         }
//     }
//     return hIdx
// }


// approach: partial sorting using bucket sort
// time: o(n)
// space: o(n)
func hIndex(citations []int) int {
    n := len(citations)
    buckets := make([]int, n+1)
    for i := 0; i < len(citations); i++ {
        if citations[i] >= n {
            buckets[n]++
        } else {
            buckets[citations[i]]++
        }
    }
    sum := 0
    for i := len(buckets)-1; i >= 0 ; i-- {
        sum += buckets[i]
        if sum >= i {
            return i
        }
    }
    return 0
}