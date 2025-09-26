func longestCommonPrefix(strs []string) string {
   if len(strs) == 0 { return "" }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        for !hasPrefix(strs[i], prefix) {
            if len(prefix) == 0 { return "" }
            prefix = prefix[:len(prefix)-1]
        }
    }
    return prefix
}
func hasPrefix(s, p string) bool {
    if len(p) > len(s) { return false }
    return s[:len(p)] == p
}