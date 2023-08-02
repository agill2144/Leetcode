var arrSize int = 1000000
type MyHashMap struct {
    arr []*int
}


func Constructor() MyHashMap {
    return MyHashMap{make([]*int, arrSize)}
}

func (this *MyHashMap) hashIdx(key int) int {
    return key % arrSize
}


func (this *MyHashMap) Put(key int, value int)  {
    idx := this.hashIdx(key)
    this.arr[idx] = &value
}


func (this *MyHashMap) Get(key int) int {
    idx := this.hashIdx(key)
    if this.arr[idx] == nil {return -1}
    return *this.arr[idx]
}


func (this *MyHashMap) Remove(key int)  {
    idx := this.hashIdx(key)
    this.arr[idx] = nil    
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */