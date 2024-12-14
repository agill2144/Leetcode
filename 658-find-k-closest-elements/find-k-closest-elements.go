func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left, right := 0, n-k
    for left <= right {
        mid := left + (right-left)/2
        if mid+k >= n {break}
        startDist := x-arr[mid]
        endDist := arr[mid+k]-x
        if endDist >= startDist {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return arr[left:left+k]
}
// tc = o(n-k) + o(k)
// sc = o(1)
// func findClosestElements(arr []int, k int, x int) []int {
//     n := len(arr)
//     left, right := 0, n-1
//     for right-left+1 != k {
//         rightDist := abs(arr[right]-x)
//         leftDist := abs(arr[left]-x)
//         if leftDist <= rightDist {
//             right--
//         } else {
//             left++
//         }
//     }
//     return arr[left:left+k]
// }
