package May_LeetCoding_Challenge

func findAnagrams(s string, p string) []int {
	compare := func(visited, window []uint8) bool {
		for i := range visited {
			if visited[i] != window[i] {
				return false
			}
		}
		return true
	}

	var (
		visited, window = make([]uint8, 26), make([]uint8, 26)
		sLen, pLen      = len(s), len(p)
		aChar           = uint8('a')
	)
	if sLen == 0 || pLen == 0 || sLen < pLen {
		return nil
	}
	for i := 0; i < pLen; i++ {
		window[p[i]-aChar]++
		visited[s[i]-aChar]++
	}
	var anagrams []int
	for i := pLen; i < sLen; i++ {
		if compare(visited, window) {
			anagrams = append(anagrams, i-pLen)
		}
		visited[s[i]-aChar]++
		visited[s[i-pLen]-aChar]--
	}
	if compare(visited, window) {
		anagrams = append(anagrams, sLen-pLen)
	}
	return anagrams
}
