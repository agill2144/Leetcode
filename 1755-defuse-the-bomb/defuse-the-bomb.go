func decrypt(nums []int, k int) []int {
    n := len(nums)
    out := make([]int, n)
    if k == 0 {return out}
    left := 0; right := 0
    sum := 0
    ptr := n-1
    c := 0
    if k < 0 { k*=-1; ptr = k}
    for c != n {
        sum += nums[right%n]
        if right-left+1 == k {
            out[ptr%n] = sum
            sum -= nums[left%n]
            left++
            ptr++
            c++
        }
        right++
    }
    return out
}