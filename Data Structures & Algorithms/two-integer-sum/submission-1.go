func twoSum(nums []int, target int) []int {
	m := map[int]int{}
    for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		idx, ok := m[complement]
		if ok {
			if i < idx {
				return []int{i, idx}
			}
			return []int{idx, i}
		}

		m[nums[i]] = i
	}

	return nil
}
