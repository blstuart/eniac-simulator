package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var step [10]struct {
	stage      int
	di, i, cdi chan pulse
	o          [6]chan pulse
	csw        int
	inff       int
	kludge     bool
}

var dec [20]struct {
	di    chan pulse
	val   int
	carry bool
	lim   [6]int
}

var assoc [8]int
var mpupdate chan int

func mpstat() string {
	var s string

	for i := 0; i < 10; i++ {
		s += fmt.Sprintf("%d", step[i].stage)
	}
	s += " "
	for i := 0; i < 20; i++ {
		s += fmt.Sprintf("%d", dec[i].val)
	}
	s += " "
	for i := 0; i < 10; i++ {
		if step[i].inff < 10 {
			s += fmt.Sprintf("%d", step[i].inff)
		} else {
			s += "*"
		}
	}
	return s
}

func mpreset() {
	for i := 0; i < 20; i++ {
		dec[i].di = nil
		for j := 0; j < 6; j++ {
			dec[i].lim[j] = 0
		}
	}
	for i := 0; i < 10; i++ {
		step[i].di = nil
		step[i].i = nil
		step[i].cdi = nil
		for j := 0; j < 6; j++ {
			step[i].o[j] = nil
		}
		step[i].csw = 0
		step[i].kludge = false
	}
	for i := 0; i < 8; i++ {
		assoc[i] = 0
	}
	mpclear()
	mpupdate <- 1
}

func mpclear() {
	for i := 0; i < 20; i++ {
		dec[i].val = 0
		dec[i].carry = false
	}
	for i := 0; i < 10; i++ {
		step[i].stage = 0
		step[i].inff = 0
	}
}

func mpplug(jack string, ch chan pulse) {
	var n int

	if unicode.IsDigit(rune(jack[0])) {
		fmt.Sscanf(jack, "%ddi", &n)
		dec[20-n].di = ch
	} else {
		prog := int(jack[0] - 'A')
		if prog > 7 { // Correct for the lack of program I
			prog--
		}
		switch jack[1:] {
		case "di":
			step[prog].di = ch
		case "i":
			step[prog].i = ch
		case "cdi":
			step[prog].cdi = ch
		default:
			if len(jack) < 3 {
				fmt.Println("Invalid MP jack", jack)
				break
			}
			fmt.Sscanf(jack[1:], "%do", &n)
			step[prog].o[n-1] = ch
		}
	}
	mpupdate <- 1
}

func mpctl(ch chan [2]string) {
	var d, s int

	for {
		ctl := <-ch
		switch ctl[0][0] {
		case 'a', 'A':
			switch ctl[0] {
			case "a20", "A20":
				if ctl[1] == "b" || ctl[1] == "B" {
					assoc[0] = 1
				} else {
					assoc[0] = 0
				}
			case "a18", "A18":
				if ctl[1] == "c" || ctl[1] == "C" {
					assoc[1] = 1
				} else {
					assoc[1] = 0
				}
			case "a14", "A14":
				if ctl[1] == "d" || ctl[1] == "D" {
					assoc[2] = 1
				} else {
					assoc[2] = 0
				}
			case "a12", "A12":
				if ctl[1] == "e" || ctl[1] == "E" {
					assoc[3] = 1
				} else {
					assoc[3] = 0
				}
			case "a10", "A10":
				if ctl[1] == "g" || ctl[1] == "G" {
					assoc[4] = 1
				} else {
					assoc[4] = 0
				}
			case "a8", "A8":
				if ctl[1] == "h" || ctl[1] == "H" {
					assoc[5] = 1
				} else {
					assoc[5] = 0
				}
			case "a4", "A4":
				if ctl[1] == "j" || ctl[1] == "J" {
					assoc[6] = 1
				} else {
					assoc[6] = 0
				}
			case "a2", "A2":
				if ctl[1] == "k" || ctl[1] == "K" {
					assoc[7] = 1
				} else {
					assoc[7] = 0
				}
			default:
				fmt.Printf("Invalid MP associator switch %s\n", ctl[0])
				break
			}

		case 'd', 'D':
			fmt.Sscanf(ctl[0], "d%ds%d", &d, &s)
			n, _ := strconv.Atoi(ctl[1])
			dec[20-d].lim[s-1] = n
		case 'c', 'C':
			s = int(ctl[0][1] - 'A')
			if s > 7 {
				s--
			}
			n, _ := strconv.Atoi(ctl[1])
			step[s].csw = n - 1
		default:
			fmt.Printf("Invalid master programmer switch %s\n", ctl[0])
		}
	}
}

