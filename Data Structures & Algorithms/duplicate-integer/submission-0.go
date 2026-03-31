func hasDuplicate(nums []int) bool {
    seen := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] {
			return true
		}

		seen[nums[i]] = true
	}

	return false
}
