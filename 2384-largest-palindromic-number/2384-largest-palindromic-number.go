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
        
        if freq == 1 || freq % 2 != 0 {
            heap.Push(usedKeys, top)
            if freq == 1 {continue}
        }
        if top == "0" {
            if left == 0 {
                delete(freqMap, top)
                continue
            }
        }
        
        if freq % 2 == 0 {
            k := 0
            for k < freq {
                out[left] = &top
                out[right] = &top
                left++
                right--
                k+=2
            }
            delete(freqMap, top)
        } else if freq > 1 {
            n := freq-1
            k := 0
            for k < freq-1 {
                out[left] = &top
                out[right] = &top
                left++
                right--
                k+=2                
            }
            freqMap[top] = freq-n
        }
    }
    
    if usedKeys.Len() >= 1 {
        top := heap.Pop(usedKeys).(string)
        if out[left] == nil {
            out[left] = &top
        } else {
            out[right] = &top
        }
    }
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