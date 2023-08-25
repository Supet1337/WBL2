package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	file := NewTextFile()
	file.ParseFlags()
	file.SetArgs()
	file.Read()
	res, err := grep(file)
	if err != nil {
		log.Fatal(err)
	}
	switch result := res.(type) {
	case []string:
		for _, line := range result {
			fmt.Println(line)
		}
	case map[int]string:
		for i, line := range result {
			fmt.Println(i, line)
		}
	default:
		fmt.Println(res)
	}
}

func grep(file *TextFile) (interface{}, error) {
	var prefix, postfix string
	if file.i {
		prefix = "(?i)"
	}
	if file.F {
		prefix = "^"
		postfix = "$"
	}
	regex, err := regexp.Compile(prefix + file.word + postfix)
	if err != nil {
		return "Error", fmt.Errorf("uncorrected regular expression")
	}

	switch {
	case file.A != 0:
		for i, line := range file.dateStr {
			if regex.MatchString(line) {
				if i+file.A <= len(file.dateStr) {
					return file.dateStr[i : file.A+2], nil
				}
				return file.dateStr[i:], nil
			}
		}

	case file.B != 0:
		for i, line := range file.dateStr {
			if regex.MatchString(line) {
				if i-file.B >= 0 {
					return file.dateStr[i-file.B : i+1], nil
				}
				return file.dateStr[:i+1], nil
			}
		}

	case file.C != 0:
		for i, line := range file.dateStr {
			left, right := 0, 0
			if regex.MatchString(line) {
				if i-file.C >= 0 {
					left = i - file.C
				} else {
					left = 0
				}
				if i+file.C <= len(file.dateStr) {
					right = file.C + 1
				} else {
					right = len(file.dateStr) - 1
				}

				return file.dateStr[left : right+1], nil
			}
		}

	case file.c:
		count := 0
		for _, line := range file.dateStr {
			if regex.MatchString(line) {
				count++
			}
		}
		if file.Flags.v {
			return len(file.dateStr) - count, nil
		}
		return count, nil

	case file.v:
		invertStr := []string{}
		for _, line := range file.dateStr {
			if !regex.MatchString(line) {
				invertStr = append(invertStr, line)
			}
		}
		return invertStr, nil

	case file.n:
		strNum := make(map[int]string)
		for i, line := range file.dateStr {
			if regex.MatchString(line) {
				strNum[i+1] = line
			}
		}
		return strNum, nil

	default:
		deflt := []string{}
		for _, line := range file.dateStr {
			if regex.MatchString(line) {
				deflt = append(deflt, line)
			}
		}
		return deflt, nil
	}
	return nil, nil
}

type Flags struct {
	A int  //-A - "after" печатать +N строк после совпадения
	B int  //-B - "before" печатать +N строк до совпадения
	C int  //-C - "context" (A+B) печатать ±N строк вокруг совпадения
	c bool //-c - "count" (количество строк)
	i bool //-i - "ignore-case" (игнорировать регистр)
	v bool //-v - "invert" (вместо совпадения, исключать)
	F bool //-F - "fixed", точное совпадение со строкой, не паттерн
	n bool //-n - "line num", печатать номер строки
}

// NewFlag - создание объекта с флагами
func NewFlag() *Flags {
	return &Flags{}
}

// ParseFlags - парсинг флагов и путей для файлов из коммандной строки
func (f *Flags) ParseFlags() {
	flag.IntVar(&f.A, "A", 0, "show +N strings after match")
	flag.IntVar(&f.B, "B", 0, "show +N strings before match")
	flag.IntVar(&f.C, "C", 0, "(A+B) show ±N strings around match")
	flag.BoolVar(&f.c, "c", false, "numbers of strings")
	flag.BoolVar(&f.i, "i", false, "ignore case")
	flag.BoolVar(&f.v, "v", false, "instead of match, exclude")
	flag.BoolVar(&f.F, "f", false, "exact string match, not a pattern")
	flag.BoolVar(&f.n, "n", false, "print line number")
	flag.Parse()
}

type TextFile struct {
	*Flags
	dateStr []string
	path    string
	word    string
}

// NewTextFile - создание TextFile
func NewTextFile() *TextFile {
	return &TextFile{
		Flags: NewFlag(),
	}
}

func (file *TextFile) SetArgs() {
	file.word = flag.Arg(0)
	file.path = flag.Arg(1)
}

func (t *TextFile) Read() {
	file, err := os.Open(t.path)
	if err != nil {
		log.Fatal("open file error: ", err)
	}
	defer file.Close()

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("read file error: ", err)
	}

	t.dateStr = strings.Split(string(dataBytes), "\n")
}
