package main

import (
	"fmt"

	"github.com/yasukotelin/ansiescape/sgr"
)

var (
	clearEsc = sgr.MakeString(sgr.Clear)
)

func main() {
	printAnsiColors()
	fmt.Print("\n\n")
	print256colors()
}

func printAnsiColors() {
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgBlack), "Black", clearEsc)
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgRed), "Red", clearEsc)
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgGreen), "Green", clearEsc)
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgYellow), "Yellow", clearEsc)
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgBlue), "Blue", clearEsc)
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgMagenta), "Magenta", clearEsc)
	fmt.Printf("%v%s%v ", sgr.MakeString(sgr.BgCyan), "Cyan", clearEsc)
	fmt.Printf("%v%s%v", sgr.MakeString(sgr.BgWhite), "White", clearEsc)
}

func print256colors() {
	for i := 0; i < 256; i++ {
		bgEsc := sgr.MakeString(sgr.BgColor(i))
		fmt.Printf("%v%03d%v", bgEsc, i, clearEsc)
		if (i+1)%16 == 0 {
			fmt.Print("\n\n")
		} else {
			fmt.Print(" ")
		}
	}
}
