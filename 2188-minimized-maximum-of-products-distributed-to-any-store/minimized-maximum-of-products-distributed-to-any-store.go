/*
    "minimize the maximum number"
    - binary search on answer hint
      - is the answer in a known predictable range?
      - if yes, binary search on answer
    - greedy
      - if greedy does not work
      - DP hint    

    - each store can only have 1 product type
    - if we have product type x and it has 12 quantities
    - how many uniq stores would we need to distribute all 12 ?
        - 12/2 = 6 stores
        - evenly distribute meaning, 2 per store
        - this 2 is what we are looking for in this question
        - can we increase 2 to 3 ?
        - meaninging if we have 12 quantites, and each store can have 3
        - how many stores would be needed?
        - 12/3 = 4 stores
        - what if its not perfectly divisible?
        - 12/5 = 3 ( we round up )
    - say we ONLY have 6 stores
    - and we have 2 types of productes with diff quantities [11, 6]
    - we divy 11 based on 3 per store
        - how many stores needed ?
        - 11/3 = 4 (rounded up)
        - so out of 6 stores, 4 are used
        - we are left with 2 stores
    - we divy 6 based on 3 per stroe
        - how many stores needed?
        - 6/3 = 2 stores
    - no more stores left and no more products left
    - we used a hypothetical distribution quantity = 3
    - is there a know quantity per store we can assume ?
    - we can distribute 1 per store
    - or we can distrubute up to 11 per store ( max quantity )
    - we have a known range from 1 to 11
    - our answer lies in this range
    - ranges are sorted, therefore binary search
    - mid = potential ans ( distribution amount per store, our atMax value )
    - mid does not work if number of stores needed > n
    - it means our distribution amount is too small, therefore increase it, search right
    - otherwise it worked, save this as potential answer , and keep searching left to find a smaller ans

    time = o(log maxQuantity * n)
    space = o(1)

*/
func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := -1
    for i := 0; i < len(quantities); i++ {
        right= max(right, quantities[i])
    }
    ans := -1
    for left <= right {
        mid := left +(right-left)/2
        countPerStore := mid
        storesNeeded := 0
        for j := 0; j < len(quantities); j++ {
            storesNeeded += int(math.Ceil(float64(quantities[j])/float64(countPerStore)))
        }
        // when does mid not work?
        // when we need more stores than we have (n)
        if storesNeeded > n {
            // increase count per product type
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}