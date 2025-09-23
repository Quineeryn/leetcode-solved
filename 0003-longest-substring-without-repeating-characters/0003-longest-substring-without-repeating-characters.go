func lengthOfLongestSubstring(s string) int {
    last := map[byte]int{}
    start, best := 0, 0
    for i := 0; i < len(s); i++ {
        if p, ok := last[s[i]]; ok && p > start {
            start = p
        }
        if i-start+1 > best {
            best = i - start + 1
        }
        last[s[i]] = i + 1
    }
    return best
}
