package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {

		}
	case "add":
	case "remove":
	default:
	}
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list -- View the existing music lib
		lib add <name> <artist> <source> <type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
	`)

	//lib = library.NewMusicmanager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter commend ->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {

		} else if tokens[0] == "play" {

		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
