/*
    approach: binary search
    - we know that koko can eat any pile of bananas in an hour
    - we can pick any number from the piles array and use that
    - so if we pick the largest pile, we will find an answer
    - it may not the correct answer tho, 
    - because we want to minimize the number of bananas to be finished within h hours
    - so really the number of bananas to be eaten per hour lies between 1 banana to max number of numbers in a pile
    - which means if the max number of bananas in a pile is 11 then,
    - k can be anywhere between 1 to 11 ( inclusive )
    - so we can brute for from 1 to 11 or from 11 to 1 but because this is a range ( 1 to 11 )
    - we can apply binary search to find that k between 1 to 11
    - left will start from 1 ( because this is not idx or an array but rather a range from 1 to N)
        - left is not 0 because eating 0 bananas in an hour, will never resolve to an answer
        - therefore left starts from 1
    - right ptr will be the max number of bananas from the piles array ( 11 in this case if you are following the above example )
    - Then each mid point within our binary search is a potential k
        - now we find how long does it take to eat all the bananas if k = mid value
        - we will loop over the piles array and for each pile we will find how long it will take to finish this file IF rate of eating is mid
        - if piles[i] is 3 and mid is 6
        - then it will take 1 hour to finish this pile
        - english version: if I can eat 6 bananas in an hour, how long does it take to eat 3 banans ?
        - 3/6 = 1 ( because we round up -- because koko cannot go to next pile if the pile had less bananas than the mid value)
        - if piles[i] = 11 and mid is 6
        - then it will take 2 hours to finish this pile
        - english version: if I can eat 6 bananas in an hour, how long does it take to eat 11 banans ?
            - in the first hour we eat 6, and then in the next hour we eat the remaining 5 
        - 11/6 = 2 ( because we round up -- because koko cannot go to next pile if the pile had less bananas than the mid value)
        - This way we will go thru the whole piles array to find how long it will take to eat all the bananas
    - Once we have found how long it takes to finish the piles when mid is some N value
    - if the total hours taken for this mid is less than h 
        - save this as a potential k
        - and because we want to minimize rate of bananas per hour, we move the right pointer back to mid-1
        - because anything to the right of mid is going to give us a k larger than our current mid
    - otherwise , we need to increase the number of bananas to be eaten per hour
        - this means that the current mid does not work because the total time it took at $mid per hour was > allowed h
        - therefore move left ptr to mid+1
        
    
    time:
    - let n be th number of piles
    - o(n) to find the max number of bananas in a pile
        - let this be $max
    - binary search:
        - binary search happens from 1 to $max
        - so o(log$max)
        - however for each evaluation in a binary search, we run the the piles array again
        - therefore o(log$max) x o(n)
    - o(log$max) x o(n)
    
    space: o(1)
    
*/
func minEatingSpeed(piles []int, h int) int {
    left := 1
    right := math.MinInt64
    for i := 0; i < len(piles); i++ {
        if piles[i] > right {
            right = piles[i]
        }
    }
    
    k := math.MaxInt64
    for left <= right {
        mid := left + (right-left)/2 
        
        totalHours := 0
        for i := 0; i < len(piles); i++ {
            totalHours += int(math.Ceil(float64(piles[i])/float64(mid)))

        }
        if totalHours <= h {
            k = mid
            right = mid-1
        } else {
            left = mid+1
        }
        
    }
    return k
}