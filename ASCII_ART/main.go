package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ONLY ONE BUDDY!")
		return
	}

	FileLine, err := os.Open("standard.txt")
	if err != nil {
		fmt.Print("no file exists")
		return
	}
	defer FileLine.Close()

	LineScan := bufio.NewScanner(FileLine)
	lines := []string{}
	AllLines := [][]string{}
	for LineScan.Scan() {
		lines = append(lines, LineScan.Text())
		if len(lines) == 9 {
			AllLines = append(AllLines, lines)
			lines = nil
		}
	}

	input := os.Args[1]

	parts := strings.Split(input, "\\n")

	str2 := ""
	for _, part := range parts {
		//whenever i see "" means there was /n there before i split
		if part == "" {
			str2 += "\n"
			continue
		}

		for line := 1; line < 9; line++ {
			for _, ch := range part {
				if ch >= ' ' && ch <= '~' {
					ch = ch - ' '
					str2 += AllLines[ch][line]
				}
			}
			//newline after each line of the ASCII art
			str2 += "\n"
		}
	}
	fmt.Print(str2)
}