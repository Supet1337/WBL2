package internal

import (
	"sort"
	"strings"
)

// RunSort - ф-я запуска соритровки
func RunSort(file *TextFile) {
	if file.U {
		file.dateStr = DeleteDuplicate(file.dateStr)
	}
	if file.N {
		file.dateStr = SortByNumbers(file.dateStr, file.Flag)
	} else if file.K > 1 {
		file.dateStr = SortByColumn(file.dateStr, file.Flag)
	}
	if file.R {
		file.dateStr = Reverse(file.dateStr)
	}

}

func Reverse(str []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	return str
}

func SortByColumn(str []string, flag *Flag) []string {
	mp := make(map[string][]string)
	keys := make([]string, 0, len(str))
	resultData := make([]string, 0, len(str))

	for _, line := range str {
		column := strings.Split(line, " ")
		key := column[flag.K-1]
		keys = append(keys, key)
		mp[key] = []string{line}

	}
	sort.Strings(keys)

	for _, k := range keys {
		resultData = append(resultData, mp[k]...)
	}

	return resultData
}

func DeleteDuplicate(str []string) []string {
	mp := make(map[string]struct{})

	for _, el := range str {
		if _, exist := mp[el]; !exist {
			mp[el] = struct{}{}
		}
	}

	data := make([]string, 0, len(mp))

	for el := range mp {
		data = append(data, el)
	}
	return data
}

func SortByNumbers(str []string, flag *Flag) []string {
	mp := make(map[string][]string)
	keys := make([]string, 0, len(str))
	resultData := make([]string, 0, len(str))

	for _, line := range str {
		column := strings.Split(line, " ")
		key := column[0]
		keys = append(keys, key)
		mp[key] = []string{line}

	}
	sort.Strings(keys)

	for _, k := range keys {
		resultData = append(resultData, mp[k]...)
	}

	return resultData
}
