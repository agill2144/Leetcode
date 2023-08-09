
// use heaps to be build largest possible palindrome
func largestPalindromic(num string) string {
 
    freqMap := map[string]int{}
    for i := 0; i < len(num); i++ {
        freqMap[string(num[i])]++
    }
    
    mx := &maxHeap{items: []string{}}
    usedKeys := &maxHeap{items: []string{}}

    for k , _ := range freqMap {
        heap.Push(mx, k)
    }
    
    out := make([]*string, len(num))
    left := 0
    right := len(out)-1
    
    
    for mx.Len() != 0 { 
        top := heap.Pop(mx).(string)
        freq := freqMap[top]
        
        // when we have left over counts for a key, push it to other pq
        if freq == 1 || freq % 2 != 0 {
            heap.Push(usedKeys, top)
        }
        
        // the only time where we cannot use a key to build the palindrome is when freq is only 1
        if freq == 1 {continue}
        // or
        // when a value to push to output is 0 and this is the first value in output, cannot have leading 0s
        if top == "0" {
            if left == 0 {
                delete(freqMap, top)
                continue
            }
        }
        
        if freq % 2 == 0 {
            for freq != 0 {
                out[left] = &top
                out[right] = &top
                left++
                right--
                freq-=2
            }
            delete(freqMap, top)
        } else if freq > 1 {
            n := freq-1
            k := 0
            for k < n {
                out[left] = &top
                out[right] = &top
                left++
                right--
                k+=2                
            }
            freqMap[top] = freq-n
        }
    }
    // keys we could not use the first time, because of their odd or 1 freq count
    // take the biggest one
    if usedKeys.Len() >= 1 {
        top := heap.Pop(usedKeys).(string)
        if out[left] == nil {
            out[left] = &top
        } else {
            out[right] = &top
        }
    }
    // build out the final string
    outStr := new(strings.Builder)
    for i := 0; i < len(out); i++ {
        if out[i] != nil {
            outStr.WriteString(*out[i])
        }
    }
    if outStr.String() =="" {return "0"}
    return outStr.String()
}

type maxHeap struct {
    items []string
}
func (m *maxHeap) Less(i,j int) bool {
    return strings.Compare(m.items[i], m.items[j]) == 1
}
func (m *maxHeap) Swap(i,j int) {
    m.items[i],m.items[j] = m.items[j],m.items[i]
}
func (m *maxHeap) Len() int {
    return len(m.items)
}
func (m *maxHeap) Push(x interface{}) {
    m.items = append(m.items, x.(string))
}
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}