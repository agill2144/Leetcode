var arrSize int = 1001
var innerArrSize int = 1001
type MyHashMap struct {
    arr [][]*int
}


func Constructor() MyHashMap {
    return MyHashMap{make([][]*int, arrSize)}
}

func (this *MyHashMap) hashIdx(key int) int {
    return key % arrSize
}

func (this *MyHashMap) innerHashIdx(key int) int {
    return key / innerArrSize
}


func (this *MyHashMap) Put(key int, value int)  {
    idx := this.hashIdx(key)
    idx2 := this.innerHashIdx(key)
    if this.arr[idx] == nil {
        this.arr[idx] = make([]*int, innerArrSize)
    }
    this.arr[idx][idx2] = &value    
}


func (this *MyHashMap) Get(key int) int {
    idx := this.hashIdx(key)
    idx2 := this.innerHashIdx(key)
    if this.arr[idx] == nil || this.arr[idx][idx2] == nil {return -1}
    return *this.arr[idx][idx2]
}


func (this *MyHashMap) Remove(key int)  {
    idx := this.hashIdx(key)
    idx2 := this.innerHashIdx(key)
    if this.arr[idx] == nil {return}
    this.arr[idx][idx2] = nil    
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */