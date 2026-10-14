func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))

	for _, ch := range s {
		// Dùng nháy đơn ' cho rune
		if ch == '{' || ch == '(' || ch == '[' {
			stack = append(stack, ch)
		} else {
			// Nếu gặp ngoặc đóng mà stack rỗng -> sai
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1] // Lấy phần tử đỉnh stack

			// Kiểm tra cặp ngoặc có khớp nhau không
			if (ch == ')' && top != '(') ||
				(ch == '}' && top != '{') ||
				(ch == ']' && top != '[') {
				return false
			}

			// Khớp thì pop phần tử đỉnh ra khỏi stack
			stack = stack[:len(stack)-1]
		}
	}

	// Hợp lệ khi tất cả ngoặc mở đều đã được đóng hết
	return len(stack) == 0
}