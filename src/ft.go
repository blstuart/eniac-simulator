package main

import (
	"fmt"
	"strconv"
)

type ft struct {
	jack             [27]chan pulse
	inff1, inff2     [11]bool
	opsw             [11]int
	rptsw            [11]int
	argsw            [11]int
	pm1, pm2         int
	cons             [8]int
	del              [8]int
	sub              [12]int
	tab              [104][14]int
	arg              int
	ring             int
	add              bool
	subtr            bool
	argsetup         bool
	gateh42, gatee42 bool
	update           chan int
	resp             chan int
	whichrp          bool
	px4119           bool
}

var ftable [3]ft

func ftstat(unit int) (s string) {
	for i := 0; i < 11; i++ {
		if ftable[unit].inff2[i] {
			s += "1"
		} else {
			s += "0"
		}
	}
	s += fmt.Sprintf(" %d %d", ftable[unit].arg, ftable[unit].ring)
	if ftable[unit].add {
		s += " 1"
	} else {
		s += " 0"
	}
	if ftable[unit].subtr {
		s += " 1"
	} else {
		s += " 0"
	}
	if ftable[unit].argsetup {
		s += " 1"
	} else {
		s += " 0"
	}
	return
}

func ftreset(unit int) {
	f := &ftable[unit]
	for i := 0; i < 27; i++ {
		f.jack[i] = nil
	}
	for i := 0; i < 11; i++ {
		f.inff1[i] = false
		f.inff2[i] = false
		f.opsw[i] = 0
		f.rptsw[i] = 0
		f.argsw[i] = 0
	}
	f.pm1 = 0
	f.pm2 = 0
	for i := 0; i < 8; i++ {
		f.cons[i] = 0
		f.del[i] = 0
	}
	for i := 0; i < 12; i++ {
		f.sub[i] = 0
	}
	for i := 0; i < 104; i++ {
		for j := 0; j < 14; j++ {
			f.tab[i][j] = 0
		}
	}
	f.arg = 0
	f.ring = 0
	f.add = false
	f.subtr = false
	f.argsetup = false
	f.gateh42 = false
	f.gatee42 = false
	f.whichrp = false
	f.px4119 = false
	f.update <- 1
}

func ftplug(unit int, jack string, ch chan pulse) {
	jacks := [22]string{"1i", "1o", "2i", "2o", "3i", "3o", "4i", "4o",
		"5i", "5o", "6i", "6o", "7i", "7o", "8i", "8o", "9i", "9o",
		"10i", "10o", "11i", "11o"}
	switch {
	case jack == "arg", jack == "ARG":
		ftable[unit].jack[0] = ch
	case jack == "A":
		ftable[unit].jack[1] = ch
	case jack == "B":
		ftable[unit].jack[2] = ch
	case jack == "NC":
		ftable[unit].jack[3] = ch
	case jack == "C":
		ftable[unit].jack[4] = ch
	default:
		for i, j := range jacks {
			if j == jack {
				ftable[unit].jack[i+5] = ch
				break
			}
		}
	}
	ftable[unit].update <- 1
}

