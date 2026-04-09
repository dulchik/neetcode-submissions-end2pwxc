func groupAnagrams(strs []string) [][]string {
    type pair struct { h1, h2 uint64 }
    memo := make(map[pair][]string)
    
    primes := [26]uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    
    const mod1 = 1e9 + 7
    const mod2 = 1e9 + 9

    for _, s := range strs {
        var h1, h2 uint64 = 1, 1
        for i := 0; i < len(s); i++ {
            p := primes[s[i]-'a']
            h1 = (h1 * p) % mod1
            h2 = (h2 * p) % mod2
        }
        
        key := pair{h1, h2}
        memo[key] = append(memo[key], s)
    }

    ans := make([][]string, 0, len(memo))
    for _, v := range memo {
        ans = append(ans, v)
    }
    return ans
}