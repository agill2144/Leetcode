/*
    subarray sum whose is divisible by k ( i.e $sum % k == 0 )
    we can check from start(0th idx) -> ith element by maintaining a runningSum
    and then check runningSum %k == 0
    
    what about other nested subarrays that do not start from idx 0?
    if their remainder is not 0, we need to save their remainders
        - this means at some ith position we had some sum that left $r as remainder
    if we later run into the same remainder again, it means,
    that some other number was added to runningSum and this other number IS PERFECTLY DIVISIBLE BY K
    eg;
    k = 5
    23 + 1 = 24 ; 24 % 5 = 4
    23 + 1 + 7 = 31; 31 % 5 = 1
    23 + 1 + 7 + 2 = 33; 33 % 5 = 3
    23 + 1 + 7 + 2 + 2 = 35; 35 % 5 = 5
    ------ no same remainders yet ----------------
    23 + 1 + 7 + 2 + 2 + 4 = 39; 39%5 = 4 ( same remainder! )
         4               4
            sum after prev
            4 is exactly 
            divisible by k
            7+2+2+4 = 15
    23 +     15 ;
    15 % 5 = 0; and we are left with 23 % 5 = 4
    
    Therefore the only time we see the same remainder again is if
    we added some number to runningSum that was perfectly divisible by k
    
    Another edge is
        in golang, C++, when we do -1 % 2 , we get the result as -1 , 
        whereas in other langs like python, the remainder is 1 ( the correct ans )
        so the remedy here is to just add the divisor(k) whenever the remainder is -ve
    
*/
func subarraysDivByK(nums []int, k int) int {
    count := 0
    freq := map[int]int{0:1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        
        /*
            I still dont have a full grasp on this theory; 
            
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