func ftctl(unit int, ch chan [2]string) {
	var digit, row, val int
	var bank, ilk rune

	ops := [10]string{"A-2", "A-1", "A0", "A+1", "A+2", "S+2", "S+1", "S0", "S-1", "S-2"}
	for {
		swval := <-ch
		switch {
		case swval[0][:2] == "op":
			sw, _ := strconv.Atoi(swval[0][2:])
			for i, o := range ops {
				if o == swval[1] {
					ftable[unit].opsw[sw-1] = i
					break
				}
			}
		case swval[0][:2] == "cl":
			sw, _ := strconv.Atoi(swval[0][2:])
			switch swval[1] {
			case "0":
				ftable[unit].argsw[sw-1] = 0
			case "NC", "nc":
				ftable[unit].argsw[sw-1] = 1
			case "C", "c":
				ftable[unit].argsw[sw-1] = 2
			}
		case swval[0][:2] == "rp":
			sw, _ := strconv.Atoi(swval[0][2:])
			val, _ := strconv.Atoi(swval[1])
			ftable[unit].rptsw[sw-1] = val - 1
		case swval[0] == "mpm1":
			switch swval[1][0] {
			case 'P', 'p':
				ftable[unit].pm1 = 0
			case 'M', 'm':
				ftable[unit].pm1 = 1
			case 'T', 't':
				ftable[unit].pm1 = 2
			}
		case swval[0] == "mpm2":
			switch swval[1][0] {
			case 'P', 'p':
				ftable[unit].pm2 = 0
			case 'M', 'm':
				ftable[unit].pm2 = 1
			case 'T', 't':
				ftable[unit].pm2 = 2
			}
		case swval[0][0] == 'A', swval[0][0] == 'B':
			offset := 0
			fmt.Sscanf(swval[0], "%c%d%c", &bank, &digit, &ilk)
			if bank == 'B' {
				offset = 1
			}
			switch ilk {
			case 'd', 'D':
				switch swval[1] {
				case "d", "D":
					ftable[unit].del[4*offset+digit-1] = 1
				case "o", "O":
					ftable[unit].del[4*offset+digit-1] = 0
				}
			case 'c', 'C':
				switch swval[1] {
				case "pm1", "PM1":
					ftable[unit].cons[4*offset+digit-1] = 10
				case "pm2", "PM2":
					ftable[unit].cons[4*offset+digit-1] = 11
				default:
					n, _ := strconv.Atoi(swval[1])
					ftable[unit].cons[4*offset+digit-1] = n
				}
			case 's', 'S':
				switch swval[1] {
				case "s", "S":
					ftable[unit].sub[6*offset+digit-5] = 1
				case "0":
					ftable[unit].sub[6*offset+digit-5] = 0
				}
			}
		case swval[0][0] == 'R':
			n, _ := fmt.Sscanf(swval[0], "R%c%dL%d", &bank, &row, &digit)
			if n == 3 {
				val, _ = strconv.Atoi(swval[1])
				if bank == 'A' {
					ftable[unit].tab[row+2][7-digit] = val
				} else {
					ftable[unit].tab[row+2][13-digit] = val
				}
			} else {
				fmt.Sscanf(swval[0], "R%c%dS", &bank, &row)
				if swval[1] == "m" || swval[1] == "M" {
					val = 1
				} else {
					val = 0
				}
				if bank == 'A' {
					ftable[unit].tab[row+2][0] = val
				} else {
					ftable[unit].tab[row+2][13] = val
				}
			}
		case swval[0] == "ninep" || swval[0] == "Ninep":
			if swval[1][0] == 'C' || swval[1][0] == 'c' {
				ftable[unit].px4119 = true
			} else {
				ftable[unit].px4119 = false
			}
		default:
			fmt.Printf("Invalid function table switch %s\n", swval[0])
		}
	}
}

