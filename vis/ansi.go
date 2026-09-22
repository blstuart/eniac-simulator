package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var msgin *bufio.Scanner

func main() {
	msgin = bufio.NewScanner(os.Stdin)
	reset()
	cls()
	seperator()
	toinp()
	initdraw()
	fmt.Println("update")
	procupd()
}

func reset() {
	fmt.Fprint(os.Stderr, "\x1bc")
}

func cls() {
	fmt.Fprint(os.Stderr, "\x1b[2J")
}

func toinp() {
	fmt.Fprint(os.Stderr, "\x1b[23;24r")
	fmt.Fprint(os.Stderr, "\x1b[23;1H")
	fmt.Fprint(os.Stderr, "\x1b[0J")
}

func seperator() {
	fmt.Fprint(os.Stderr, "\x1b[22;1H")
	fmt.Fprint(os.Stderr, "\x1b[7m")
	fmt.Fprint(os.Stderr, "                                                                               ")
	fmt.Fprint(os.Stderr, "\x1b[0m")
}

func save() {
	fmt.Fprint(os.Stderr, "\x1b7")
}

func restore() {
	fmt.Fprint(os.Stderr, "\x1b[23;24r")
	fmt.Fprint(os.Stderr, "\x1b8")
}

func initdraw() {
	save()
	fmt.Fprint(os.Stderr, "\x1b[1;1H")
	fmt.Fprintln(os.Stderr, "Init: 000000000000000   Cyc: 00 C")
	fmt.Fprintln(os.Stderr, "MP: 0000000000 00000000000000000000 0000000000")
	fmt.Fprintln(os.Stderr, "Accs:")
	for i := 0; i < 10; i++ {
		fmt.Fprintln(os.Stderr,
			"P 0000000000 0000000000 0 000000000000   P 0000000000 0000000000 0 000000000000")
	}
	fmt.Fprintln(os.Stderr, "Div: 0 0 00000000 n+")
	fmt.Fprintln(os.Stderr, "Mult:  0 000000000000000000000000 0 0")
	fmt.Fprintln(os.Stderr, "FT1: 00000000000 00 -3 0 0 0")
	fmt.Fprintln(os.Stderr, "FT2: 00000000000 00 -3 0 0 0")
	fmt.Fprintln(os.Stderr, "FT3: 00000000000 00 -3 0 0 0")
	fmt.Fprintln(os.Stderr, "CT: 000000000000000000000000000000")
	fmt.Fprint(os.Stderr, "Card: ")
	restore()
}

