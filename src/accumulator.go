package main

import (
	"fmt"
	"strconv"
)

/*
 * Bit positions for the ST1 and ST2 connectors
 */
const (
	pin1 = 1 << iota
	pin2
	pin3
	pin4
	pin5
	pin6
	pin7
	pin8
	pin9
	pin10
	_
	_
	_
	pin14
	pin15
	pin16
	pin17
)

/*
 * Signal names on ST1 and ST2 connectors
 */
const (
	stα        = pin1
	stβ        = pin2
	stγ        = pin3
	stδ        = pin4
	stε        = pin5
	stA        = pin6
	stAS       = pin7
	stS        = pin8
	stCLR      = pin9
	stCORR     = pin10
	stPMinp    = pin14
	stCORRsrc  = pin14
	stDEC10CAR = pin15
	stDEC1inp  = pin15
	stSF0out   = pin16
	stDEC1sub  = pin16
	stSFSWinp  = pin17
	stSF10out  = pin17
)

type accumulator struct {
	α, β, γ, δ, ε, A, S chan pulse
	ctlterm             [20]chan pulse
	inff1, inff2        [12]bool
	opsw                [12]byte
	clrsw               [12]bool
	rptsw               [8]byte
	sigfig              int
	sc                  byte
	val                 [10]byte
	decff               [10]bool
	sign                bool
	h50                 bool
	rep                 int
	whichrp             bool
	change              chan int
	lbuddy, rbuddy      int
}

var units [20]accumulator

func accstat(unit int) string {
	var s string

	if units[unit].sign {
		s = "M "
	} else {
		s = "P "
	}
	for i := 9; i >= 0; i-- {
		s += fmt.Sprintf("%d", units[unit].val[i])
	}
	s += " "
	for i := 9; i >= 0; i-- {
		s += b2is(units[unit].decff[i])
	}
	s += fmt.Sprintf(" %d ", units[unit].rep)
	for _, f := range units[unit].inff1 {
		s += b2is(f)
	}
	return s
}

func accreset(unit int) {
	u := &units[unit]
	u.α = nil
	u.β = nil
	u.γ = nil
	u.δ = nil
	u.ε = nil
	for i := 0; i < 12; i++ {
		u.ctlterm[i] = nil
		u.inff1[i] = false
		u.inff2[i] = false
		u.opsw[i] = 0
		u.clrsw[i] = false
	}
	for i := 0; i < 8; i++ {
		u.rptsw[i] = 0
	}
	u.sigfig = 10
	u.sc = 0
	u.h50 = false
	u.rep = 0
	u.whichrp = false
	u.lbuddy = unit
	u.rbuddy = unit
	accclear(unit)
	u.change <- 1
}

func accclear(acc int) {
	for i := 0; i < 10; i++ {
		units[acc].val[i] = 0
		units[acc].decff[i] = false
	}
	if units[acc].sigfig < 10 {
		units[acc].val[9-units[acc].sigfig] = 5
	}
	units[acc].sign = false
}

func accinterconnect(p1 []string, p2 []string) {
	unit1, _ := strconv.Atoi(p1[0][1:])
	unit2 := -1
	if len(p2) > 1 && p2[0][0] == 'a' {
		unit2, _ = strconv.Atoi(p2[0][1:])
	}
	switch {
	case p2[0] == "m" && p2[1] == "l":
		units[unit1-1].lbuddy = -1
	case p2[0] == "m" && p2[1] == "r":
		units[unit1-1].lbuddy = -2
	case p2[0] == "d" && p2[1] == "sv":
		units[unit1-1].lbuddy = -3
	case p2[0] == "d" && p2[1] == "su2q":
		units[unit1-1].lbuddy = -4
	case p2[0] == "d" && p2[1] == "su2s":
		units[unit1-1].lbuddy = -5
	case p2[0] == "d" && p2[1] == "su3":
		units[unit1-1].lbuddy = -6
	case p1[1] == "st1" || p1[1] == "il1":
		if unit2 != -1 && unit1 != unit2 {
			units[unit1-1].lbuddy = unit2 - 1
			units[unit2-1].rbuddy = unit1 - 1
		}
	case p1[1] == "st2" || p1[1] == "ir1":
		if unit2 != -1 && unit1 != unit2 {
			units[unit1-1].rbuddy = unit2 - 1
			units[unit2-1].lbuddy = unit1 - 1
		}
	case p1[1] == "su1" || p1[1] == "il2":
		if unit2 != -1 && unit1 != unit2 {
			units[unit1-1].lbuddy = unit2 - 1
			units[unit2].rbuddy = unit1 - 1
		}
	case p1[1] == "su2" || p1[1] == "ir2":
		if unit2 != -1 && unit1 != unit2 {
			units[unit1-1].rbuddy = unit2 - 1
			units[unit2-1].lbuddy = unit1 - 1
		}
	}
	if unit2 != -1 && unit1 != unit2 {
		units[unit1-1].change <- 1
		units[unit2-1].change <- 1
	}
}