func incstep(st int) {
	if step[st].kludge {
		return
	}
	if step[st].stage >= step[st].csw {
		step[st].stage = 0
	} else {
		step[st].stage++
	}
	step[st].kludge = true
}

func cycgate(st int) bool {
	stage := step[st].stage
	switch st {
	case 0:
		if assoc[0] == 1 {
			return false
		} else {
			return dec[0].val == dec[0].lim[stage]
		}
	case 1:
		if assoc[1] == 0 && dec[2].val != dec[2].lim[stage] {
			return false
		} else if dec[1].val != dec[1].lim[stage] {
			return false
		} else if assoc[0] == 1 && dec[0].val != dec[0].lim[stage] {
			return false
		}
		return true
	case 2:
		if assoc[2] == 0 && dec[6].val != dec[6].lim[stage] {
			return false
		} else if dec[5].val != dec[5].lim[stage] {
			return false
		} else if dec[4].val != dec[4].lim[stage] {
			return false
		} else if dec[3].val != dec[3].lim[stage] {
			return false
		} else if assoc[1] == 1 && dec[2].val != dec[2].lim[stage] {
			return false
		}
		return true
	case 3:
		if assoc[3] == 0 && dec[8].val != dec[8].lim[stage] {
			return false
		} else if dec[7].val != dec[7].lim[stage] {
			return false
		} else if assoc[2] == 1 && dec[6].val != dec[6].lim[stage] {
			return false
		}
		return true
	case 4:
		if assoc[3] == 0 {
			return dec[9].val == dec[9].lim[stage]
		} else {
			return dec[9].val == dec[9].lim[stage] &&
				dec[8].val == dec[8].lim[stage]
		}
	case 5:
		if assoc[4] == 1 {
			return false
		} else {
			return dec[10].val == dec[10].lim[stage]
		}
	case 6:
		if assoc[5] == 0 && dec[12].val != dec[12].lim[stage] {
			return false
		} else if dec[11].val != dec[11].lim[stage] {
			return false
		} else if assoc[4] == 1 && dec[10].val != dec[10].lim[stage] {
			return false
		}
		return true
	case 7:
		if assoc[6] == 0 && dec[16].val != dec[16].lim[stage] {
			return false
		} else if dec[15].val != dec[15].lim[stage] {
			return false
		} else if dec[14].val != dec[14].lim[stage] {
			return false
		} else if dec[13].val != dec[13].lim[stage] {
			return false
		} else if assoc[5] == 1 && dec[12].val != dec[12].lim[stage] {
			return false
		}
		return true
	case 8:
		if assoc[7] == 0 && dec[18].val != dec[18].lim[stage] {
			return false
		} else if dec[17].val != dec[17].lim[stage] {
			return false
		} else if assoc[6] == 1 && dec[16].val != dec[16].lim[stage] {
			return false
		}
		return true
	case 9:
		if assoc[7] == 0 {
			return dec[19].val == dec[19].lim[stage]
		} else {
			return dec[19].val == dec[19].lim[stage] &&
				dec[18].val == dec[18].lim[stage]
		}
	}
	return false
}

func incdec(de int) {
	dec[de].val++
	if dec[de].val == 10 {
		dec[de].val = 0
		dec[de].carry = true
	}
}

