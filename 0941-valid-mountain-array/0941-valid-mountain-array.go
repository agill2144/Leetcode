/* approach: two pointers 
    time: o(n)
    space: o(1)
*/
func validMountainArray(arr []int) bool {
    n := len(arr)
    left := 0
    for left < len(arr)-1{
        if arr[left] >= arr[left+1] {
            break
        }
        left++
    }

    right := len(arr)-1
    for right >= 1 {
        if arr[right] >= arr[right-1] {
            break
        }
        right--
    }
    return (left != 0 && right != 0) && (left != n-1 && right != n-1) && left==right 

}

/* approach: binary search + linear scan
    - GIVEN there is ONLY 1 PEAK
    - use binary search to find the peak ( mid > mid-1 && mid > mid+1 )
    - Then linear scan from left to peak, all elements must be increasing order ( strictly increasing )
    - Then linear scan from right to peak, all elements must be increasing order ( strictly increasing )
    
    time: o(logn) + o(n)
    space: o(1)
*/
// func validMountainArray(arr []int) bool {
//     if len(arr) < 3 {return false}

//     left := 0
//     right := len(arr)-1
//     peakIdx := -1
//     for left <= right {
//         mid := left+(right-left)/2
//         if (mid==0||arr[mid] > arr[mid-1]) && (mid==len(arr)-1||arr[mid] > arr[mid+1]) {
//             peakIdx = mid
//             break
//         } else if (mid==0 ||arr[mid] > arr[mid-1]) {
//             left=mid+1
//         } else {
//             right = mid-1
//         }
//     }
//     // if we didnt find a peak, return false
//     if peakIdx == -1 || peakIdx == 0 || peakIdx == len(arr)-1 {return false}
    
//     // walk up from left to peak ( everything should be incresing )
//     for i := 1; i < peakIdx; i++ {
//         if arr[i] <= arr[i-1] {return false}
//     }
    
//     // walk down from peak ( everything should be decreasing )
//     for i := peakIdx + 1; i < len(arr); i++ {
//         if arr[i] >= arr[i-1] {return false}
//     }
    
    
//     return true
// }