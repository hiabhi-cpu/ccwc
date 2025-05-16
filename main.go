package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	cmdList := make(map[string]func([]byte))

	cmdList["-l"] = lineCommand
	cmdList["-w"] = wordCommand
	cmdList["-c"] = charCommand
	cmdList["-m"] = mCommand
	cmds := []string{"-l", "-w", "-c", "-m"}

	if len(os.Args) == 1 {
		data, err := readFromConsole()
		check(err)
		if len(data) == 0 {
			check(errors.New("Usage: ccw <command> [args...]"))
		}
		noCommands(data, cmdList)
		fmt.Println()
	} else {
		if strings.Contains(os.Args[len(os.Args)-1], ".") { //has file name
			fileName := os.Args[len(os.Args)-1]
			argsCmd := os.Args[1 : len(os.Args)-1]
			if strings.Contains(os.Args[1], "-") {
				hasCommands(getDataInBytes(fileName), cmds, argsCmd, cmdList)
			} else {
				noCommands(getDataInBytes(fileName), cmdList)
			}
			fmt.Println(fileName)
		} else { //has no file name
			data, err := readFromConsole()
			check(err)
			if len(data) == 0 {
				check(errors.New("Usage: ccw <command> [args...]"))
			}
			argsCmd := os.Args[1:len(os.Args)]
			if strings.Contains(os.Args[1], "-") {
				hasCommands(data, cmds, argsCmd, cmdList)
			} else {
				noCommands(data, cmdList)
			}
			fmt.Println()
		}

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func hasCommands(data []byte, cmds, argsCmd []string, cmdList map[string]func([]byte)) {
	fmt.Println("has cmds")
	for _, c := range argsCmd {
		if !slices.Contains(cmds, c) {
			check(errors.New("Give correct command"))
		}
	}
	for _, c := range cmds {
		if slices.Contains(argsCmd, c) {
			f, e := cmdList[c]
			if !e {
				check(errors.New("Give proper command name"))
			}
			f(data)
		}
	} //give a conter and check if wrong command was given then panic
}

func noCommands(data []byte, cmdList map[string]func([]byte)) {
	for i, c := range cmdList {
		if i == "-m" {
			continue
		}
		c(data)
	}
}

func getDataInBytes(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	check(err)
	return data
}

func charCommand(data []byte) {
	fmt.Print(len(data), "\t")
}

func lineCommand(data []byte) {
	cnt := strings.Count(string(data), "\n")
	fmt.Print(cnt, "\t")
}

func wordCommand(data []byte) {
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

func mCommand(data []byte) {
	fmt.Print(utf8.RuneCount(data), "\t")
}

func readFromConsole() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)

	return io.ReadAll(reader)
}
