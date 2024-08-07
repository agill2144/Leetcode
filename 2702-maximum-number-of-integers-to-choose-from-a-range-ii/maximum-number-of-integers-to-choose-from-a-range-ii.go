// time = o(b) + o(logn)
// space = o(b)
// binary search on answers; we have to pick x numbers from range 1 to n
// we can pick 1, 2 ,3,4 or 5 numbers 
// this range is known and is a sorted range
// therefore mid = number of elemets we picked 
// how do we know if mid is valid?
// if the sum of all numbers from 1 to mid <= maxSum ( after remove banned numbers )
// if mid is valid, it means we can pick more numbers; therefore left = mid+1
// otherwise search left; right = mid-1
func maxCount(banned []int, n int, maxSum int64) int {    
    // we have to moved banned to a set
    // because there could be duplicate numbers
    // and a banned number that appears twice; should only be used once!
    set := map[int]struct{}{}
    for _, e := range banned {set[e] = struct{}{}}

    left := 1
    right := n
    maxCount := 0    
    for left <= right {
        // mid defines the number of numbers we have picked from 1 to mid
        mid := left + (right-left)/2
        
        // using arithmetic progression; get sum of all numbers until mid
        totalSum := int64(mid * (mid + 1) / 2)
        
        // remove banned numbers from the totalSum
        // if they are in part of totalSum ( ele <= mid )
        count := mid
        for b, _ := range set {
            if b <= mid {
                totalSum -= int64(b)
                count--
            }
        }
        
        // because mid defines we picked $mid numbers
        // and if this sum worked ( sum <= maxSum ), we can increase the number of numbers to pick
        // therefore increase mid ;left = mid+1
        if totalSum <= maxSum {
            maxCount = count
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return maxCount
}

// sort + two pointers
// time = o(blogb) + o(n)
// func maxCount(banned []int, n int, maxSum int64) int {
//     // sort the banned array
//     sort.Ints(banned)

//     // we want to to maximize the number of integers we picked
//     // therefore start from 1 and go to n
//     // starting with smaller number increases the amount of integers we choose
//     // therefore.
//     // we are allowed to pick numbers from 1 to n ( inclusive )
//     // now using 2 ptrs, we can check if ith num == banned[ptr]
//     // if yes, we can skip it
//     bPtr := 0
//     var rSum int64 = 0
//     count := 0
//     for i := 1; i <= n; i++ {
//         if bPtr < len(banned) && i == banned[bPtr] {
//             // cannot use this number
//             bPtr++
//             continue
//         }

//         // if after adding we exceed the maxSum, break and dont choose this number
//         if int64(i) + rSum > maxSum {break}

//         // can use this number
//         rSum += int64(i)
//         count++
//     }
//     return count
// }