func procupd() {
	for {
		if !msgin.Scan() {
			reset()
			os.Exit(0)
		}
		s := msgin.Text()
		p := strings.Split(s, " ")
		switch p[0] {
		case "exit":
			reset()
			return
		case "up":
			fmt.Println("update")

		// Initiating unit
		case "init":
			save()
			fmt.Fprintf(os.Stderr, "\x1b[1;7H%s", p[1])
			restore()

		// Cycle unit
		case "cy":
			cy, _ := strconv.Atoi(p[1])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[1;30H%02d", cy)
			restore()
		case "cm":
			save()
			fmt.Fprintf(os.Stderr, "\x1b[1;33H%s", p[1])
			restore()

		// Accumulators
		case "ad":
			unit, _ := strconv.Atoi(p[1])
			dig, _ := strconv.Atoi(p[2])
			d, _ := strconv.Atoi(p[3])
			row := 4 + unit / 2
			col := 41 * (unit % 2) + dig
			save()
			if dig == 0 {
				if d == 0 {
					fmt.Fprintf(os.Stderr, "\x1b[%d;%dHP", row, col + 1)
				} else {
					fmt.Fprintf(os.Stderr, "\x1b[%d;%dHM", row, col + 1)
				}
			} else {
				fmt.Fprintf(os.Stderr, "\x1b[%d;%dH%d", row, col + 2, d)
			}
			restore()
		case "ac":
			unit, _ := strconv.Atoi(p[1])
			dig, _ := strconv.Atoi(p[2])
			d, _ := strconv.Atoi(p[3])
			row := 4 + unit / 2
			col := 41 * (unit % 2) + dig
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;%dH%d", row, col + 13, d)
			restore()
		case "ar":
			unit, _ := strconv.Atoi(p[1])
			rep, _ := strconv.Atoi(p[2])
			row := 4 + unit / 2
			col := 25 + 41 * (unit % 2)
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;%dH%d", row, col, rep)
			restore()
		case "af":
			unit, _ := strconv.Atoi(p[1])
			prog, _ := strconv.Atoi(p[2])
			val, _ := strconv.Atoi(p[3])
			row := 4 + unit / 2
			col := 27 + 41 * (unit % 2) + prog
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;%dH%d", row, col, val)
			restore()

		// Divider/Square Rooter
		case "d":
			flags := ""
			if p[4][0] == '1' {
				flags += " divff"
			}
			if p[4][1] == '1' {
				flags += " clrff"
			}
			if p[4][2] == '1' {
				flags += " coinff"
			}
			if p[4][3] == '1' {
				flags += " dpg"
			}
			if p[4][4] == '1' {
				flags += " ng"
			}
			if p[4][5] == '1' {
				flags += " psrcff"
			}
			if p[4][6] == '1' {
				flags += " denomff"
			}
			if p[4][7] == '1' {
				flags += " n+"
			}
			if p[4][8] == '1' {
				flags += " n-"
			}
			if p[4][9] == '1' {
				flags += " qa"
			}
			if p[4][10] == '1' {
				flags += " SAC"
			}
			if p[4][11] == '1' {
				flags += " -2"
			}
			if p[4][12] == '1' {
				flags += " -1"
			}
			if p[4][13] == '1' {
				flags += " NAC"
			}
			if p[4][14] == '1' {
				flags += " dA"
			}
			if p[4][15] == '1' {
				flags += " na"
			}
			if p[4][16] == '1' {
				flags += " da"
			}
			if p[4][17] == '1' {
				flags += " dg"
			}
			if p[4][18] == '1' {
				flags += " npg"
			}
			if p[4][19] == '1' {
				flags += " +2"
			}
			if p[4][20] == '1' {
				flags += " +1"
			}
			if p[4][21] == '1' {
				flags += " sa"
			}
			if p[4][22] == '1' {
				flags += " dS"
			}
			if p[4][23] == '1' {
				flags += " nb"
			}
			if p[4][24] == '1' {
				flags += " db"
			}
			if p[4][25] == '1' {
				flags += " A1"
			}
			if p[4][26] == '1' {
				flags += " A2"
			}
			if p[4][27] == '1' {
				flags += " A3"
			}
			if p[4][28] == '1' {
				flags += " A4"
			}
			save()
			fmt.Fprintf(os.Stderr, "\x1b[14;6H\x1b[K%s %s %s%s", p[1], p[2], p[3], flags)
			restore()

		// Multiplier
		case "m":
			save()
			fmt.Fprintf(os.Stderr, "\x1b[15;7H%s", s[2:])
			restore()
		// Master programmer
		case "mps":
			stage, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[2;%dH%d", stage + 5, val)
			restore()
		case "mpi":
			stage, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[2;%dH%d", stage + 37, val)
			restore()
		case "mpd":
			decade, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[2;%dH%d", decade + 16, val)
			restore()

		// Function tables
		case "ftar":
			unit, _ := strconv.Atoi(p[1])
			arg, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;18H%02d", unit + 16, arg)
			restore()
		case "ftr":
			unit, _ := strconv.Atoi(p[1])
			ring, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;21H%+d", unit + 16, ring)
			restore()
		case "ftad":
			unit, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;24H%d", unit + 16, val)
			restore()
		case "ftsu":
			unit, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;26H%d", unit + 16, val)
			restore()
		case "ftse":
			unit, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[%d;28H%d", unit + 16, val)
			restore()

		// Constant Transmitter
		case "ct":
			prog, _ := strconv.Atoi(p[1])
			ff, _ := strconv.Atoi(p[2])
			save()
			fmt.Fprintf(os.Stderr, "\x1b[19;%dH%d", prog + 5, ff)
			restore()

		// Punch
		case "punch":
			save()
			fmt.Fprintf(os.Stderr, "\x1b[21;1H%s", s[6:])
			restore()
		}
	}
}
