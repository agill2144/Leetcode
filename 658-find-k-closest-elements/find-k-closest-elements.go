// sorted;
// 1. two pointers
func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left := 0
    right := n-1
    for right-left+1 != k {
        leftDist := abs(arr[left]-x)
        rightDist := abs(arr[right]-x)
        if leftDist <= rightDist {
            right--
        } else {
            left++
        }
    }
    return arr[left:right+1]
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}