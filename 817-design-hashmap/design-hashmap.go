var (
    inner = 1001
    outter = 1001
)

type MyHashMap struct {
    data [][]*int    
}


func Constructor() MyHashMap {
    return MyHashMap{
        data: make([][]*int, outter),
    }    
}

func (this *MyHashMap) hash1(key int)int {return key%outter}
func (this *MyHashMap) hash2(key int)int {return key/inner}

func (this *MyHashMap) Put(key int, value int)  {
    idx1 := this.hash1(key)
    idx2 := this.hash2(key)
    if this.data[idx1] == nil {this.data[idx1] = make([]*int, inner)}
    this.data[idx1][idx2] = &value
}


func (this *MyHashMap) Get(key int) int {
    idx1 := this.hash1(key)
    idx2 := this.hash2(key)
    if this.data[idx1] == nil || this.data[idx1][idx2] == nil {return -1}
    return *this.data[idx1][idx2]
}


func (this *MyHashMap) Remove(key int)  {
    idx1 := this.hash1(key)
    idx2 := this.hash2(key)
    if this.data[idx1] == nil {return}
    this.data[idx1][idx2] = nil
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */