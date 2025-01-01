/*
    approach: brute force
    - traverse the matrix from r1,c1 to r2,c2
    - and keep track of the sum
    
    q = len(queries)
    tc = 1(q*mn)
    sc = o(mn) - we stored the entire matrix 

    approach: prefix sum 1D array
    - if we had 1d array
    - we could have a prefix sum of all elements
    - then to calculate a subarray sum, given we have idxs of the subarray
    - the res will be = sum upto rightIdx - sum upto leftIdx-1
    - nums = [5,6,3]
    - sum = [5,11,14]
    - idx = 1,2 ; should be 9
    - sum up to rightIdx (2) = 14
    - sum up to leftIdx-1 (1-1 = 0) = 5
    - 14-5 = 9, as expected
    - then we can pretend each row is its own 1D array
    - and keep a prefix sum for each row within the same matrix
    - now we can start from row1 and col1 and get its prefix sum 
    - until and inlcuding row2 and col2
    - we would get prefix sum of a subarray in each row and keep adding to our total
    - at the end would have sum of the square region
    
    tc = o(m); worst case we are give row=0 until row=m-1
    - then we are looping over each row getting preifx sum for each row
    - however prefix sum in each row is done in constant time
    sc = o(mn) - we stored the entire matrix 

    approach: prefixSum on matrix
    - fullSum of a cell at x = sum from 0,0 to that cell x
    - to get fullSum of a cell, its the top sum + leftSum + currVal
    - however we will end up adding an extra cell , that is row-1 and col-1
    - therefore we will remove it
    - hence fullSum of cell at x = leftSum + topSum + curr.Val - overlappingCellSum
        - leftSum and rightSum, both had $overlappingCellSum
        - so we added it twice, therefore the "-overlappingCellSum"
    - give r1,c1 and r2,c2 is the matrix we need the sum for
    - if we had fullSum upto r2,c2 (from 0,0 to r2,c2)
    - then there will be 2 regions we need to remove sum from
    - we need to remove fullSum of 2 regions, top-region and leftRegion
    - topRegion fullSum is at = r1-1,c2
    - leftRegion fullSum is at = r2,c1-1
    - again, we have overlapping sum we removed twice
    - topRegion and leftRegion removed some overlapping fullSum twice
    - this overlapping fullSum region is at r1-1,c1-1
    - becuase we dont want to remove it twice, we will add it back once
    - hence sum of r1,c1 to r2,c2 is
        - fullSum(r2,c2)-topRegion-leftRegion+overlappingRegion
        - topRegion fullSum = r1-1,c2
        - leftRegion fullSum = r2,c1-1
        - overlappingRegion fullSum = r1-1,c1-1

    tc = o(q) for SumRegion
      = o(mn) for Constructor ( but this only happens once! )
    sc = o(mn) - we stored the entire matrix 
*/

type NumMatrix struct {
    matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m := len(matrix)
    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if j-1 >= 0 {matrix[i][j] += matrix[i][j-1]}
            if i-1 >= 0 {matrix[i][j] += matrix[i-1][j]}
            // remove the overlapping sum that was added twice
            // once it was added from left cell and other time it was added from top cell
            if i-1 >= 0 && j-1 >= 0 {matrix[i][j] -= matrix[i-1][j-1]}
        }
    }
    return NumMatrix{matrix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    fullSum := this.matrix[row2][col2]
    // remove top row
    if row1-1 >= 0 {fullSum -= this.matrix[row1-1][col2]}
    // remove left col
    if col1-1 >= 0 {fullSum -= this.matrix[row2][col1-1]}
    // add the overlapping sum that was removed twice
    if row1-1 >= 0 && col1-1 >= 0 {fullSum += this.matrix[row1-1][col1-1]}
    return fullSum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */