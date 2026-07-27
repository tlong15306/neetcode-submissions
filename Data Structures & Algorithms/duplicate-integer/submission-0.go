func hasDuplicate(nums []int) bool {
    // Khởi tạo map: Key là số nguyên (giá trị), Value là số nguyên (số lần xuất hiện)
    counts := make(map[int]int) 
    
    for _, num := range nums {
        counts[num]++ // Tăng số lần xuất hiện của số 'num' lên 1
        
        // Nếu số này xuất hiện nhiều hơn 1 lần -> Phát hiện trùng lặp
        if counts[num] > 1 {
            return true
        }
    }
    
    // Đã duyệt hết mảng mà không có số nào lặp lại
    return false
}

