func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

    for _, w := range strs {
        b := []byte(w)
        sort.Slice(b, func(i,j int) bool { return b[i] < b [j] } )

        key := string(b)

        groups[key] = append(groups[key],w)
    }
    res := make([][]string, 0, len(groups))

    for _, g := range groups { res = append(res, g)}
    return res

}