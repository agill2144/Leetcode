/*
    approach: brute force
    - nested loop
    tc = o(n^2)
    sc = o(1)
    
    approach: prefix / suffix array
    - get a prefix prod for each ith position ( excluding that position )
    - get a suffix prod for each ith position ( excluding that position )
    - now for each ith position prodExceptItself is prefix[i] * suffix[i]
    tc = o(3n) = o(n)
    sc = o(2n) = o(n)
*/
func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefix := make([]int, n)
    prefix[0] = 1
    for i := 1; i < n; i++ {
        prefix[i] = nums[i-1]*prefix[i-1]
    }
    suffix := make([]int, n)
    suffix[n-1] = 1
    for i := n-2; i >= 0; i-- {
        suffix[i] = suffix[i+1] * nums[i+1]
    }

    out := make([]int, n)
    for i := 0; i < n; i++ {
        out[i] = prefix[i] * suffix[i]
    }
    return out
}