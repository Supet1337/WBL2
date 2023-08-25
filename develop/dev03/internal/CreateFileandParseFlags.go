package internal

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Flag struct {
	K int
	R bool
	U bool
	N bool
}

// NewFlag - создание объекта с флагами
func NewFlag() *Flag {
	return &Flag{}
}

// ParseFlags - парсинг флагов и путей для файлов из коммандной строки
func (f *Flag) ParseFlags() {
	flag.IntVar(&f.K, "k", 1, "указание строки для соритровки")
	flag.BoolVar(&f.N, "n", false, "сортировка по числовому значению")
	flag.BoolVar(&f.R, "r", false, "сортировка в обратном порядке")
	flag.BoolVar(&f.U, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()
}

// TextFile - структура для работы с файлами
type TextFile struct {
	*Flag
	dateStr []string
	Sorted  bool
	inPath  string
	outPath string
}

// NewTextFile - создание TextFile
func NewTextFile() *TextFile {
	return &TextFile{
		Flag: NewFlag(),
	}
}

// Read - ф-я чтения из файла, переданного в коммандноцй строки
func (t *TextFile) Read() {
	file, err := os.Open(t.inPath)
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

// Write - ф-я записи в файл новый отсортированных данных
func (t *TextFile) Write() {
	file, err := os.Create(t.outPath)
	if err != nil {
		log.Fatal("create file error: ", err)
	}
	defer file.Close()

	for _, line := range t.dateStr {
		if _, err = file.WriteString(line + "\n"); err != nil {
			log.Fatal("write file error: ", err)
		}
	}
}

// SetInputPath - путь исходного файла
func (t *TextFile) SetInputPath() {
	t.inPath = flag.Arg(0)
}

// SetOutputPath - путь выходного файла
func (t *TextFile) SetOutputPath() {
	t.outPath = flag.Arg(1)
}
