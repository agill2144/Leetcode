/*
    approach: linear scan / two ptrs
    - maintain a ptr to keep track of "potential" missing number
    - another ptr in input arr
    - if 2 pointers value are not same;
        - it means we have ran into a missing number
        - increment a missing counter by 1
        - if our missing counter == k; we have seen exactly k missing elements
            - meaning the current value our ptr is at, is the missing number
        - otherwise, increment ptr by 1
    - otherwise, we have same value, meaning nothing is missing
        - move both ptrs by 1

    time = o(n)
    space = o(1)
*/
func findKthPositive(arr []int, k int) int {
    mc := 0
    ptr := 1
    i := 0
    for i < len(arr) {
        if ptr != arr[i] {
            mc++
            if mc == k {
                // dry run this on board to understand better
                // make k outside of arr, while arr is missing some numbers
                return ptr+(k-mc)
            }
            ptr++
        } else {
            ptr++
            i++
        }
    }
    // when nothing was missing in arr, our ptr value went ahead by 1, when we were at the last
    // therefore the -1 at the end
    // dry run on board to see better
    return ptr+(k-mc)-1
}