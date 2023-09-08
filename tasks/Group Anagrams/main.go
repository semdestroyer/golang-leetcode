package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {

	newStrs := make(map[string][]string)
	keys := make([]string, 0)
	for _, str := range strs {
		sorted := sortRunes(str)
		newStrs[sorted] = append(newStrs[sorted], str)
	}
	for k := range newStrs {
		keys = append(keys, k)
	}
	res := make([][]string, len(keys))
	for i, key := range keys {
		res[i] = newStrs[key]
	}
	return res
}

func sortRunes(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1; j++ {
			if runes[j] > runes[j+1] {
				tmp := runes[j+1]
				runes[j+1] = runes[j]
				runes[j] = tmp
			}

		}
	}
	return string(runes)
}
