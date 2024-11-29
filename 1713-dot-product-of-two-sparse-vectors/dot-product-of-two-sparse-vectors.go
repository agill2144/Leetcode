/*
    approach: only store non-zero values in hashmap
    - prod of a number x with 0 will always be 0
    - therefore, there is no point in storing 0s
    - instead lets store all non-zero values and their idx positions
    - to make searching of idx position in another instance of this class
        - if we stored them in a map
        - we can easily lookup in o(1) time
    - therefore store all non-zero elements in an idx map
    tc = o(v1+v2) - when we have a not-so-sparse array
    sc = o(v1+v2)
*/
type SparseVector struct {
    idx map[int]int // {idx: val}
}

func Constructor(nums []int) SparseVector {
    idx := map[int]int{}
    for i := 0; i < len(nums); i++ {
        idx[i] = nums[i]        
    }
    return SparseVector{idx}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    sum := 0
    for key, val1 := range this.idx {
        val2, ok := vec.idx[key]
        if ok {
            sum += (val1*val2)
        }
    }
    return sum
}


/*
    approach: brute force
    - save the entire array
    - and do the product and save the sum
    tc = o(v1+v2)
    sc = o(v1+v2)
*/
// type SparseVector struct {
//     items []int
// }

// func Constructor(nums []int) SparseVector {
//     return SparseVector{nums}    
// }

// // Return the dotProduct of two sparse vectors
// func (this *SparseVector) dotProduct(vec SparseVector) int {
//     sum := 0
//     for i := 0; i < len(this.items); i++ {
//         sum += (this.items[i] * vec.items[i])
//     }
//     return sum
// }

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */