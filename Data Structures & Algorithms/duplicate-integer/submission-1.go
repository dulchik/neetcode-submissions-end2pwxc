func hasDuplicate(nums []int) bool {
    dup := make(map[int]int)
    for _, num := range nums {
        dup[num] += 1
    }
    for i, _ := range dup {
        if dup[i] > 1 {
            return true
        }
    }
    return false
}
