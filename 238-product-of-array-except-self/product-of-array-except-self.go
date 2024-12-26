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

    approach: prefix / suffix array
    - instead of allocating both prefix and suffix array
    - once we have prefix, we can re-use prefix array as suffix array
    - now each ith element in the prefix array is the prod of i except itself
    - we also wont need a separate output array
    - HOWEVER! we cannot store suffix product right away in prefix array
    - because each prefix[i] would now have the prod of prefix[i-1] * prefix[i+1]
    - and now when i moves back 1 more position while processing suffix, the prefix[i+1] is NOT the suffix prod
        - it is not the right running product
        - prefix[i+1] is prod of i+1 except itself
        - basically a pollution
    - hence we need to split and maintain a suffix prod separately in a running suffix prod array
        - call it "rrp" = right running prod    
    tc = o(2n) = o(n)
    sc = o(n)
*/
func productExceptSelf(nums []int) []int {
    n := len(nums)
    prefix := make([]int, n)
    prefix[0] = 1
    for i := 1; i < n; i++ {
        prefix[i] = nums[i-1]*prefix[i-1]
    }
    rrp := 1
    for i := n-2; i >= 0; i-- {
        rrp *= nums[i+1]
        prefix[i] *= rrp
    }
    return prefix
}