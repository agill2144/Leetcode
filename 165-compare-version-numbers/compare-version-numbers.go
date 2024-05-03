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