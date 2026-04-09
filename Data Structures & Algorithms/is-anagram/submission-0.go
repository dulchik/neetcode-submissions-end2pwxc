func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    dups := make(map[rune]int)
    dupt := make(map[rune]int)
    for i, char := range s {
        dups[char]++
        dupt[rune(t[i])]++
    }

    for k, v := range dups {
        if dupt[k] != v {
            return false
        }
    }
    return true
}