func clrdecset(st int) {
	switch st {
	case 0:
		if assoc[0] == 0 {
			dec[0].val = 0
			dec[0].carry = false
		}
	case 1:
		if assoc[0] == 1 {
			dec[0].val = 0
			dec[0].carry = false
		}
		dec[1].val = 0
		dec[1].carry = false
		if assoc[1] == 0 {
			dec[2].val = 0
			dec[2].carry = false
		}
	case 2:
		if assoc[1] == 1 {
			dec[2].val = 0
			dec[2].carry = false
		}
		dec[3].val = 0
		dec[3].carry = false
		dec[4].val = 0
		dec[4].carry = false
		dec[5].val = 0
		dec[5].carry = false
		if assoc[2] == 0 {
			dec[6].val = 0
			dec[6].carry = false
		}
	case 3:
		if assoc[2] == 1 {
			dec[6].val = 0
			dec[6].carry = false
		}
		dec[7].val = 0
		dec[7].carry = false
		if assoc[3] == 0 {
			dec[8].val = 0
			dec[8].carry = false
		}
	case 4:
		if assoc[3] == 1 {
			dec[8].val = 0
			dec[8].carry = false
		}
		dec[9].val = 0
		dec[9].carry = false
	case 5:
		if assoc[4] == 0 {
			dec[10].val = 0
			dec[10].carry = false
		}
	case 6:
		if assoc[4] == 1 {
			dec[10].val = 0
			dec[10].carry = false
		}
		dec[11].val = 0
		dec[11].carry = false
		if assoc[5] == 0 {
			dec[12].val = 0
			dec[12].carry = false
		}
	case 7:
		if assoc[5] == 1 {
			dec[12].val = 0
			dec[12].carry = false
		}
		dec[13].val = 0
		dec[13].carry = false
		dec[14].val = 0
		dec[14].carry = false
		dec[15].val = 0
		dec[15].carry = false
		if assoc[6] == 0 {
			dec[16].val = 0
			dec[16].carry = false
		}
	case 8:
		if assoc[6] == 1 {
			dec[16].val = 0
			dec[16].carry = false
		}
		dec[17].val = 0
		dec[17].carry = false
		if assoc[7] == 0 {
			dec[18].val = 0
			dec[18].carry = false
		}
	case 9:
		if assoc[7] == 1 {
			dec[18].val = 0
			dec[18].carry = false
		}
		dec[19].val = 0
		dec[19].carry = false
	}
}

func incdecset(st int) {
	switch st {
	case 0:
		if assoc[0] == 0 {
			incdec(0)
		}
	case 1:
		if assoc[1] == 0 {
			incdec(2)
			if dec[2].carry {
				incdec(1)
				dec[2].carry = false
			}
		} else {
			incdec(1)
		}
		if assoc[0] == 1 && dec[1].carry {
			incdec(0)
			dec[1].carry = false
		}
	case 2:
		if assoc[2] == 0 {
			incdec(6)
			if dec[6].carry {
				incdec(5)
				dec[6].carry = false
			}
		} else {
			incdec(5)
		}
		if dec[5].carry {
			incdec(4)
			dec[5].carry = false
		}
		if dec[4].carry {
			incdec(3)
			dec[4].carry = false
		}
		if assoc[1] == 1 && dec[3].carry {
			incdec(2)
			dec[3].carry = false
		}
	case 3:
		if assoc[3] == 0 {
			incdec(8)
			if dec[8].carry {
				incdec(7)
				dec[8].carry = false
			}
		} else {
			incdec(7)
		}
		if assoc[2] == 1 && dec[8].carry {
			incdec(7)
			dec[8].carry = false
		}
	case 4:
		incdec(9)
		if assoc[3] == 1 && dec[9].carry {
			incdec(8)
			dec[9].carry = false
		}
	case 5:
		if assoc[4] == 0 {
			incdec(10)
		}
	case 6:
		if assoc[5] == 0 {
			incdec(12)
			if dec[12].carry {
				incdec(11)
				dec[12].carry = false
			}
		} else {
			incdec(11)
		}
		if assoc[4] == 1 && dec[11].carry {
			incdec(10)
			dec[11].carry = false
		}
	case 7:
		if assoc[6] == 0 {
			incdec(16)
			if dec[16].carry {
				incdec(15)
				dec[16].carry = false
			}
		} else {
			incdec(15)
		}
		if dec[15].carry {
			incdec(14)
			dec[15].carry = false
		}
		if dec[14].carry {
			incdec(13)
			dec[14].carry = false
		}
		if assoc[5] == 1 && dec[13].carry {
			incdec(12)
			dec[13].carry = false
		}
	case 8:
		if assoc[7] == 0 {
			incdec(19)
			if dec[19].carry {
				incdec(18)
				dec[19].carry = false
			}
		} else {
			incdec(18)
		}
		if assoc[6] == 1 && dec[18].carry {
			incdec(17)
			dec[18].carry = false
		}
	case 9:
		incdec(19)
		if assoc[7] == 1 && dec[19].carry {
			incdec(18)
			dec[19].carry = false
		}
	}
}

