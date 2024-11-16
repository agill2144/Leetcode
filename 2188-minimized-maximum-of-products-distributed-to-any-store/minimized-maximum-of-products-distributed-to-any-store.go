/*
    - each store can only have 1 product type
    - if we have quantites = [15,10,10], and stores = 7
    - and we want to minimize the number of products per store
    - then lets start super small, at 1, give only 1 quantity of a product to a store
    - number of stores needed to distribute 15 quantities, if a store can only have 1 ?
        - 15 stores ( 15/1 = 15)
        - but we only have 7 stores, therefore our min is too small
        - lets try 2
    - number of stores needed to distribute 15 products if a store can only have 2 ?
        - 8 stores (ceil(15/2))
        - again, we only have 7 stores
        - lets try 5
    - number of stores needed to distribute 15 products if a store can only have 5 ?
        - 3 stores ( 15/5 = 3)
        - number of stores needed to distribute 10 products
        - 2 stores ( 10/5 = 2)
        - number of stores needed to distribute 10 products
        - 2 stores ( 10/5 = 2
        - total stores = 7
        - which is <= our n
        - therefore 5 is the smallest possible distribution we can have
        - if we try 6, it will surely work
        - 7,8,9,... anything beyond 5 will work
    - does 4 work?
        - 15/4 = 4
        - 10/4 = 3
        - 10/4 = 3
        - 10 stores needed, no 4 does not work
    - this is a range, 1,2,3,4,5.... till max quantity
    - sorted range
    - therefore we can use binary search
    - mid = potential ans ( distribution amount atMax )
    - mid does not work if number of stores needed > n
    - it means our distribution amount is too small, therefore increase it, search right
    - otherwise it worked, save this as potential answer , and keep searching left to find a smaller ans

    time = o(log maxQuantity * n)
    space = o(1)
*/
func minimizedMaximum(n int, quantities []int) int {
    left := 1
    right := -1
    for i := 0; i < len(quantities); i++ {right = max(right, quantities[i])}
    ans := 0
    for left <= right {
        mid := left + (right-left)/2
        count := 0
        for i := 0; i < len(quantities); i++ {
            count += int(math.Ceil(float64(quantities[i])/float64(mid)))
        }
        if count > n {
            left = mid+1
        } else {
            ans = mid
            right = mid-1
        }
    }
    return ans
}