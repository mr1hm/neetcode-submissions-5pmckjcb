func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	out := [][]string{}
	for i := 0; i < len(strs); i++ {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		
		m[string(s)] = append(m[string(s)], strs[i])
	}

	for _, v := range m {
		out = append(out, v)
	}

	return out
}
