func hasDuplicate(nums []int) bool {
    
    var dict = make(map[int]bool)

    for _,item:= range nums {
        if dict[item] == true{
            return true
        }else{
            dict[item]=true
        }
    }
    return false
}
