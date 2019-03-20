package main

type runeArray []rune

func (r runeArray) Less(i, j int) bool { return r[i] < r[j] }
func (r runeArray) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runeArray) Len() int           { return len(r) }

func sortString(str string) string {
	r := make(runeArray, len(str))
	for i, v := range str {
		r[i] = v
	}
	sort.Sort(r)
	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string = make([][]string, 0)
	dict := make(map[string]int)
	for i := 0; i < len(strs); i++ {
		str := strs[i]
		sortedStr := sortString(str)
		if idx, ok := dict[sortedStr]; !ok {
			res = append(res, []string{str})
			dict[sortedStr] = len(res) - 1
		} else {
			res[idx] = append(res[idx], str)
		}
	}
	return res
}
