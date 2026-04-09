func groupAnagrams(strs []string) [][]string {
	storage := make(map[[26]int][]string)
	for _, s := range strs {
		var key [26]int
		for _, symbol := range s {
			key[int(symbol)-97]++
		}

		storage[key] = append(storage[key], s)
	}
	var result [][]string
	for _, group := range storage {
		result = append(result, group)
	}

	return result
}