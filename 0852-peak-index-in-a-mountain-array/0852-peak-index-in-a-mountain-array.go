func peakIndexInMountainArray(arr []int) int {
    left := 0
    right := len(arr)-1
    for left <= right {
        mid := left+(right-left)/2
        if (mid==0||arr[mid] > arr[mid-1]) && (mid==len(arr)-1||arr[mid] > arr[mid+1]) {
            return mid
        } else if (mid==0||arr[mid] > arr[mid-1]) {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}