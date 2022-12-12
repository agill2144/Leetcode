/*
    sneaky sliding window problem...
    1. Get the total sum of all points first
    2. then create a window of totalLen-k
    3. [1,2,3,4,5,6,1]  ; k = 3
        - window would be [1,2,3,4]
        - make sure you have a windowSum
    4. Then remove windowSum from totalSum
    5. That gives you the answer for what if you picked last k cards
    6. Then slide the window ( of same size ) forward
        - [1,2,3,4,5,6,1]  ; k = 3
        - prev window =  [1,2,3,4]
        - new window = [2,3,4,5]
    7. Then subtract new window sum from totalSum
    8. That gives you the answer for what if you 1 picked card from front and remaining from back
    
    9. Then slide the window ( of same size ) forward
        - [1,2,3,4,5,6,1]  ; k = 3
        - prev window =  [2,3,4,5]
        - new window = [3,4,5,6]
    10. Then subtract new window sum from totalSum
    11. That gives you the answer for what if you picked 2 cards from front and 1 from back
    
     12. Then slide the window ( of same size ) forward
        - [1,2,3,4,5,6,1]  ; k = 3
        - prev window =  [2,3,4,5]
        - new window = [4,5,6,1]
    13. Then subtract new window sum from totalSum
    14. That gives you the answer for what if you picked all k cards from front
    15. Return the max out from each window evaluation
    
    time: o(n)
    space: o(1)
*/


func maxScore(cardPoints []int, k int) int {
    total := 0
    for i := 0; i < len(cardPoints); i++ {
        total += cardPoints[i]
    }
    if k == len(cardPoints) {
        return total
    }
    windowSize := len(cardPoints)-k
    left := 0
    sum := 0
    max := 0
    for right := 0; right < len(cardPoints); right++ {
        sum += cardPoints[right]
        if right-left+1 == windowSize {
            diff := total-sum
            if diff > max {
                max = diff
            }
            sum -= cardPoints[left]
            left++
        }
    }
    return max
}