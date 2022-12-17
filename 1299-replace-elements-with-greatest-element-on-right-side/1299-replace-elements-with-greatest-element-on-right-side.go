// time: o(n)
// space: o(1)

func replaceElements(arr []int) []int {
    if arr == nil || len(arr) == 0 {return nil}
    // feels like dp, can be solved bottom up with the same intuition
    // find the best answer by looking at the smallest subproblem
    // the smallest problem we have and we have a definitive constant answer is -1 for the last idx
    // then, traverse back and look right to answer the question for i-1
    // that will be max(arr[i+1], out[i+1])

    out := make([]int, len(arr))
    out[len(out)-1] = -1
    if len(out) == 1 {return out}
    
    for i := len(out)-2; i >= 0; i-- {
        out[i] = max(arr[i+1], out[i+1])
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}