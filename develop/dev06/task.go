package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ct struct {
	fields    string
	delimiter string
	separated bool
}

func main() {
	flags := NewFlag()
	cutRun(flags)
}

func NewFlag() *Ct {
	ct := Ct{}
	flag.StringVar(&ct.fields, "f", "", "List of fields to cut")
	flag.StringVar(&ct.delimiter, "d", "\t", "Set custom delimeter")
	flag.BoolVar(&ct.separated, "s", false, "Get only separated strings")
	flag.Parse()
	return &ct
}

func cut(row string, ct *Ct) (string, error) {
	var res strings.Builder
	fields := make(map[int]bool)
	var delimeter string
	if ct.delimiter != "\t" {
		if len(ct.delimiter) == 1 {
			delimeter = ct.delimiter
		} else {
			return "", fmt.Errorf("you could set only one character for delimeter")
		}
	}
	if ct.fields != "" {
		rangeR := strings.Split(ct.fields, ",")
		for _, dPart := range rangeR {
			dRange := strings.Split(strings.TrimSpace(dPart), "-")
			if len(dRange) == 2 {
				dLeft, err := strconv.Atoi(dRange[0])
				if err != nil {
					return "", fmt.Errorf("invalid left value %s", dLeft)
				}
				dRight, err := strconv.Atoi(dRange[1])
				if err != nil {
					return "", fmt.Errorf("invalid right value %s", dRight)
				}
				if dLeft < 1 || dLeft > dRight {
					return "", fmt.Errorf("your range has started from 0 or left border more than right border")
				}
				for i := dLeft; i <= dRight; i++ {
					fields[i] = true
				}
			} else {
				numOfField, err := strconv.Atoi(strings.TrimSpace(dPart))
				if err != nil {
					return "", fmt.Errorf("invalid field value %s", dPart)
				}
				fields[numOfField] = true
			}
		}
	}
	sliceOfRows := strings.Split(row, delimeter)
	if len(sliceOfRows) == 1 && ct.separated {
		return "", nil
	}
	isDelim := false
	for i, val := range sliceOfRows {
		_, ok := fields[i+1]
		if ok {
			if isDelim {
				res.WriteString(delimeter + val + "\n")
			} else {
				res.WriteString(val + "\n")
				isDelim = true
			}
		}
	}
	return res.String(), nil
}

func cutRun(config *Ct) {
	var str strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str.WriteString(scanner.Text())
	}
	result, err := cut(str.String(), config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(result)
}
