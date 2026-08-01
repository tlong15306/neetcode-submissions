func twoSum(nums []int, target int) []int {
    // Create a map to store numbers and their indices
    seen := make(map[int]int)
    
    for i, num := range nums {
        complement := target - num
        
        // If we've already seen the complement, we found our pair!
        if j, ok := seen[complement]; ok {
            return []int{j, i}
        }
        
        // Otherwise, add the current number and its index to the map
        seen[num] = i
    }
    
    return nil
}