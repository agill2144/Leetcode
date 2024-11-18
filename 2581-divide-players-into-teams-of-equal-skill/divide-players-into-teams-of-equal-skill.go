func dividePlayers(skill []int) int64 {
    n := len(skill)
    teams := n/2
    sum := 0
    freq := map[int]int{}
    for i := 0; i < n; i++ {
        sum += skill[i]
        freq[skill[i]]++
    }
    if sum % teams != 0 {return int64(-1)}
    skillPerTeam := sum / teams
    var total int64 = 0
    for i := 0; i < len(skill); i++ {
        diff := skillPerTeam - skill[i]
        if freq[diff] == 0 {return int64(-1)}
        freq[diff]--; if freq[diff] == 0 {delete(freq, diff)}
        total += (int64(skill[i]) * int64(diff))
    }
    // why the divide by 2 here?
    // because every compliment search in the 2nd pass is added twice
    // skill = 6
    // and we have a pair 2+4, then 4+2 will come up too
    // therefore dupes, therefore each compliment is added twice, therefore removing by dividing the total by 2
    // Return half of total chemistry (as each pair is counted twice)
    return total/2
}
