func compareVersion(version1 string, version2 string) int {
    id1, id2 := 0, 0
    l1, l2 := len(version1), len(version2)

    for id1 < l1 || id2 < l2{
        n1, n2 := 0,0
        for id1 < l1 && version1[id1]!='.'{
            n1 = n1*10 + int(version1[id1]-'0')
            id1++
        }
        id1++

        for id2 < l2 && version2[id2]!='.'{
            n2 = n2*10 + int(version2[id2]-'0')
            id2++
        }
        id2++

        if n1 > n2 {
            return 1
        } else if n1 < n2 {
            return -1
        }
    }
    return 0
}