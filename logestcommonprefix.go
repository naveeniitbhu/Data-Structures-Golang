package main

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}

	for i:=0; i<len(strs); i++ {
		ch := strs[0][i]
		for j:=1; j<len(strs); j++ {
			if i>= len(strs[j]) || ch != strs[j][i] {
				return prefix
			}
		}
		prefix += string(ch)
	}
	return prefix
}