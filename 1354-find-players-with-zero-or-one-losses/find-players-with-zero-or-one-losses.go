func findWinners(matches [][]int) [][]int {
    scores := map[int][]int{} // { id : [wins, loss] }
    for i := 0; i < len(matches); i++ {
        wPlayer := matches[i][0]
        lPlayer := matches[i][1]
        if _, ok := scores[wPlayer]; !ok {scores[wPlayer] = make([]int, 2)}
        if _, ok := scores[lPlayer]; !ok {scores[lPlayer] = make([]int, 2)}
        scores[wPlayer][0]++
        scores[lPlayer][1]++
    }
    out := make([][]int, 2)
    for player, counts := range scores {
        _, losses := counts[0], counts[1]
        if losses == 1 {out[1] = append(out[1], player)}
        if losses == 0  {out[0] = append(out[0], player)}
    }
    for _, l := range out {sort.Ints(l)}
    return out
}