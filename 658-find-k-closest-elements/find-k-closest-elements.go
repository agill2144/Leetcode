func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left := 0
    right := n-k
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

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}