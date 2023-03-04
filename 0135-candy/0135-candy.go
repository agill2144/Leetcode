/*
    approach: greedy
    - First of all every child gets 1 candy ( regardless of their ratings compared with their neighboring childs )
    - Then for each ith child we will check whether it has more ratings than its neighboring child ; i-1 and i+1
        - HOWEVER, we have not calculated the final candies for the right child if we are looping from left to right
        - THEREFORE we cannot compare an ith child with its right child
    - So we will do this in 2 passes;
        - Left to right loop, compare ith rating with i-1th rating
        - if the i-1th rating is more, we KNOW FOR SURE, THAT THIS iTH CHILD GETS MORE THAN PREVIOUS CHILD
        - THEREFORE CANDIES[i] = previous child candies + 1
        - Now we can do the 2nd pass, in which we will compare ith child ratings with its right neighbor
        - if ratings[i] is higher than the right neighbor, than this ith child SHOULD GET MORE CANDIES COMPARED TO ITS RIGHT NEIGHBOR
        - However if candies[i] is already more than its right neighbor, than we are fine, no need to give more
        - otherwise, candies[i] = candies[i+1]+1 ( 1 more than its right neihbor )
    
    time : o(2n)
    space: o(n)
*/
func candy(ratings []int) int {
    candies := make([]int, len(ratings))
    candies[0] = 1
    for i := 1; i < len(candies); i++ {
        candies[i] = 1
        if ratings[i] > ratings[i-1] {
            candies[i] = candies[i-1]+1
        }
    }
    
    total := candies[len(candies)-1]
    for i := len(candies)-2; i>=0 ;i-- {
        if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
            candies[i] = candies[i+1]+1
        }
        total += candies[i]
    }
    return total
}