func arrayRankTransform(arr []int) []int {
    idxs := map[int][]int{}
    for i := 0; i < len(arr); i++ {
        num := arr[i]
        idxs[num] = append(idxs[num], i)
    }
    sort.Ints(arr)
    out := make([]int, len(arr))
    rank := 1
    for i := 0; i < len(arr); i++ {
        if i > 0 && arr[i] != arr[i-1] {rank++}
        num := arr[i]
        indices := idxs[num]
        out[indices[0]] = rank
        indices = indices[1:]
        if len(indices) == 0 {delete(idxs, num)} else {idxs[num] = indices}
    }
    return out
}