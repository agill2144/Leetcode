type RandomizedSet struct {
    items []int
    idx map[int]int 
}


func Constructor() RandomizedSet {
    return RandomizedSet{items: []int{}, idx: map[int]int{}}    
}


func (this *RandomizedSet) Insert(val int) bool {
    _, exists := this.idx[val]
    if exists {return false}
    this.items = append(this.items, val)
    this.idx[val] = len(this.items)-1    
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    currPos, exists := this.idx[val]
    if !exists {return false}
    lastPos := len(this.items)-1
    lastVal := this.items[lastPos]
    this.items[currPos], this.items[lastPos] = lastVal, val
    this.items = this.items[:len(this.items)-1]
    this.idx[lastVal] = currPos
    delete(this.idx, val)    
    return true
}


func (this *RandomizedSet) GetRandom() int {
    n := len(this.items)
    return this.items[rand.Intn(n)]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */