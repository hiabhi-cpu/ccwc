package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println("Hello")
	if len(os.Args) < 1 {
		log.Fatal("Usage: ccw <command> [args...]")
		return
	}
	if !strings.Contains(os.Args[1], "-") {
		if !strings.Contains(os.Args[1], ".") {
			check(errors.New("Give proper file name"))
		}
		//give all outputs
	}
	fileName := os.Args[len(os.Args)-1]
	if !strings.Contains(fileName, ".") {
		check(errors.New("Give proper file name"))
	}

	cmdList := make(map[string]func(string))

	cmdList["-l"] = lineCommand
	cmdList["-w"] = wordCommand
	cmdList["-c"] = charCommand
	cmds := []string{"-l", "-w", "-c"}
	argsCmd := os.Args[1 : len(os.Args)-1]
	if len(argsCmd) == 0 {
		for _, c := range cmdList {
			c(fileName)
		}
	}
	for _, c := range cmds {
		if slices.Contains(argsCmd, c) {
			f, e := cmdList[c]
			if !e {
				check(errors.New("Give proper command name"))
			}
			f(fileName)
		}
	}

	// for _, val := range os.Args[1 : len(os.Args)-1] {

	// }

	fmt.Println(fileName)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDataInBytes(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	check(err)
	return data
}

func charCommand(fileName string) {
	data := getDataInBytes(fileName)
	fmt.Print(len(data), "\t")
}

func lineCommand(fileName string) {
	data := getDataInBytes(fileName)
	cnt := strings.Count(string(data), "\n")
	fmt.Print(cnt, "\t")
}

func wordCommand(fileName string) {
	data := getDataInBytes(fileName)
	cnt := 0

	inward := false
	for _, ch := range data {
		if unicode.IsSpace(rune(ch)) {
			inward = false
		} else {
			if !inward {
				cnt++
				inward = true
			}
		}
	}

	fmt.Print(cnt, "\t")
}
