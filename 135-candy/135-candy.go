/*
    Note: every child gets 1 candy regardless of the rating
    the child(s) with more ratings gets more than its neighboring child
    
approach: greedy
- The child with more ratings than its neighboring child gets 1 more candy if this childs rating is higher than its neighboring child
    - but which neighboring child?
    - immediate left and right neighbors
    - So thats how we will be greedy
    - We will do a initial left pass where each child is compared with its immediate left neighbor
    - if current child rating > immediate left child rating
    - then the current child gets 1 more candy than of left child
        - candies[i] = candies[left]+1
    - Then we will do a right pass, and repeat, however with a catch
    - if a ith child when doing the right pass, ( we will be comparing ith child rating with its immediate right child i+1)
    - if the ith rating > i+1 child's rating
    - which means that the current ith child should get more candy than of its immediate right neighbor
        - ONLY IF ith child candy count is <= i+1 childs candy count.
        - there is no need to give this ith child more candy just because its rating is higher - we also have to check if it already has more candies than i+1 child
    - Note that in left or right pass, we wont have any left or right child to compare with when starting left at idx 0 and starting right at last element
    - therefore we will start left from idx 1 and right from lastElement-1
    - Finally, we will sum all the numbers in our candies array

time: o(n) -- o(2n) but 2 is constant
space: o(n)
*/
func candy(ratings []int) int {
    candies := make([]int, len(ratings))
    for idx, _ := range candies {
        candies[idx] = 1
    }
    
    for i := 1; i < len(ratings); i++ {
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1]+1
        }
    }
    
    total := candies[len(candies)-1]
    for i := len(candies)-2; i>=0; i-- {
        if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
            candies[i] = candies[i+1]+1
        }
        total += candies[i]
    }
    return total
       
}