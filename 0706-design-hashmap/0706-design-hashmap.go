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
                - Why is our nested array using a different hash func ?
                - When we used mod , on the outter array, we had collision
                - then using the same again in the inner array, we will again have same collision in the nested array, therefore we cannot use
                    the same mod hash func for the inner array

*/

const arrSize int = 1001
const innerArrSize int = 1001

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