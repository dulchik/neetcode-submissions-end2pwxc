import "maps"

func isAnagram(s string, t string) bool {
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
