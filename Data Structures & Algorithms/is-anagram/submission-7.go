func isAnagram(s string, t string) bool {
    if len(t)!=len(s){
        return false
    }
    counts := make(map[rune]int)
    for _,ch:= range s{
         counts[ch]++
    }
    for _,ch:= range t{
        count, exists := counts[ch]
        if !exists || count == 0 {
            return false
        }
        counts[ch]--
    }
    return true
}
