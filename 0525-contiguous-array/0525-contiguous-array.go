/*
    approach: brute force
    - form each subarray
    - check if number of zeros and ones are equal
    - we could maintain 2 different vars ( zeroCount and oneCount )
    - and once they are both equal, we can check the size of this subarray and
        save if its better than before
    - OR instead of counting in 2 vars, we can maintain 1 count
        - if ele == 1; count++
        - if ele == 0; count--
    - at any point count becomes 0, the subarray is a balanced subarray ( equal number of 0s and 1s )

    time = o(n^2)
    space = o(1)

    
    when we have nested iterations, consider the following techniques
    1. binary search
    2. two ptrs
    3. hashing
    4. running sum
    5. sliding window
    6. dp
    
    Why cant we use sliding window here ?
    - we dont have a definitive way that helps us decide when to shrink our window
    
    approach: running sum
    - the n^2 time is because of nested iteration
    - when is a subarray considered balanced ?
        - when count is 0
    - so if we have a runningCount called X up to a certain idx
    - what will make this subarray balanced ? , the only way is if we X was 0
    - how can make X 0 ?
        - subtract X from X
    - Therefore we need to search for the same X value before ( lets call this Y )
    - if we have seen a Y which is equal to X, it means;
        - up to this x idx, our count is X
        - up to that y idx, the count was y
        - X and Y are the same count
        - [. . . . . . . . . .]
               Y       X
        - if we remove those Y elements from this X array, our count will become 0
        - therefore balanced.
    - If we have, then we can check whether at this point , can there be a subarray thats balanced ?
    - How do we search previous running-counts 
        - we can use a map to maintain previous running-counts
        - because we want to return the size of the longest subarray, we will store running-count as {count:idx}
        - if we have already saved a running-count;
            - we wont override its idx
            - we will be greedy and keep the old one to increase chances of finding the longest one

    
    Why start our idxMap with {0:-1} ?
    - because we could have a case at the beginning where running count becomes 0
    - but we have not had added a 0 yet in our hashmap, AND if we reached a 0 without having any initial 0 in idx map, this is a balanced subarray right
    - therefore we need to start our idx map with {0:-1}
    Example: [1,1,1,0,0,0]
    { 1:0,2:1,3:2, 2:3, 1:4, 0:5 }
    - at last idx, our count will become 0, and this is balanced and the largest one
    - but we do not have a 0 saved in our hashmap
    - we will have an answer but that wont be the longest subarray ( rSum 1 is repeating )
    - however we miss the fact that the rSum ended with 0 and because we never had a 0 to compare with, we never assume that as a potential answer.
    - Therefore we need to start our hashmap with {0:-1}

    
    time: o(n)
    space: o(n) // unique counts at each idx; example [1,1,1,1,1,1,1]
*/
func findMaxLength(nums []int) int {
    countToIdx := map[int]int{0:-1}
    runningCount := 0
    maxLen := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0{
            runningCount--
        } else {
            runningCount++
        }
        idx, ok := countToIdx[runningCount]
        if ok {
            if i-(idx+1)+1 > maxLen {
                maxLen = i-(idx+1)+1
            }
        } else {
            countToIdx[runningCount] = i
        }
    }
    return maxLen
}


// brute force, form each subarray and keep track of count
// time = o(n^2)
// space = o(1)
// func findMaxLength(nums []int) int {
//     max := math.MinInt64
//     for i := 0; i < len(nums); i++ {
//         count := 0
//         for j := i; j < len(nums); j++ {
//             if nums[j] == 0 {
//                 count--
//             } else {
//                 count++
//             }
//             if count == 0 {
//                 if j-i+1 > max {
//                     max = j - i + 1
//                 }
//             }
//         }
//     }
//     if max == math.MinInt64 {max = 0}
//     return max
// }
