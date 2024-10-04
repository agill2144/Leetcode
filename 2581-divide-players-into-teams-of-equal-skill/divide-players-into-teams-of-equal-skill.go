func dividePlayers(skill []int) int64 {
    n := len(skill)
    total := 0
    freq := map[int]int{}
    for i := 0; i < len(skill); i++ {
        total += skill[i]
        freq[skill[i]]++
    }
    numTeams := n/2
    if total % numTeams != 0 {return int64(-1)}
    var totalChem int64 = 0
    skillPerTeam := total / numTeams
    for i := 0; i < n; i++ {
        diff := skillPerTeam - skill[i]
        c := freq[diff]
        if c == 0 {return -1}
        freq[diff]--
        if c == 1 {delete(freq, diff)}
        totalChem += (int64(skill[i]) * int64(diff) )
    }

    return totalChem/2
}