func mpunit(cyctrunk chan pulse) {
	var p pulse

	mpupdate = make(chan int)
	go mpunit2()
	resp := make(chan int)
	for {
		p = <-cyctrunk
		cyc := p.val
		if cyc&Cpp != 0 {
			for i, s := range step {
				if cycgate(i) {
					clrdecset(i)
					incstep(i)
				}
				// Unclear what this should be: probably > 3 and < 12
				if s.inff >= 6 {
					incdecset(i)
					step[i].inff = 0
					if s.o[s.stage] != nil {
						s.o[s.stage] <- pulse{1, resp}
						<-resp
					}
				}
			}
		} else if cyc&Tenp != 0 {
			for i := 0; i < len(step); i++ {
				step[i].kludge = false
			}
		}
		if p.resp != nil {
			p.resp <- 1
		}
		// Simulate "flip-flop...time constant approximately equal to that
		// of the slow buffer output of a transceiver."  Huskey TM II, Ch X
		for i, s := range step {
			if s.inff > 0 {
				step[i].inff++
			}
		}
	}
}

func mpunit2() {
	var p pulse

	for {
		select {
		case <-mpupdate:
		case p = <-dec[0].di:
			incdec(0)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[1].di:
			incdec(1)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[2].di:
			incdec(2)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[3].di:
			incdec(3)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[4].di:
			incdec(4)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[5].di:
			incdec(5)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[6].di:
			incdec(6)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[7].di:
			incdec(7)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[8].di:
			incdec(8)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[9].di:
			incdec(9)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[10].di:
			incdec(10)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[11].di:
			incdec(11)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[12].di:
			incdec(12)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[13].di:
			incdec(13)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[14].di:
			incdec(14)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[15].di:
			incdec(15)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[16].di:
			incdec(16)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[17].di:
			incdec(17)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[18].di:
			incdec(18)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-dec[19].di:
			incdec(19)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[0].di:
			incstep(0)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[0].i:
			step[0].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[0].cdi:
			step[0].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[1].di:
			incstep(1)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[1].i:
			step[1].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[1].cdi:
			step[1].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[2].di:
			incstep(2)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[2].i:
			step[2].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[2].cdi:
			step[2].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[3].di:
			incstep(3)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[3].i:
			step[3].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[3].cdi:
			step[3].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[4].di:
			incstep(4)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[4].i:
			step[4].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[4].cdi:
			step[4].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[5].di:
			incstep(5)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[5].i:
			step[5].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[5].cdi:
			step[5].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[6].di:
			incstep(6)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[6].i:
			step[6].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[6].cdi:
			step[6].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[7].di:
			incstep(7)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[7].i:
			step[7].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[7].cdi:
			step[7].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[8].di:
			incstep(8)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[8].i:
			step[8].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[8].cdi:
			step[8].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[9].di:
			incstep(9)
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[9].i:
			step[9].inff = 1
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-step[9].cdi:
			step[9].stage = 0
			if p.resp != nil {
				p.resp <- 1
			}
		}
	}
}
