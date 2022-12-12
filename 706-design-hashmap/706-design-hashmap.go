/*
    Hashmaps use arrays under the hood to store items.
    How does it determine the idx to store a key val pair?
    - "Hash"maps use hashing functions to hash a given key that maps out to a idx in the underlying array
    - Then that idx is used.
    - This is true for any hashing ds
    - The same hashing function is used to Get key value from hashmap
    - This is how o(1) is achieved - Get(key -> hash -> idx), Put(key -> hash -> idx)
    
    
    Key properties for any hashing functions:
    - It should be consistent
    - Given a hash key - it should produce the same hash value over and over again and it should not change
    - Should avoid and or handle any collisions
    
    Lets talk collisions
    - Collision happens when a hashing function maps 2 different hash keys to same underlying idx
    - For example if our hash func is based using  : key % arraySize
    - Given our array size is 1k, then a key 0 and a key 1000 will produce the same idx
    - Therefore overriding the same idx value in the underlying array, therefore loss of data
    - How do handle collisions?
        - You can make sure your array size is big enough where no 2 keys collide using %
        - if you cannot use a bigger array, then the choices are to handle it with one of:
            1. Chaining ( array of LL - each idx will be its own LL )
                - Pros: No more space issues, and grows and shrinks dynamically
                - Cons: Time becomes o(n) for each Get, Put
            2. Double hashing ( array of array , i.e 2D array )
                - Intially allocate outter array of sqrt(hashMapSize) -- so if 10k is max, then allocate only 1k in outter array
                - When a Put(key) is called, and the hash func points to an idx whose nested array is not yet initialized, initialize it then 
                    - So, we will eventually in iterations allocate the 10k size but not right away - so somewhat space optimized.
                    - What will be the inner array size? outterArraySize+1
                - Now we will have a nested array in a single cell, what will be the idx in nested array where we will store the value of this key?
                - We will hash the key one more time, using  key / innerArraySize

*/

const (
    arraySize = 1000
    innerArraySize = arraySize + 1 
)

func hash(key int) int {return key % arraySize }
func hash2(key int) int {return key / arraySize }
func toPtrInt(n int) *int {return &n}

type MyHashMap struct {
    items [][]*int // using ptr to an int to distinguish between null or 0
}


func Constructor() MyHashMap {
    return MyHashMap{
        items: make([][]*int, arraySize),
    }
}

func (this *MyHashMap) Put(key int, value int)  {
    idx := hash(key)
    innerIdx := hash2(key)
    
    if this.items[idx] == nil {
        this.items[idx] = make([]*int, innerArraySize)
    }
    this.items[idx][innerIdx] = toPtrInt(value)
}


func (this *MyHashMap) Get(key int) int {
    idx := hash(key)
    innerIdx := hash2(key)
    if this.items[idx] == nil || this.items[idx][innerIdx] == nil {
        return -1
    }
    return *this.items[idx][innerIdx]
}


func (this *MyHashMap) Remove(key int)  {
    idx := hash(key)
    innerIdx := hash2(key)
    if this.items[idx] == nil {
        return
    }
    this.items[idx][innerIdx] = nil
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */