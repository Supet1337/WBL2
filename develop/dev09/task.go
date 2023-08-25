package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	Wget("https://losst.pro/komanda-cut-linux")
}

func Wget(url string) {
	client := &http.Client{}
	resp, err := client.Get(url)
	fmt.Println(resp.Status)
	if err != nil {
		panic(err)
	}
	Write(resp)
}

func Write(resp *http.Response) {
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("create file error: ", err)
	}
	defer file.Close()
	buffer := bufio.NewWriter(file)
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		panic(err)
	}
}
