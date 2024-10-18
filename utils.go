package main

import (
	"fmt"
	"time"
)

func delay(milliseconds time.Duration) {
	time.Sleep(milliseconds * time.Millisecond)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func moveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", y, x)
}
