func numOfSubarrays(arr []int) int {
    n := len(arr)
    rSum := 0
    even := 1
    odd := 0
    total := 0
    mod := 1000000007
    for i := 0; i < n; i++ {
        rSum += arr[i]
        if rSum % 2 == 0 {
            total += odd
            even++
        } else {
            total += even
            odd++
        }
        total %= mod
    }
    return total
}