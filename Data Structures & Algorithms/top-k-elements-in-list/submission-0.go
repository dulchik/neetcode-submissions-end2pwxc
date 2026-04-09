func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, n := range nums {
		freq[n]++
	}
 
	// 2) convert to a slice of pairs
    type pair struct {
        val int
        cnt int
    }
    pairs := make([]pair, 0, len(freq))
    for v, c := range freq {
        pairs = append(pairs, pair{val: v, cnt: c})
    }

    // 3) sort by frequency desc
    sort.Slice(pairs, func(i, j int) bool {
        if pairs[i].cnt == pairs[j].cnt {
            return pairs[i].val < pairs[j].val // tie-breaker (optional)
        }
        return pairs[i].cnt > pairs[j].cnt
    })

    // 4) collect top k values
    res := make([]int, 0, k)
    for i := 0; i < k && i < len(pairs); i++ {
        res = append(res, pairs[i].val)
    }
    return res
}
