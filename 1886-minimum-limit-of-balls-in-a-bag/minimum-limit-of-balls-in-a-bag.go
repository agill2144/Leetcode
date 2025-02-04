/*
    NOTE:
    - 1 operation = creates 2 bags

    approach: binary search on answers
    - the answer lies within a range
    - the min penalty is 1
    - the max penalty is max element in nums
    - ranges are sorted, hence binary search on our answer space (1 to max element)
    - we are using binary search to evaluate whether x is a possible answer
    - x = max number of balls
    - x becomes a valid answer IF
    - we are able to take <= maxOperations to divide input into x number of balls per bag
    - but we want smallest such x
    - therefore save this x as a potential answer and search left
    
    How to evaluate whether x(penalty) works?
    - well it works if were able to divide balls no more than x per bag 
        without exceeding maxOperations
    - in 1 operation, we can divide a bag into 2 bags
    - so i have a bag of 2 balls, in 1 operation, i get 2 bags in return
    - x = atMax bags allowed PER bag
    - if we have 8 balls in a bag and our atMax allowed is 2 per bag
    - number of operations needed would be ceil(8/2) - 1
        - i.e: math.Ceil(nums[i]/x) - 1
    - x = atMax bags allowed per bag = max penalty = mid value
    - we want smallest such x
    - when does mid not work?
        - when operations needed to fulfill max balls per bag exceed allowed maxOps
        - meaning, our bag size is too small, we need to allow more balls
        - hence left = mid+1
    - when does mid work?
        - when operations needed to fulfill max balls per <= allowed maxOps
        - save this as potential answer
        - keep search left for smaller ans
        - hence right = mid-1
    
    tc = o(n) + o(log(maxEle) * n)
    sc = o(1)
*/
func minimumSize(nums []int, maxOps int) int {
    left := 1
    right := slices.Max(nums)
    if maxOps <= 0 {return right}
    res := -1
    for left <= right {
        mid := left + (right-left)/2
        atMax := mid
        ops := 0
        for i := 0; i < len(nums); i++ {
            ops += int(math.Ceil(float64(nums[i])/float64(atMax)))-1
            if ops > maxOps {break}
        }
        if ops <= maxOps {
            res = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return res
}