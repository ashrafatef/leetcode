package main


func longestCommonPrefix(strs []string) string {
	prefix := ""
	flag := true
	index := 0
	character := ""
	for flag == true && len(strs) != 0 {
		for v := 0; v < len(strs); v++ {
			if len(strs[v]) <= index {
				flag = false
				break
			}
			if v == 0 {
				character = string(strs[v][index])
			}
			if string(strs[v][index]) != string(character) {
				flag = false
				break
			}
		}
		prefix = prefix + string(character)
		index++
	}
	if len(prefix) == 0 {
		return ""
	}
	if flag == false {
		return prefix[:len(prefix)-1]
	}
	return prefix
}

