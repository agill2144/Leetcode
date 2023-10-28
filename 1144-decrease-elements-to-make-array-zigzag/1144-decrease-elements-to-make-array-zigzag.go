// greedy greedy greedy
// time = o(n) + o(n) + o(n)
// space = o(n)
func movesToMakeZigzag(nums []int) int {
    n := len(nums)
    odd := make([]int, n)
    copy(odd, nums)
    oddMoves := math.MaxInt64
    evenMoves := math.MaxInt64

    for i := 0; i < n; i++ {
        if i % 2 != 0 {
            curr := odd[i]
            left := math.MinInt64
            right := math.MinInt64
            if i != 0 {left = odd[i-1]}
            if i != n-1 {right = odd[i+1]}
            if curr > left && curr > right {continue}

            if left >= curr {
                if oddMoves == math.MaxInt64 {oddMoves = 0}
                newVal := curr-1
                oddMoves += (left - newVal)
                odd[i-1] = newVal
            }

            if right != math.MinInt64 && right >= curr {
                if oddMoves == math.MaxInt64 {oddMoves = 0}
                newVal := curr-1
                oddMoves += (right-newVal)
                odd[i+1] = newVal
            }
        } else {
            curr := nums[i]
            left := math.MinInt64
            right := math.MinInt64
            if i != 0 {left = nums[i-1]}
            if i != n-1 {right = nums[i+1]}

            if curr > left && curr > right {continue}

            if left != math.MaxInt64 && left >= curr {
                if evenMoves == math.MaxInt64 {evenMoves = 0}
                newVal := curr-1
                evenMoves += (left-newVal)
                nums[i-1] = newVal
            }

            if right >= curr {
                if evenMoves == math.MaxInt64 {evenMoves = 0}
                newVal := curr-1
                evenMoves += (right-newVal)
                nums[i+1] = newVal
            }            
        }
    }
    
    if evenMoves == math.MaxInt64 {evenMoves = 0}
    if oddMoves == math.MaxInt64 {oddMoves = 0}
    return min(evenMoves, oddMoves)
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}

func min(x, y int) int {
    if x < y {return x}
    return y
}