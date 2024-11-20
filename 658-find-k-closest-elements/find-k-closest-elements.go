func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left := 0
    right := n-k-1
    for left <= right {
        mid := left + (right-left)/2
        // if mid+k >= len(arr){break}
        startDist := x-arr[mid]
        endDist := arr[mid+k]-x
        if endDist >= startDist {
            right--
        } else {
            left++
        }
    }
    return arr[left:left+k]
}
// func findClosestElements(arr []int, k int, x int) []int {
//     left := 0
//     right := len(arr)-1
//     for right-left+1 != k {
//         leftDist := abs(x-arr[left])
//         rightDist := abs(x-arr[right])
//         if rightDist >= leftDist {
//             right--
//         } else {
//             left++
//         }
//     }
//     return arr[left:left+k]
// }

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}