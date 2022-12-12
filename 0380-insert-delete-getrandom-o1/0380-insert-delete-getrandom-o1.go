type RandomizedSet struct {
    items []int
    valueToIdx map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        items: []int{},
        valueToIdx: map[int]int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    _, ok := this.valueToIdx[val]
    if ok {
        return false
    }
    this.items = append(this.items, val)
    this.valueToIdx[val] = len(this.items)-1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if idx, ok := this.valueToIdx[val]; ok {
        
        lastIdx := len(this.items)-1
        lastVal := this.items[lastIdx]
        
        // swap
        this.items[idx], this.items[lastIdx] = this.items[lastIdx], this.items[idx]
        
        // update lastVal in map to point to idx since its swapped now
        this.valueToIdx[lastVal] = idx
        
        // drop the last item from items
        this.items = this.items[:lastIdx]
        
        // drop this item from our map
        delete(this.valueToIdx, val)
        
        return true
    }
    return false
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