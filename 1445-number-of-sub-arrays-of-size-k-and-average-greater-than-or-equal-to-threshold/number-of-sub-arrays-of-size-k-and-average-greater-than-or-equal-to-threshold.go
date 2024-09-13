func numOfSubarrays(arr []int, k int, threshold int) int {
    count := 0
    rSum := 0
    left := 0
    for i := 0; i < len(arr); i++ {
        rSum += arr[i]
        if i-left+1 == k {
            if rSum/k >= threshold {
                count++
            }
            rSum -= arr[left]
            left++
        }
    }
    return count
}