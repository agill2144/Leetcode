/*
    approach:
    - its possible input nums is empty
        - then the whole range is missing
        - return lower,upper
    - its possible that nums[0] starts after lower
        - meaning lower is smaller than starting number
        - meaning the range from lower to startNumber-1 is missing
        - therefore check and add this range first in output array
    - now we can start looping from idx 1, to check if there is a range missing 
        between curr number and prev number
    - if prev number + 1 == curr number, no range missing, continue
    - otherwise, some range is missing
        - the missing range start value is prev+1 and end value is curr-1
        - now check if this range is within our allowed lower and upper range
        - i.e prev+1 >= lower and curr+1 <= upper
        - if yes, add this missing range to our output array
    - another potential range that could be missing is after the loop is over
    - just like a range could be missing from the beginning
    - a range could also be missing at the end
    - i.e nums[n-1] < upper
    - then missing end range is [nums[n-1]+1, upper]
*/
func findMissingRanges(nums []int, lower int, upper int) [][]int {
    if len(nums) == 0 {return [][]int{{lower, upper}}}
    out := [][]int{}
    if nums[0] > lower {
        out = append(out, []int{lower, nums[0]-1})        
    }
    for i := 1; i < len(nums); i++ {
        curr := nums[i]
        prev := nums[i-1]
        if prev + 1 == curr {continue}
        if curr > upper {break}
        currPrev := curr-1
        prevNext := prev+1
        if prevNext >= lower && currPrev <= upper {
            out = append(out, []int{prevNext, currPrev})
        }
    }
    if nums[len(nums)-1] < upper {
        out = append(out, []int{nums[len(nums)-1]+1, upper})
    }
    return out
}