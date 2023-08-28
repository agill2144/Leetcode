func subarraysDivByK(nums []int, k int) int {
    count := 0
    freq := map[int]int{0:1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := (sum % k + k) % k
        
        /*
            in golang C++, when we do -1 % 2 , we get the result as -1 , 
            whereas in other langs like python, the remainder is 1 ( the correct ans )
            so the remedy here is to just add the divisor(k) whenever the remainder is -ve
        */
        if rem < 0 {rem += k}
        val, ok := freq[rem]
        if ok {
            count += val
        }
        freq[rem]++
    }
    return count
}