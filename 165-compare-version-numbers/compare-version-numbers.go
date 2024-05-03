// get chunks using 2 ptrs, instead of comparing each ptr val
// once you have chunks, compare the chunks
// return in which ever case we can get out, otherwise continue
// the reason why we are able to get away with || operator in our while loop
// is because "If a version number does not specify a revision at an index, then treat the revision as 0."
// therefore if a revision does not exist, we can imply 0 for that revision and still compare and be a valid comparison
func compareVersion(version1 string, version2 string) int {
    v1 := 0; v1Rev := 0
    v2 := 0; v2Rev := 0

    for v1 < len(version1) || v2 < len(version2) {
        v1Rev = 0
        v2Rev = 0
        for v1 < len(version1) && version1[v1] != '.' {
            v1Rev = v1Rev * 10 + int(version1[v1]-'0')
            v1++
        }
        v1++ // escape the dot
        for v2 < len(version2) && version2[v2] != '.' {
            v2Rev = v2Rev * 10 + int(version2[v2]-'0')
            v2++
        }
        v2++ // escape the dot
        if v1Rev < v2Rev {
            return -1
        } else if v1Rev > v2Rev {
            return 1
        }
    }
    return 0
}