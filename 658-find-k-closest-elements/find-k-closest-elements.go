func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left := 0
    right := n-k
    for left <= right {
        mid := left+(right-left)/2
        if mid + k >= len(arr) {break}
        leftDist := x-arr[mid]
        rightDist := arr[mid+k]-x
        if leftDist <= rightDist {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return arr[left:left+k]
}
// sorted;
// 1. two pointers
// tc = o(n-k) + o(k)
// sc = o(1)
// func findClosestElements(arr []int, k int, x int) []int {
//     n := len(arr)
//     left := 0
//     right := n-1
//     for right-left+1 != k {
//         leftDist := abs(arr[left]-x)
//         rightDist := abs(arr[right]-x)
//         if leftDist <= rightDist {
//             right--
//         } else {
//             left++
//         }
//     }
//     return arr[left:right+1]
// }

// func abs(x int) int {
//     if x < 0 {return x*-1}
//     return x
// }