func accplug(unit int, jack string, ch chan pulse) {
	jacks := [20]string{"1i", "2i", "3i", "4i", "5i", "5o", "6i", "6o", "7i", "7o",
		"8i", "8o", "9i", "9o", "10i", "10o", "11i", "11o", "12i", "12o"}

	switch {
	case jack == "α", jack == "a", jack == "alpha":
		units[unit].α = ch
	case jack == "β", jack == "b", jack == "beta":
		units[unit].β = ch
	case jack == "γ", jack == "g", jack == "gamma":
		units[unit].γ = ch
	case jack == "δ", jack == "d", jack == "delta":
		units[unit].δ = ch
	case jack == "ε", jack == "e", jack == "epsilon":
		units[unit].ε = ch
	case jack == "A":
		units[unit].A = ch
	case jack == "S":
		units[unit].S = ch
	case jack[0] == 'I':
	default:
		foundjack := false
		for i, j := range jacks {
			if j == jack {
				units[unit].ctlterm[i] = ch
				foundjack = true
				break
			}
		}
		if !foundjack {
			fmt.Println("Invalid jack:", jack, "on accumulator", unit+1)
		}
	}
	units[unit].change <- 1
}

func accctl(unit int, ch chan [2]string) {
	units[unit].lbuddy = unit
	units[unit].rbuddy = unit
	for {
		ctl := <-ch
		prog, _ := strconv.Atoi(ctl[0][2:])
		prog--
		switch ctl[0][:2] {
		case "op":
			switch ctl[1] {
			case "α", "a", "alpha":
				units[unit].opsw[prog] = 0
			case "β", "b", "beta":
				units[unit].opsw[prog] = 1
			case "γ", "g", "gamma":
				units[unit].opsw[prog] = 2
			case "δ", "d", "delta":
				units[unit].opsw[prog] = 3
			case "ε", "e", "epsilon":
				units[unit].opsw[prog] = 4
			case "0":
				units[unit].opsw[prog] = 5
			case "A":
				units[unit].opsw[prog] = 6
			case "AS":
				units[unit].opsw[prog] = 7
			case "S":
				units[unit].opsw[prog] = 8
			default:
				fmt.Println("Invalid operation code:", ctl[1], "on unit",
					unit+1, "program", prog+1)
			}
		case "cc":
			switch ctl[1] {
			case "0":
				units[unit].clrsw[prog] = false
			case "C", "c":
				units[unit].clrsw[prog] = true
			default:
				fmt.Println("Invalid clear/correct setting:", ctl[1], "on unit",
					unit+1, "program", prog+1)
			}
		case "rp":
			rpt, err := strconv.Atoi(ctl[1])
			if err == nil {
				units[unit].rptsw[prog-4] = byte(rpt - 1)
			} else {
				fmt.Println("Invalid repeat count:", ctl[1], "on unit",
					unit+1, "program", prog+1)
			}
		case "sf":
			n, _ := strconv.Atoi(ctl[1])
			units[unit].sigfig = n
		case "sc":
			switch ctl[1] {
			case "0":
				units[unit].sc = 0
			case "SC", "sc":
				units[unit].sc = 1
			default:
				fmt.Println("Invalid selective clear setting:", ctl[1],
					"on unit", unit+1)
			}
		}
	}
}

/*
 * Implement the PX-5-109 terminator and PX-5-110
 * and PX-5-121 interconnect cables
 */

func su1(unit int) int {
	u := &units[unit]
	x := 0
	if u.rbuddy >= 0 && u.rbuddy != unit {
		x = su1(u.rbuddy)
	}
	for i := 0; i < 12; i++ {
		if u.inff1[i] || u.inff2[i] {
			switch u.opsw[i] {
			case 0:
				x |= stα
			case 1:
				x |= stβ
			case 2:
				x |= stγ
			case 3:
				x |= stδ
			case 4:
				x |= stε
			case 6:
				x |= stA
			case 7:
				x |= stAS
			case 8:
				x |= stS
			}
			if u.clrsw[i] {
				if u.opsw[i] >= 5 {
					if i < 4 || u.rep == int(u.rptsw[i-4]) {
						x |= stCLR
					}
				} else {
					x |= stCORR
				}
			}
		}
	}
	return x
}

