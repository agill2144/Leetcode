type RandomizedSet struct {
    idx map[int]int
    nums []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{idx:map[int]int{}, nums:[]int{}}    
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.idx[val]
    if ok {return false}
    this.nums = append(this.nums, val)
    n := len(this.nums)
    this.idx[val] = n-1    
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    n := len(this.nums)
    currPos, ok := this.idx[val]
    if !ok {return false}
    lastPos := n-1
    lastVal := this.nums[lastPos]
    this.nums[currPos], this.nums[lastPos] = this.nums[lastPos], this.nums[currPos]
    this.idx[lastVal] = currPos
    delete(this.idx, val)
    this.nums = this.nums[:n-1]
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */