func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        // Start j at i + 1 to avoid using the same element twice
        for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
                // Return the indices as a slice
                return []int{i, j}
            }
		}
	}
	return nil
}
