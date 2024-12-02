func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left := 0
    right := n-k
    for left <= right {
        mid := left + (right-left)/2
        startDist := x-arr[mid]
        if mid+k == n {break}
        endDist := arr[mid+k]-x
        if endDist >= startDist {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return arr[left:left+k]
}
/*
    approach: 2 ptrs
    - left ptr at 0
    - right ptr at n-1
    - calc dist of both
    - move away from ptr whose dist is farther away
    - what if dists are same?
        - move away from right
        - "|a - x| == |b - x| and a < b"
        - a = left; b = right
        - we want to keep a if a is smaller
        - this is a sorted array, therefore a will always be smaller
        - therefore we move from right ptr
    time = o(n-k) + k
    space = o(1)

*/
// func findClosestElements(arr []int, k int, x int) []int {
//     n := len(arr)
//     left := 0
//     right := n-1
//     // n-k
//     for right-left+1 != k {
//         leftDist := abs(arr[left]-x)
//         rightDist := abs(arr[right]-x)
//         if rightDist >= leftDist {
//             right--
//         } else {
//             left++
//         }
//     }
//     return arr[left:right+1]
// }

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}