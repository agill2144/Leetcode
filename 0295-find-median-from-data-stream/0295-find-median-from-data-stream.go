/*
    approach: use 2 heaps
    - Median always has to do with 1 or more middle element in a sorted array
    - Therefore we need to make sure we can access atleast of the middle elements in a sorted array QUICKLY
    - For example;
        - [1,2,3,4,5,6,7]
        - odd lenght list, therefore median is middle element (4)
        - [1,2,3,4,5,6]
        - even length list, therefore median is avg of 2 middle elements ( (3+4)/2 = 3.5)
        - Notice how we only really cared about 3,4 or just 3, the middle element(s)
    - But how can we further optimize o(n) time for addNum() ?
    - What else helps sort things?
        - sorting algos
        - heaps
    - Can we use a max or min heap?
        - max heap will only give us the largest number
        - min heap on its own will give us the smallest number
    - But we can use BOTH and have the array split into 2 heaps ( min and max heaps )
    - minHeap will be for the right half
    - maxHeap will be the the left half
    - For example;
        - [1,2,3,4] [5,6,7]
        - maxHeap    minHeap
        - maxHeap will guarantee largest number of left half is at the root therefore 4 will be at the root of maxHeap
        - minHeap will guarantee smallest number of right half is at the root, therefore 5 will be at the root of minHeap
    - Now we have the middle elements AT THE SURFACE OF BOTH HEAPS!
    - How do we continue to keep the correct middle elements at the surface of both heaps???
    - How do we decide which heap addNum() will push to??
    
        - [1,2,3,4] [5,6,7]
        - maxHeap    minHeap
    
    - If maxHeap is supposed to keep the largest element of the left half
    - If minHeap is suppsed to keep the smallest element of the right half
    - Then its fair to say that the top of maxHeap SHOULD be <= top of minHeap
    - So that left half only contains elements <= right half!
    - Thats one thing to check ALWAYS!
    
    - We also know that if someone keeps adding element to the left half ALWAYS
    - Then simply pushing elements to left half will mess up our middle elements
    - We need to make sure that we can rebalance our heaps after every add SUCH that the median stays on the surface of both heaps
    - if we dont do the above, then median will get burried deep into some heap and findMedian will become time expensive again.
    - Therefore we need to rebalance both heaps such that the size of both of them is almost the same ( can be equal or 1 element off incase we have a odd list size )
    - So if left half has 2 more elements than of right half
        - pop an element from the left half
        - push that element to right half
    - vice versa for the left half
    
    - Now our heaps will be balanced such that the middle elements are always on the surface of both heaps
    - We also will make sure that top of maxHeap <= top of minHeap
        - if not, pop from left and push to right heap
    
    - Now, how do we decide which heap addNum() will push elements to.
    - It does not matter, because of our 2 checks we will do.
        - top of maxHeap <= top of minHeap
        - well balanced heaps
    - So pick any heap ( say left and push blindly ) and run checks and rebalance as needed
    
    
    time:
        addNum(): o(5logN) or o(logN)
        findMedian() : o(1)
    space: 
        addNum() : o(1)
        findMedian() : o(1)        
*/

type MedianFinder struct {
    left *maxHeap
    right *minHeap
}


func Constructor() MedianFinder {
    return MedianFinder{
        left: &maxHeap{items: []int{}},
        right: &minHeap{items: []int{}},
    }
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.right, num)
    
    // ensure left top <= right top
    if this.left.Len() != 0 && this.right.Len() != 0 && 
    this.right.items[0] < this.left.items[0] {
        heap.Push(this.left, heap.Pop(this.right))
    }
    
    // ensure its balanced, so that the median is at the top of surface of both heaps
    // otherwise median could get burried somewhere deep inside a heap
    if this.left.Len() > this.right.Len() + 1 {
        heap.Push(this.right, heap.Pop(this.left))
    } else if this.right.Len() > this.left.Len() {
        heap.Push(this.left, heap.Pop(this.right))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    if this.left.Len() > this.right.Len() {
        return float64(this.left.items[0])
    } else if this.right.Len() > this.left.Len() {
        return float64(this.right.items[0])
    }
    return float64(this.left.items[0]+this.right.items[0]) / 2.0
}


type minHeap struct {items []int}
func (m *minHeap) Less(i, j int) bool { return m.items[i] < m.items[j] }
func (m *minHeap) Swap(i, j int)  { m.items[i], m.items[j] = m.items[j], m.items[i] }
func (m *minHeap) Len() int  { return len(m.items) }
func (m *minHeap) Push(x interface{}) { m.items = append(m.items, x.(int)) }
func (m *minHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}

type maxHeap struct {items []int}
func (m *maxHeap) Less(i, j int) bool { return m.items[i] > m.items[j] }
func (m *maxHeap) Swap(i, j int)  { m.items[i], m.items[j] = m.items[j], m.items[i] }
func (m *maxHeap) Len() int  { return len(m.items) }
func (m *maxHeap) Push(x interface{}) { m.items = append(m.items, x.(int)) }
func (m *maxHeap) Pop() interface{} {
    out := m.items[len(m.items)-1]
    m.items = m.items[:len(m.items)-1]
    return out
}


/*
    approach: brute force
    - make addNum insert n in its correct idx position
        - append incoming number to a list blindly and then
        - bubble the last appended item to its correct idx position in o(n) time
    - time:
        addNum() : o(n)
        findMedian() : o(1)
    - space: 
        addNum() : o(1)
        findMedian() : o(1)        
*/
// type MedianFinder struct {
//     items []int
// }

// func Constructor() MedianFinder {
//     return MedianFinder{
//         items: []int{},
//     }
// }

// func (this *MedianFinder) AddNum(num int)  {
//     this.items = append(this.items, num)
//     // bubble the last appended item to its correct idx position
//     for i := len(this.items)-2; i >= 0; i-- {
//         if this.items[i+1] < this.items[i] {
//             this.items[i+1], this.items[i] = this.items[i], this.items[i+1]
//         } else {
//             break
//         }
//     }        
// }


// func (this *MedianFinder) FindMedian() float64 {
//     mid := (len(this.items)-1)/2
//     if len(this.items) % 2 == 0 {
//         return float64(this.items[mid]+this.items[mid+1])/2
//     }
//     return float64(this.items[mid])
// }
