func twoSum(nums []int, target int) []int {
	n := len(nums)
	hash := make(map[int]int, n)
	for i := 0; i < n; i++ {
		minus := target - nums[i]
		if index, found := hash[minus]; found {
			return []int{index, i}
		}
		hash[nums[i]] = i
	}
	return []int{}
}
