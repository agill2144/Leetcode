/*
    approach: running sum
    - We will use running sum / count principle
    - Since we need the longest subarray, we will use a map to store {$rs:$idxSeenAt}
    - running sum / count will be incremented by 1 if ith val == 1 
    - running sum / count will be decremented by 1 if ith val == 0
    - If we run into a running sum at an ith position that we have seen before ( lookup new rs in map )
    - Than that means from that idx from map lookup+1 till current ith idx - we have a balanced subarray
    - Calc this window size ( right-left+1 ) -- i-(idxFromMap+1)+1
    - Then save the above window size to max as needed.
    
    Time: o(n)
    Space: o(1)
*/
func findMaxLength(nums []int) int {
    count := 0
    max := 0
    countIdx := map[int]int{0:-1}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0{count--}
        if nums[i] == 1{count++}
        idx, seen := countIdx[count]
        idx++
        if !seen {
            countIdx[count] = i
        } else {
            // current subarray size
            cl := i-idx+1
            if cl > max {
                max = cl
            }
        }
    }
    return max
}