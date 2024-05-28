package main

import "fmt"

func ClearScreen() {
	fmt.Print("\033[H\033[2J") // Overwrites the terminal buffer https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
}
