
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i, num := range nums {
		item := target - num
		_, exists := dict[item]
		if exists {
			return []int{dict[item], i}
		}
		dict[num] = i
	}
	return nil
}

