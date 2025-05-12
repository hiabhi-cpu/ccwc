package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello")
	if len(os.Args) < 2 {
		log.Fatal("Usage: ccw <command> [args...]")
		return
	}
	cmdName := os.Args[1]
	fmt.Print(cmdName)
	dat, err := os.ReadFile("text.txt")
	check(err)
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
