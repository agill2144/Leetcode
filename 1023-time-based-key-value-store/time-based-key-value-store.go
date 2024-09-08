type valNode struct {
    time int
    val string
}

type TimeMap struct {
    data map[string][]*valNode
}


func Constructor() TimeMap {
    return TimeMap{data:map[string][]*valNode{}}    
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.data[key] = append(this.data[key], &valNode{timestamp, value})

}


func (this *TimeMap) Get(key string, timestamp int) string {
    // we want the rightmost on left side of timestamp
    ans := ""
    vals := this.data[key]
    if len(vals) == 0 {
        return ans
    }

    left := 0
    right := len(this.data[key])-1
    for left <= right {
        mid := left + (right-left)/2
        if vals[mid].time <= timestamp {
            ans = vals[mid].val
            if vals[mid].time == timestamp {break}
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return ans
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */