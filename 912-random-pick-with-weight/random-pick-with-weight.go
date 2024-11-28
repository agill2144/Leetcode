/*
     0 1 2 3  
    [2,1,4,3]

    instead of prob of ith idx = i/4
    its now based on weight

    total weight = 10
    instead of 4 choices, we now have to consider 10 idxs
    0 -> 2/10
    1 -> 1/10
    2 -> 4/10
    3 -> 3/10

    [0,0,1,2,2,2,2,3,3,3]
    prob -> 1/10

    truncate the space    
    [0(2), 1(1), 2(4), 3(3)]
 
     0 1 2 3
    [2,3,7,10]

    idx 0 -> [0,2) -> 0,1
    idx 1 -> [2,3) -> 2
    idx 2 -> [3,7) -> 3,4,5,6
    idx 3 -> [7,10) -> 7,8,9
*/
type Solution struct {
    items []int    
}


func Constructor(w []int) Solution {
    items := []int{}
    for i := 0; i < len(w); i++ {
        items = append(items,w[i])
        if i-1 >= 0{items[i] += items[i-1]}
    }
    return Solution{items}
}


func (this *Solution) PickIndex() int {
    n := len(this.items)
    r := rand.Intn(this.items[n-1])
    left := 0
    right := n-1
    ans := -1
    for left <= right {
        mid := left + (right-left)/2
        if this.items[mid] > r {
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */