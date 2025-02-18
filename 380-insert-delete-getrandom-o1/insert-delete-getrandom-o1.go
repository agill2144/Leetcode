type RandomizedSet struct {
    idxs map[int]int
    nums []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{idxs:map[int]int{}, nums:[]int{}}    
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.idxs[val]
    if ok {return false}
    this.nums = append(this.nums, val)
    this.idxs[val] = len(this.nums)-1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    n := len(this.nums)
    idx, ok := this.idxs[val]
    if !ok {return false}
    lastIdx := n-1
    lastVal := this.nums[lastIdx]

    this.nums[idx], this.nums[lastIdx] = lastVal, val
    this.idxs[lastVal] = idx
    this.nums = this.nums[:n-1]
    delete(this.idxs, val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    n := len(this.nums)
    return this.nums[rand.Intn(n)]    
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */