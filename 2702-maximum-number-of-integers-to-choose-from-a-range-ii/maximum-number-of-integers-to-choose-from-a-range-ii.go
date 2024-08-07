// two pointers
func maxCount(banned []int, n int, maxSum int64) int {
    // sort the banned array
    sort.Ints(banned)

    // we want to to maximize the number of integers we picked
    // therefore start from 1 and go to n
    // starting with smaller number increases the amount of integers we choose
    // therefore.
    // we are allowed to pick numbers from 1 to n ( inclusive )
    // now using 2 ptrs, we can check if ith num == banned[ptr]
    // if yes, we can skip it
    bPtr := 0
    var rSum int64 = 0
    count := 0
    for i := 1; i <= n; i++ {
        if bPtr < len(banned) && i == banned[bPtr] {
            // cannot use this number
            bPtr++
            continue
        }

        // if after adding we exceed the maxSum, break and dont choose this number
        if int64(i) + rSum > maxSum {break}

        // can use this number
        rSum += int64(i)
        count++
    }
    return count
}