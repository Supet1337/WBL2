package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Anagram(arr *[]string) *map[string][]string {
	//Приводим буквы к нижнему регистру
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = strings.ToLower((*arr)[i])
	}

	sortedWord := make([]string, len(*arr))
	copy(sortedWord, *arr)

	//Сортируем каждое слово
	for i, _ := range sortedWord {
		sortedWord[i] = SortString(sortedWord[i])
	}

	//Ищем уникальные ключи
	uniqueUnsKeys := make(map[string][]string)

	for _, el := range sortedWord {
		if _, exist := uniqueUnsKeys[el]; !exist {
			uniqueUnsKeys[el] = []string{}
		}
	}

	uniqueSrtKeys := make(map[string][]string)

	//сортируем ключи обратно
	for unsK := range uniqueUnsKeys {
		for _, k := range *arr {
			if SortString(k) == unsK {
				uniqueSrtKeys[k] = []string{}
				break
			}
		}
	}

	//Формируем множество
	for _, k := range *arr {
		for srtK := range uniqueSrtKeys {
			if SortString(k) == SortString(srtK) {
				uniqueSrtKeys[srtK] = append(uniqueSrtKeys[srtK], k)
			}
		}
	}

	//Сортируем
	for s := range uniqueSrtKeys {
		sort.Slice(uniqueSrtKeys[s], func(i, j int) bool {
			return uniqueSrtKeys[s][i] < uniqueSrtKeys[s][j]
		})
	}

	return &uniqueSrtKeys
}

func main() {
	arr := []string{"пятаК", "пятка", "тяпка", "листок", "слиток", "столик"}
	mp := Anagram(&arr)
	fmt.Println(mp)

}
