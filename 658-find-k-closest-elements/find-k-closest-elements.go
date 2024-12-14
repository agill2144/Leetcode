func findClosestElements(arr []int, k int, x int) []int {
    n := len(arr)
    left, right := 0, n-1
    for right-left+1 != k {
        rightDist := abs(arr[right]-x)
        leftDist := abs(arr[left]-x)
        if leftDist <= rightDist {
            right--
        } else {
            left++
        }
    }
    return arr[left:left+k]
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}