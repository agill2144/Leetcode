
type Solution struct {
    prefixSum []int
}


func Constructor(w []int) Solution {
    prefixSum := make([]int, len(w))
    for i := 0; i < len(w); i++ {
        if i == 0 {
            prefixSum[i] = w[i]
        } else {
            prefixSum[i] = prefixSum[i-1] + w[i]
        }
    }
    for i := 0; i < len(w); i++ {
        fmt.Println(prefixSum[i])
    }
    return Solution{prefixSum}
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