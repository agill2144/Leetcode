/*
    approach: brute force
    - make addNum insert n in its correct idx position
    - time:
        addNum() : o(n)
        findMedian() : o(1)
    - space: 
        addNum() : o(1)
        findMedian() : o(1)        
*/
type MedianFinder struct {
    items []int
}

func Constructor() MedianFinder {
    return MedianFinder{
        items: []int{},
    }
}

func (this *MedianFinder) AddNum(num int)  {
    this.items = append(this.items, num)
    // bubble the last appended item to its correct idx position
    for i := len(this.items)-2; i >= 0; i-- {
        if this.items[i+1] < this.items[i] {
            this.items[i+1], this.items[i] = this.items[i], this.items[i+1]
        } else {
            break
        }
    }        
}


func (this *MedianFinder) FindMedian() float64 {
    mid := (len(this.items)-1)/2
    if len(this.items) % 2 == 0 {
        return float64(this.items[mid]+this.items[mid+1])/2
    }
    return float64(this.items[mid])
}
