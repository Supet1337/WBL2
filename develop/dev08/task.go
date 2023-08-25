package main

import (
	"bufio"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for {
		path, err := os.Getwd()
		if err != nil {
			return
		}
		fmt.Print(path, ">>> ")
		if stdin.Scan() {
			cmd := strings.Split(stdin.Text(), "|")
			start(cmd)
		}
	}
}

func pwd() {
	workDir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(workDir)
}

func cd(command []string) {
	err := os.Chdir(command[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func psFunc(...[]string) {
	fmt.Printf("%7s %-9s\n", "PID", "TTY")
	res, err := ps.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	for _, v := range res {
		fmt.Printf("%7d %-9s\n", v.Pid(), v.Executable())
	}
}

func echo(command []string) {
	for _, el := range command[1:] {
		fmt.Println(el)
	}
}

func kill(command []string) {
	pid, err := strconv.Atoi(command[1])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	proc.Kill()
}

func start(data []string) {
	for _, command := range data {
		commandSl := strings.Split(command, " ")
		switch commandSl[0] {
		case "pwd":
			pwd()
		case "ps":
			psFunc(commandSl)
		case "kill":
			kill(commandSl)
		case "cd":
			cd(commandSl)
		case "echo":
			echo(commandSl)
		default:
			fmt.Println("No command")
		}

	}
}
