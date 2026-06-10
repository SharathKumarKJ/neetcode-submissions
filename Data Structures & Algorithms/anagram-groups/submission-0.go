func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, word := range strs {
		chars := []rune(word)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)
		dict[key] = append(dict[key], word)
	}

	var result = make([][]string, 0)
	for _, item := range dict {
		result = append(result, item)
	}
	return result
}