func st1(unit int) int {
	u := &units[unit]
	x := 0
	if u.lbuddy == unit {
		x = su1(unit)
	} else if u.lbuddy == -1 {
		x = su1(unit)
		if Multl {
			x |= stα
		}
	} else if u.lbuddy == -2 {
		x = su1(unit)
		if Multr {
			x |= stα
		}
	} else if u.lbuddy == -3 {
		x = su1(unit)
		x |= divsr.sv
	} else if u.lbuddy == -4 {
		x = su1(unit)
		/* Wiring for PX-5-134 for quotient */
		x |= divsr.su2 & stα
		x |= (divsr.su2 & su2qA) << 2
		x |= (divsr.su2 & su2qS) << 3
		x |= (divsr.su2 & su2qCLR) << 3
	} else if u.lbuddy == -5 {
		x = su1(unit)
		/* Wiring for PX-5-135 for shift */
		x |= (divsr.su2 & su2sα) >> 1
		x |= (divsr.su2 & su2sA) << 3
		x |= (divsr.su2 & su2sCLR) << 2
	} else if u.lbuddy == -6 {
		x = su1(unit)
		/* Wiring for PX-5-136 for denominator */
		x |= divsr.su3 & (stα | stβ | stγ)
		x |= (divsr.su3 & su3A) << 2
		x |= (divsr.su3 & su3S) << 3
		x |= (divsr.su3 & su3CLR) << 3
	} else {
		x = st2(u.lbuddy) & 0x1c3ff
	}
	return x
}

func st2(unit int) int {
	//	u := &units[unit]
	x := st1(unit) & 0x03ff

	return x
}

func accrecv(unit, dat int) {
	u := &units[unit]
	for i := 0; i < 10; i++ {
		if dat&1 == 1 {
			u.val[i]++
			if u.val[i] >= 10 {
				u.decff[i] = true
				u.val[i] -= 10
			}
		}
		dat >>= 1
	}
	if dat&1 == 1 {
		u.sign = !u.sign
	}
}

func docpp(u *accumulator, resp chan int, cyc int) {
	for i := 0; i < 4; i++ {
		if u.inff2[i] {
			u.inff2[i] = false
		}
	}
	if u.h50 {
		u.rep++
		rstrep := false
		for i := 4; i < 12; i++ {
			if u.inff2[i] && u.rep == int(u.rptsw[i-4])+1 {
				u.inff2[i] = false
				rstrep = true
				t := (i-4)*2 + 5
				if u.ctlterm[t] != nil {
					u.ctlterm[t] <- pulse{1, resp}
					<-resp
				}
			}
		}
		if rstrep {
			u.rep = 0
			u.h50 = false
		}
	}
}

func ripple(unit int) {
	u := &units[unit]
	for i := 0; i < 9; i++ {
		if u.decff[i] {
			u.val[i+1]++
			if u.val[i+1] == 10 {
				u.val[i+1] = 0
				u.decff[i+1] = true
			}
		}
	}
	if u.lbuddy < 0 || u.lbuddy == unit {
		if u.decff[9] {
			/*
			 * Connection PX-5-121 pins 14, 15
			 */
			u.sign = !u.sign
		}
	} else {
		/*
		 * PX-5-110, pin 15 straight through
		 */
		if u.decff[9] {
			units[u.lbuddy].val[0]++
			if units[u.lbuddy].val[0] == 10 {
				units[u.lbuddy].val[0] = 0
				units[u.lbuddy].decff[0] = true
			}
		}
		ripple(u.lbuddy)
	}
}

func doccg(u *accumulator, unit int, resp chan int) {
	curprog := st1(unit)
	u.whichrp = false
	if curprog&0x1f != 0 {
		if u.rbuddy == unit {
			ripple(unit)
		}
	} else if (curprog & stCLR) != 0 {
		for i := 0; i < 10; i++ {
			u.val[i] = byte(0)
		}
		u.sign = false
	}
}

func dorp(u *accumulator) {
	if !u.whichrp {
		/*
		 * Ugly hack to avoid races.  Effectively this is
		 * a coarse approximation to the "slow buffer
		 * output" described in 1.2.9 of the Technical
		 * Manual Part 2.
		 */
		for i := 0; i < 12; i++ {
			if u.inff1[i] {
				u.inff2[i] = true
				u.inff1[i] = false
				if i >= 4 {
					u.h50 = true
				}
			}
		}
		for i := 0; i < 10; i++ {
			u.decff[i] = false
		}
		u.whichrp = true
	}
}

func dotenp(u *accumulator, unit int) {
	curprog := st1(unit)
	if curprog&(stA|stAS|stS) != 0 {
		for i := 0; i < 10; i++ {
			u.val[i]++
			if u.val[i] == 10 {
				u.val[i] = 0
				u.decff[i] = true
			}
		}
	}
}

