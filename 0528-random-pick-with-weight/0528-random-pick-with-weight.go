type Solution struct {
    prefixSum []int
}


func Constructor(w []int) Solution {
    prefixSum := make([]int, len(w))
    total := 0
    for i := 0; i < len(w); i++ {
        total += w[i]
        prefixSum[i] = total
    }
    return Solution{prefixSum: prefixSum}
}


func (this *Solution) PickIndex() int {
    n := len(this.prefixSum)
    if n == 0 {
        return -1
    }
    total := this.prefixSum[n-1]
    randNum := rand.Intn(total)
    return this.BinarySearch(randNum)
}

func (this *Solution) BinarySearch(target int) int {
    left, right := 0, len(this.prefixSum)-1
    for left < right {
        mid := left + (right - left) / 2
        if this.prefixSum[mid] > target {
            right = mid
        } else {
            left = mid+1
        }
    }
    return left
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */