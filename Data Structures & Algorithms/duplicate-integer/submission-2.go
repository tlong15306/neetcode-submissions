func hasDuplicate(nums []int) bool {
    // Khởi tạo "HashSet" trong Go bằng map với value là struct{} rỗng
    seen := make(map[int]struct{})
    
    for _, num := range nums {
        // Cú pháp _, ok giúp kiểm tra xem 'num' đã tồn tại trong map (Set) chưa
        if _, ok := seen[num]; ok {
            return true // Đã tồn tại -> Trùng lặp
        }
        
        // Nếu chưa tồn tại, thêm 'num' vào Set
        seen[num] = struct{}{}
    }
    
    return false

}
