type RandomizedSet struct {
    items []int
    idx map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{items: []int{}, idx: map[int]int{}}
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.idx[val]
    if ok {return false}
    this.items = append(this.items,val)
    this.idx[val] = len(this.items)-1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    currIdx, ok := this.idx[val]
    if !ok {return false}
    lastIdx, lastVal := len(this.items)-1, this.items[len(this.items)-1]    
    
    // swap
    this.items[currIdx], this.items[lastIdx] = this.items[lastIdx], this.items[currIdx]
    // update idx of last val to the swapped idx
    this.idx[lastVal] = currIdx
    // delete this value from idx map
    delete(this.idx, val)
    this.items =this.items[:len(this.items)-1]
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.items[rand.Intn(len(this.items))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */