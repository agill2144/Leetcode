type dataNode struct {
    val string
    time int
}

type TimeMap struct {
    data map[string][]*dataNode
}


func Constructor() TimeMap {
    return TimeMap{
        data: map[string][]*dataNode{},
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    node := &dataNode{val: value, time: timestamp}
    this.data[key] = append(this.data[key], node)
}


func (this *TimeMap) Get(key string, timestamp int) string {
    // we want the greatest time from left side of passed-in time
    // --------------.---.--.---t----.----
    nodes, ok := this.data[key]
    if !ok {return ""}
    
    left := 0
    right := len(nodes)-1
    var ans *dataNode
    for left <= right {
        mid := left + (right-left)/2
        if nodes[mid].time <= timestamp {
            ans = nodes[mid]
            left = mid+1
        } else {
            right = mid-1
        }
    }
    if ans == nil {return ""}
    return ans.val
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */