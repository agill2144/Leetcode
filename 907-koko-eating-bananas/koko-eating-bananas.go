/*
    approach: binary search
    - we know that koko can eat any pile of bananas in an hour
    - we can pick any number from the piles array and use that
    - so if we pick the largest pile, we will find an answer
    - it may not the correct answer tho, 
    - because we want to minimize the number of bananas to be finished within h hours
    - so really the number of bananas to be eaten per hour lies between 1 banana to max number of bananas in a pile
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
        - and because we want to minimize speed, we move the right pointer back to mid-1
        - because anything to the right of mid is going to give us a k larger than our current mid
    - otherwise , we need to increase the speed
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
    right := 0
    for i := 0; i < len(piles); i++ {
        right += piles[i]
    }
    ans := -1
    for left <= right {
        // if atMax, eating speed is $mid
        mid := left + (right-left)/2

        // how many hours does it take to finish at $mid eating speed
        hoursTaken := 0
        for j := 0; j < len(piles); j++ {
            hoursTaken += int(math.Ceil(float64(piles[j])/float64(mid)))
        }
        
        // did it work?
        // it works when we were able to finish all bananes within h hours
        // i.e if hours taken was <= h
        if hoursTaken <= h {
            ans = mid
            // however, we want the smallest such ans, therefore search left
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}

/*
    approach: brute force
    - we can be greedy and say that the eating speed is the total number of bananas
    - this way we know koko will finish all bananas in <= h hours.
    - however we want the smallest such eating speed ( binary search hint )
    - then we can start a range loop from totalBananas and walk back to 1 banana
        - in each iteration, we need to check how many hours does it take to finish the piles
        - if atMax the eating speed is $atMax / $i
        - if it works, i.e it took <= h hours, this is a potential ans, save it 
        - but keep walking back because we want smallest such eating speed
        - as soon as it does not work, we can break and return the last saved ans
    
    time = o(n) + o(totalSum * n)
    space = o(1)
*/
// func minEatingSpeed(piles []int, h int) int {
//     start := 1
//     end := 0

//     // o(n)
//     for i := 0; i < len(piles); i++ {
//         end += piles[i]
//     }

//     // o($sumOfAllPiles * n)
//     for i := start ; i <= end; i++ {
//         atMaxEatingSpeed := i
//         hoursTaken := 0
//         for j := 0; j < len(piles); j++ {
//             hoursTaken += int(math.Ceil(float64(piles[j])/float64(atMaxEatingSpeed)))
//         }

//         if hoursTaken <= h {
//             return atMaxEatingSpeed
//         }
//     }
//     return -1
// }