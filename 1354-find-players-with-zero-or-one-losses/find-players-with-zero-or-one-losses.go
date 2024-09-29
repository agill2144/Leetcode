func findWinners(matches [][]int) [][]int {
    // no-sorting , map each player to an idx
    // each idx == playerID
    // therefore allocate a big enough array such that 
    // all possible playerID's can be mapped to an idx
    scores := make([][]int, 100001)
    for i := 0; i < len(matches); i++ {
        wPlayer := matches[i][0]
        lPlayer := matches[i][1]
        if scores[wPlayer] == nil {scores[wPlayer] = make([]int, 2)}
        if scores[lPlayer] == nil {scores[lPlayer] = make([]int, 2)}
        scores[wPlayer][0]++
        scores[lPlayer][1]++
    }
    out := make([][]int, 2)
    for i := 0; i < len(scores); i++ {
        if scores[i] != nil {
            losses := scores[i][1]
            if losses == 0 {out[0] = append(out[0],i)}
            if losses == 1 {out[1] = append(out[1],i)}
        }
    }
    return out

}
// func findWinners(matches [][]int) [][]int {
//     scores := map[int][]int{} // { id : [wins, loss] }
//     for i := 0; i < len(matches); i++ {
//         wPlayer := matches[i][0]
//         lPlayer := matches[i][1]
//         if _, ok := scores[wPlayer]; !ok {scores[wPlayer] = make([]int, 2)}
//         if _, ok := scores[lPlayer]; !ok {scores[lPlayer] = make([]int, 2)}
//         scores[wPlayer][0]++
//         scores[lPlayer][1]++
//     }
//     out := make([][]int, 2)
//     for player, counts := range scores {
//         losses := counts[1]
//         if losses == 0 {out[0] = append(out[0], player)}
//         if losses == 1 {out[1] = append(out[1], player)}
//     }
//     for _, l := range out {sort.Ints(l)}
//     return out
// }