package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(str string) (string, error) {
	runestr := []rune(str)
	var output string
	for i := 0; i < len(runestr); i++ {
		if unicode.IsDigit(runestr[i]) {
			if i == 0 {
				return "", errors.New("Incorrect")
			}
			intsymbol, _ := strconv.Atoi(string(runestr[i]))
			output += strings.Repeat(string(runestr[i-1]), intsymbol-1)
		} else {
			output += string(runestr[i])
		}
	}
	return output, nil
}

func main() {
	result, err := Unpack("45")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
