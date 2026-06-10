
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	charCount:= make(map[rune]int)

	for _,ch:= range s {
		charCount[ch]++
	}
	for _,ch:=range t {
		if charCount[ch]>0 {
			charCount[ch]--
		}else{
			return false
		}
	}
	return true 
}
