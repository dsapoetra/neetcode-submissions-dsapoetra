import "slices"

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, w := range strs {
    	r := []rune(w)
    	slices.Sort(r)
    	key := string(r)
    	groups[key] = append(groups[key], w)
	}
	// {"abc": ["abc", "bca", "cab"]}
	result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result

}
