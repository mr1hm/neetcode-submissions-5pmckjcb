func isAnagram(s string, t string) bool {
	seen := map[byte]int{}

	for i := 0; i < len(s); i++ {
		seen[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		seen[t[i]]--
	}

	for _, v := range seen {
		if v != 0 {
			return false
		}
	}

	return true
}
