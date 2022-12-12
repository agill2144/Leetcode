/*
    approach: groupping diagonal cells together

    - kind of like groupping nodes in vertical order traversal of a binary tree
        - where we use the width of the tree to group nodes based on a width value
        - the width value would be reduced by 1 if we go left and plus 1 if we go right
        - and we did this in preorder fashion from left to right and used width as our groupping key in a map
        - then we simply looped over the map from min width to max width ( in that range )
    - likewise, we will create a group for each diagonal traverse
    - group by what tho?
        - group by row+col idx for each cell
        - so the group key in our map would be row+col idx
    - Notice in this matrix each col len could be diff, so its not fully filled out
        - this is called jaggered matrix
        - does that mean, as soon as we see a jaggered matrix like this, think grouping nodes if we have to do some sort of traversal? 
        - maybe?
    - anyways, if we started creating a map which would act as a tool to group by row+col key,
    - and if we did a normal matrix traversal from top to bottom left to right,
    - we will end up with incorrect ans for each group
    - therefore we will start from the last row and 0th col and work up instead of the normal start from 0th row and work down
        - this is to maintain the sort of each diagonal group
    - finally - we need to create a result array which contains all of the diagonal groups from top to bottom of the matrix
    - meaning in our map we need to start from key 0 to maxKeyVal that was ever stored
        - which means when we are building out our map, we need to keep track of the max width/sum 
    
    time: o(mn)
    space: o(mn) - worse case this is not a jaggered matrix and we ended up storing and grouping all elements in our map ( but only once )
*/
func findDiagonalOrder(nums [][]int) []int {
    diagGroup := map[int][]int{}
    m := len(nums)
    
    max := 0
    
    // time: o(mn)
    for i := m-1; i >= 0; i-- {
        for j := 0; j < len(nums[i]); j++ {
            rowColSum := i+j
            diagGroup[rowColSum] = append(diagGroup[rowColSum], nums[i][j])
            if rowColSum > max {max = rowColSum}
        }
    }
    out := []int{}
    
    // time: o(mn)
    for i := 0; i <= max; i++ {
        out = append(out, diagGroup[i]...)
    }
    
    // time: 2o(mn) = o(mn)
    // space: 
    
    return out
    
    
}