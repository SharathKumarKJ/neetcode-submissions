type Pair struct {
	Key   int
	Value int
}

func topKFrequent(nums []int, k int) []int {
    dict := make(map[int]int)
	for _, item := range nums {
		dict[item]++
	}

	pairItem := make([]Pair, 0, len(dict))
	for key, value := range dict {
		pairItem = append(pairItem, Pair{Key: key, Value: value})
	}

	sort.Slice(pairItem, func(i, j int) bool {
		return pairItem[i].Value > pairItem[j].Value
	})

	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, pairItem[i].Key)
	}
	return result
}
