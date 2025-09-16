func isAnagram(s string, t string) bool {
    if len([]rune(s)) != len([]rune(t)){
        return false
    }

    cnt := make(map[rune]int)

    for _, r := range s {cnt[r]++}
    for _, r := range t {cnt[r]--

        if cnt[r] < 0 {
            return false
        }
    }    
    return true
}