func addlookup(f *ft, c int) {
	a := 0
	b := 0
	arg := f.arg
	if c&Ninep != 0 {
		as := f.pm1 == 1 || f.pm1 == 2 && f.tab[arg][0] == 1
		bs := f.pm2 == 1 || f.pm2 == 2 && f.tab[arg][13] == 1
		if as {
			a |= 1 << 10
		}
		if bs {
			b |= 1 << 10
		}
		for i := 0; i < 4; i++ {
			if f.del[i] == 0 {
				if f.cons[i] == 10 && as {
					a |= 1 << (9 - uint(i))
				} else if f.cons[i] == 11 && bs {
					a |= 1 << (9 - uint(i))
				}
			}
			if f.del[i+4] == 0 {
				if f.cons[i+4] == 10 && as {
					b |= 1 << (9 - uint(i))
				} else if f.cons[i+4] == 11 && bs {
					b |= 1 << (9 - uint(i))
				}
			}
		}
		for i := 0; i < 4; i++ {
			if f.cons[i] == 9 {
				a |= 1 << (9 - uint(i))
			}
			if f.cons[i+4] == 9 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if f.tab[arg][i+1] == 9 {
				a |= 1 << (5 - uint(i))
			}
			if f.tab[arg][i+7] == 9 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Fourp != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x >= 4 && x <= 8 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x >= 4 && x <= 8 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x >= 4 && x <= 8 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x >= 4 && x <= 8 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Twopp != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x == 2 || x == 3 || x == 8 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x == 2 || x == 3 || x == 8 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x == 2 || x == 3 || x == 8 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x == 2 || x == 3 || x == 8 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Twop != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x > 5 && x < 9 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x > 5 && x < 9 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x > 5 && x < 9 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x > 5 && x < 9 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Onep != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x < 9 && x%2 == 1 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x < 9 && x%2 == 1 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x < 9 && x%2 == 1 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x < 9 && x%2 == 1 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if a != 0 && f.jack[1] != nil {
		f.jack[1] <- pulse{a, f.resp}
		<-f.resp
	}
	if b != 0 && f.jack[2] != nil {
		f.jack[2] <- pulse{b, f.resp}
		<-f.resp
	}
}

func subtrlookup(f *ft, c int) {
	a := 0
	b := 0
	arg := f.arg
	if c&Ninep != 0 {
		as := f.pm1 == 0 || f.pm1 == 2 && f.tab[arg][0] == 0
		bs := f.pm2 == 0 || f.pm2 == 2 && f.tab[arg][13] == 0
		if as {
			a |= 1 << 10
		}
		if bs {
			b |= 1 << 10
		}
		for i := 0; i < 4; i++ {
			if f.del[i] == 0 {
				if f.cons[i] == 10 && as {
					a |= 1 << (9 - uint(i))
				} else if f.cons[i] == 11 && bs {
					a |= 1 << (9 - uint(i))
				}
			}
			if f.del[i+4] == 0 {
				if f.cons[i+4] == 10 && as {
					b |= 1 << (9 - uint(i))
				} else if f.cons[i+4] == 11 && bs {
					b |= 1 << (9 - uint(i))
				}
			}
		}
	}
	if c&Fourp != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x < 6 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x < 6 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x < 6 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x < 6 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Twopp != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x < 2 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x < 2 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x < 2 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x < 2 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Twop != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x == 6 || x == 7 || x < 4 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x == 6 || x == 7 || x < 4 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x == 6 || x == 7 || x < 4 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x == 6 || x == 7 || x < 4 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Onep != 0 {
		for i := 0; i < 4; i++ {
			if x := f.cons[i]; x < 10 && x%2 == 0 {
				a |= 1 << (9 - uint(i))
			}
			if x := f.cons[i+4]; x < 10 && x%2 == 0 {
				b |= 1 << (9 - uint(i))
			}
		}
		for i := 0; i < 6; i++ {
			if x := f.tab[arg][i+1]; x < 10 && x%2 == 0 {
				a |= 1 << (5 - uint(i))
			}
			if x := f.tab[arg][i+7]; x < 10 && x%2 == 0 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if c&Onepp != 0 {
		for i := 0; i < 6; i++ {
			if f.sub[i] == 1 {
				a |= 1 << (5 - uint(i))
			}
			if f.sub[i+6] == 1 {
				b |= 1 << (5 - uint(i))
			}
		}
	}
	if a != 0 && f.jack[1] != nil {
		f.jack[1] <- pulse{a, f.resp}
		<-f.resp
	}
	if b != 0 && f.jack[2] != nil {
		f.jack[2] <- pulse{b, f.resp}
		<-f.resp
	}
}

func ftunit(unit int, cyctrunk chan pulse) {
	var prog int

	f := &ftable[unit]
	f.update = make(chan int)
	go ftunit2(f)
	f.resp = make(chan int)
	for {
		p := <-cyctrunk
		if f.px4119 {
			if p.val&Cpp != 0 {
				p.val |= Ninep
			} else {
				p.val &= ^Ninep
			}
		}
		c := p.val
		if f.gatee42 {
			sw := f.opsw[prog]
			if c&Onep != 0 && (sw == 1 || sw == 3 || sw == 6 || sw == 8) {
				f.arg++
			}
			if c&Twop != 0 && (sw == 2 || sw == 3 || sw == 6 || sw == 7) {
				f.arg++
			}
			if c&Fourp != 0 && (sw == 4 || sw == 5) {
				f.arg++
			}
		}
		if f.add {
			if f.arg >= 0 && f.arg < 104 {
				addlookup(f, c)
			} else {
				fmt.Println("Invalid function table argument", f.arg)
			}
		}
		if f.subtr {
			if f.arg >= 0 && f.arg < 104 {
				subtrlookup(f, c)
			} else {
				fmt.Println("Invalid function table argument", f.arg)
			}
		}
		if c&Cpp != 0 {
			switch f.ring {
			case 0: // Stage -3
				for prog = 0; prog < 11 && !f.inff2[prog]; prog++ {
				}
				if prog >= 11 {
					break
				}
				switch f.argsw[prog] {
				case 1:
					if f.jack[3] != nil {
						f.jack[3] <- pulse{1, f.resp}
						<-f.resp
					}
				case 2:
					if f.jack[4] != nil {
						f.jack[4] <- pulse{1, f.resp}
						<-f.resp
					}
				}
				f.ring++ // Stage -2 begins
				f.gateh42 = true
			case 1:
				f.ring++ // Stage -1 begins
				f.gateh42 = false
				f.gatee42 = true
			case 2:
				f.ring++ // Stage 0 begins
				f.gatee42 = false
				/*
					if f.opsw[prog] < 5 {
						f.add = true
					} else {
						f.subtr = true
					}
				*/
			case 3: // Stage 0
				f.ring++ // Stage 1 begins
				if f.opsw[prog] < 5 {
					f.add = true
				} else {
					f.subtr = true
				}
			default: // Stages 1-9
				if f.rptsw[prog] == f.ring-4 {
					if f.jack[prog*2+6] != nil {
						f.jack[prog*2+6] <- pulse{1, f.resp}
						<-f.resp
					}
					f.arg = 0
					f.add = false
					f.subtr = false
					f.inff2[prog] = false
					f.argsetup = false
					f.ring = 0
				} else {
					f.ring++
				}
			}
		}
		if c&Ccg != 0 {
			f.whichrp = false
		}
		if c&Rp != 0 {
			if f.whichrp {
				for i, _ := range f.inff1 {
					if f.inff1[i] {
						f.inff1[i] = false
						f.inff2[i] = true
					}
				}
				f.whichrp = false
			} else {
				f.whichrp = true
			}
		}
		if f.ring == 2 && c&Onepp != 0 {
			f.argsetup = true
		}
		if p.resp != nil {
			p.resp <- 1
		}
	}
}

func ftunit2(f *ft) {
	var p pulse

	for {
		select {
		case <-f.update:
		case p = <-f.jack[5]:
			f.inff1[0] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[7]:
			f.inff1[1] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[9]:
			f.inff1[2] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[11]:
			f.inff1[3] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[13]:
			f.inff1[4] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[15]:
			f.inff1[5] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[17]:
			f.inff1[6] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[19]:
			f.inff1[7] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[21]:
			f.inff1[8] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[23]:
			f.inff1[9] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p = <-f.jack[25]:
			f.inff1[10] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case arg := <-f.jack[0]:
			if f.gateh42 {
				if arg.val&0x01 != 0 {
					f.arg++
				}
				if arg.val&0x02 != 0 {
					f.arg += 10
				}
			}
			if arg.resp != nil {
				arg.resp <- 1
			}
		}
	}
}
