package main

import (
	"fmt"

	"github.com/yasukotelin/ansiescape/sgr"
)

var (
	clearEsc = sgr.EscSeq(sgr.Clear)
)

func main() {
	printAnsiColors()
	fmt.Print("\n\n")
	print256colors()
}

func printAnsiColors() {
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgBlack), "Black", clearEsc)
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgRed), "Red", clearEsc)
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgGreen), "Green", clearEsc)
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgYellow), "Yellow", clearEsc)
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgBlue), "Blue", clearEsc)
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgMagenta), "Magenta", clearEsc)
	fmt.Printf("%v%s%v ", sgr.EscSeq(sgr.BgCyan), "Cyan", clearEsc)
	fmt.Printf("%v%s%v", sgr.EscSeq(sgr.BgWhite), "White", clearEsc)
}

func print256colors() {
	for i := 0; i < 256; i++ {
		bgEsc := sgr.EscSeq(sgr.BgColor(i))
		fmt.Printf("%v%03d%v", bgEsc, i, clearEsc)
		if (i+1)%16 == 0 {
			fmt.Print("\n\n")
		} else {
			fmt.Print(" ")
		}
	}
}
