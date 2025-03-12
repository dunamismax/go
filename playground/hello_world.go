// main.go
package main

import (
	"fmt"
	"strings"
	"time"
)

// ====================================================================
// ASCII Art Configuration
// ====================================================================
// To customize the output, simply replace the contents of the asciiArt
// slice below with your own ASCII art, one line per slice element.
// This approach safely handles any characters (including backticks) in your art.
// ====================================================================
var asciiArt = []string{
	".s    s.                                          ",
	"      SS. .s5SSSs.  .s        .s        .s5SSSs.  ",
	"sS    S%S       SS.                           SS. ",
	"SS    S%S sS    `:; sS        sS        sS    S%S ",
	"SSSs. S%S SSSs.     SS        SS        SS    S%S ",
	"SS    S%S SS        SS        SS        SS    S%S ",
	"SS    `:; SS        SS        SS        SS    `:; ",
	"SS    ;,. SS    ;,. SS    ;,. SS    ;,. SS    ;,. ",
	":;    ;:' `:;;;;;:' `:;;;;;:' `:;;;;;:' `:;;;;;:' ",
	"                                                  ",
	".s s.  s.                                         ",
	"   SS. SS. .s5SSSs.  .s5SSSs.  .s        .s5SSSs. ",
	"sS S%S S%S       SS.       SS.                 SS.",
	"SS S%S S%S sS    S%S sS    S%S sS        sS    S%S",
	"SS S%S S%S SS    S%S SS .sS;:' SS        SS    S%S",
	"SS S%S S%S SS    S%S SS    ;,  SS        SS    S%S",
	"SS `:; `:; SS    `:; SS    `:; SS        SS    `:;",
	"SS ;,. ;,. SS    ;,. SS    ;,. SS    ;,. SS    ;,.",
	"`:;;:'`::' `:;;;;;:' `:    ;:' `:;;;;;:' ;;;;;;;:'",
}

// clearScreen clears the terminal using ANSI escape codes.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// animateAscii gradually reveals the provided ASCII art from left to right.
// Each character is printed with an ANSI color code that creates a gradient
// effect (from red to yellow) based on its horizontal position.
func animateAscii(ascii string) {
	// Split the ASCII art into lines.
	lines := strings.Split(ascii, "\n")

	// Determine the maximum width among all lines.
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Reveal the art one column at a time.
	for col := 1; col <= maxWidth; col++ {
		clearScreen()
		for _, line := range lines {
			// Print each character up to the current column with a gradient.
			for j, ch := range line {
				if j < col {
					// Compute a ratio for the gradient (red to yellow).
					ratio := float64(j) / float64(maxWidth-1)
					green := int(ratio * 255)
					// ANSI escape: \033[38;2;R;G;0m sets the foreground color.
					fmt.Printf("\033[38;2;255;%d;0m%c\033[0m", green, ch)
				} else {
					// Print a space for unrevealed columns.
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		// Control the animation speed (75ms delay between each column).
		time.Sleep(75 * time.Millisecond)
	}
}

func main() {
	// Join the slice of strings into a single string with newline separators.
	art := strings.Join(asciiArt, "\n")
	animateAscii(art)
	// Pause briefly at the end so the final art remains visible.
	time.Sleep(500 * time.Millisecond)
}