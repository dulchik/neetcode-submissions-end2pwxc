import "maps"

func isAnagram(s string, t string) bool {
    countS := make(map[string]int)
    countT := make(map[string]int)

    for _, c := range strings.Split(s, "") {   
        if _, ok := countS[c]; !ok {
            countS[c] = 1
        } else {
            countS[c] += 1
        }
    }

    for _, c := range strings.Split(t, "") {   
        if _, ok := countT[c]; !ok {
            countT[c] = 1
        } else {
            countT[c] += 1
        }
    }

    return maps.Equal(countS, countT)
}
