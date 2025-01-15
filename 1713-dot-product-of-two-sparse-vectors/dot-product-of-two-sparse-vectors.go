type SparseVector struct {
    items [][]int
}

func Constructor(nums []int) SparseVector {
    items := [][]int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {continue}
        items = append(items, []int{i, nums[i]})
    }
    return SparseVector{items}
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    if len(vec.items) < len(this.items) {
        return vec.dotProduct(*this)
    }
    // this items array will always be smaller
    //  loop over smaller array ( this array )
    //  use binary search over larger array ( vec array )
    total := 0
    for i := 0; i < len(this.items); i++ {
        idx := search(vec.items, this.items[i][0])
        if idx == -1 {continue}
        total += (this.items[i][1] * vec.items[idx][1])
    }
    return total
}

func search(items [][]int, targetIdx int) int {
    left := 0
    right := len(items)-1
    for left <= right {
        mid := left+(right-left)/2
        if items[mid][0] == targetIdx {
            return mid
        } else if targetIdx > items[mid][0] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */