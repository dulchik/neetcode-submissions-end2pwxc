import "maps"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS := make(map[rune]int)
    countT := make(map[rune]int)

    for _, r := range s {   
        countS[r]++
    }

    for _, r := range t {   
        countT[r]++
    }

    return maps.Equal(countS, countT)
}
