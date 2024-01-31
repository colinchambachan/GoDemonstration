func twoSum(nums []int, target int) []int {
    // create hashmap of stored variable
    valIndexMap := make(map[int]int)

    // iterate through and see if target - val is in map
    for index, val := range nums{
        // target-val found in hashmap
		elem, ok := valIndexMap[target-val]

        if ok {
            // return index of target-val and val 
			return []int{elem, index}
        } else {
            valIndexMap[val] = index
        }

    }
    return []int{}
}