func doninep(u *accumulator, unit int, resp chan int) {
	curprog := st1(unit)
	if curprog&(stA|stAS) != 0 {
		if u.A != nil {
			n := 0
			for i := 0; i < 10; i++ {
				if u.decff[i] {
					n |= 1 << uint(i)
				}
			}
			if u.sign {
				n |= 1 << 10
			}
			if n != 0 {
				u.A <- pulse{n, resp}
				<-resp
			}
		}
	}
	if curprog&(stAS|stS) != 0 {
		if u.S != nil {
			n := 0
			for i := 0; i < 10; i++ {
				if !u.decff[i] {
					n |= 1 << uint(i)
				}
			}
			if !u.sign {
				n |= 1 << 10
			}
			if n != 0 {
				u.S <- pulse{n, resp}
				<-resp
			}
		}
	}
}

func doonepp(u *accumulator, unit int, resp chan int) {
	curprog := st1(unit)
	if curprog&stCORR != 0 {
		/*
		 * Connection of PX-5-109 pins 14, 15
		 */
		if u.rbuddy == unit {
			u.val[0]++
			if u.val[0] > 9 {
				u.val[0] = 0
				u.decff[0] = true
			}
		}
	}
	if curprog&(stAS|stS) != 0 && u.S != nil {
		if ((u.lbuddy < 0 || u.lbuddy == unit) && u.rbuddy == unit && u.sigfig > 0) ||
			(u.rbuddy != unit && u.sigfig < 10) ||
			(u.lbuddy != unit && u.lbuddy >= 0 && units[u.lbuddy].sigfig == 10 && u.sigfig > 0) ||
			(u.rbuddy != unit && u.sigfig == 10 && units[u.rbuddy].sigfig == 0) {
			u.S <- pulse{1 << uint(10-u.sigfig), resp}
			<-resp
		}
	}
}

func accunit(unit int, cyctrunk chan pulse) {
	u := &units[unit]
	u.change = make(chan int)
	u.sigfig = 10
	u.lbuddy = unit
	u.rbuddy = unit
	go accunit2(unit)

	resp := make(chan int)
	for {
		p := <-cyctrunk
		cyc := p.val
		switch {
		case cyc&Cpp != 0:
			docpp(u, resp, cyc)
		case cyc&Ccg != 0:
			doccg(u, unit, resp)
		case cyc&Scg != 0:
			if u.sc == 1 {
				accclear(unit)
			}
		case cyc&Rp != 0:
			dorp(u)
		case cyc&Tenp != 0:
			dotenp(u, unit)
		case cyc&Ninep != 0:
			doninep(u, unit, resp)
		case cyc&Onepp != 0:
			doonepp(u, unit, resp)
		}
		if p.resp != nil {
			p.resp <- 1
		}
	}
}

func accunit2(unit int) {
	var dat, prog pulse

	u := &units[unit]
	for {
		select {
		case <-u.change:
		case dat = <-u.α:
			if st1(unit)&stα != 0 {
				accrecv(unit, dat.val)
			}
			if dat.resp != nil {
				dat.resp <- 1
			}
		case dat = <-u.β:
			if st1(unit)&stβ != 0 {
				accrecv(unit, dat.val)
			}
			if dat.resp != nil {
				dat.resp <- 1
			}
		case dat = <-u.γ:
			if st1(unit)&stγ != 0 {
				accrecv(unit, dat.val)
			}
			if dat.resp != nil {
				dat.resp <- 1
			}
		case dat = <-u.δ:
			if st1(unit)&stδ != 0 {
				accrecv(unit, dat.val)
			}
			if dat.resp != nil {
				dat.resp <- 1
			}
		case dat = <-u.ε:
			if st1(unit)&stε != 0 {
				accrecv(unit, dat.val)
			}
			if dat.resp != nil {
				dat.resp <- 1
			}
		case prog = <-u.ctlterm[0]:
			if prog.val == 1 {
				u.inff1[0] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[1]:
			if prog.val == 1 {
				u.inff1[1] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[2]:
			if prog.val == 1 {
				u.inff1[2] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[3]:
			if prog.val == 1 {
				u.inff1[3] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[4]:
			if prog.val == 1 {
				u.inff1[4] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[6]:
			if prog.val == 1 {
				u.inff1[5] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[8]:
			if prog.val == 1 {
				u.inff1[6] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[10]:
			if prog.val == 1 {
				u.inff1[7] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[12]:
			if prog.val == 1 {
				u.inff1[8] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[14]:
			if prog.val == 1 {
				u.inff1[9] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[16]:
			if prog.val == 1 {
				u.inff1[10] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-u.ctlterm[18]:
			if prog.val == 1 {
				u.inff1[11] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		}
	}
}
