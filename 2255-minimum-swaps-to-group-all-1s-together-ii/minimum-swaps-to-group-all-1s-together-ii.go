// pattern: let sliding window wrap around
// dont use the right(i) ptr to stop the loop
// if right(i) ptr reaches the end, wrap it around ( reset back to 0 )
// use the left ptr to stop the loop
// once left ptr reaches the end idx, it means we have used endIdx + [0:i] 
// i.e we have used to wrap our window back to the front

func minSwaps(nums []int) int {
    k := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            k++
        }
    }
    if k == 0 {return 0}

    left := 0
    ones := 0
    n := len(nums)
    swaps := math.MaxInt64
    i := 0
    for left < n {
        ii := i % n
        if nums[ii] == 1 {
            ones++
        }
        size := i-left+1
        if size == k {
            nonOnes := size - ones
            swaps = min(swaps, nonOnes)
            if nums[left] == 1 {
                ones--
            }
            left++
        }
        i++
    }
    return swaps
}