package main

import (
	"fmt"
	"strconv"
)

type pulseset struct {
	one, two, twop, four int
}

var multin, multout [24]chan pulse
var R, D [5]chan pulse
var A, S, AS, AC, SC, ASC, RS, DS, F chan pulse
var lhppI, lhppII, rhppI, rhppII chan pulse
var stage int
var multff [24]bool
var iersw, iercl, icandsw, icandcl, sigsw, placsw, prodsw [24]int
var reset1ff, reset3ff bool
var Multl, Multr bool
var buffer61, f44 bool
var multupdate chan int

var table10 [10][10]pulseset = [10][10]pulseset{{},
	{},
	{{}, {}, {}, {}, {},
		{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}},
	{{}, {}, {}, {}, {1, 0, 0, 0},
		{1, 0, 0, 0}, {1, 0, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}},
	{{}, {}, {}, {1, 0, 0, 0}, {1, 0, 0, 0},
		{0, 1, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 0, 0}},
	{{}, {}, {1, 0, 0, 0}, {1, 0, 0, 0}, {0, 1, 0, 0},
		{0, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}},
	{{}, {}, {1, 0, 0, 0}, {1, 0, 0, 0}, {0, 1, 0, 0},
		{1, 1, 0, 0}, {1, 1, 0, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {1, 1, 1, 0}},
	{{}, {}, {1, 0, 0, 0}, {0, 1, 0, 0}, {0, 1, 0, 0},
		{1, 1, 0, 0}, {0, 0, 0, 1}, {0, 0, 0, 1}, {1, 0, 0, 1}, {0, 1, 0, 1}},
	{{}, {}, {1, 0, 0, 0}, {0, 1, 0, 0}, {1, 1, 0, 0},
		{0, 0, 0, 1}, {0, 0, 0, 1}, {1, 0, 0, 1}, {0, 1, 0, 1}, {1, 1, 0, 1}},
	{{}, {}, {1, 0, 0, 0}, {0, 1, 0, 0}, {1, 0, 1, 0},
		{0, 0, 0, 1}, {1, 1, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 1}, {0, 1, 1, 1}},
}

var table1 [10][10]pulseset = [10][10]pulseset{{},
	{{}, {1, 0, 0, 0}, {0, 1, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0},
		{1, 0, 0, 1}, {0, 1, 0, 1}, {1, 0, 1, 1}, {0, 1, 1, 1}, {1, 1, 1, 1}},
	{{}, {0, 1, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 1}, {0, 1, 1, 1},
		{}, {0, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 1, 1}, {0, 1, 1, 1}},
	{{}, {1, 1, 0, 0}, {0, 0, 1, 1}, {1, 1, 1, 1}, {0, 1, 0, 0},
		{1, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 0, 0}, {0, 0, 0, 1}, {1, 0, 1, 1}},
	{{}, {0, 1, 1, 0}, {0, 1, 1, 1}, {0, 1, 0, 0}, {0, 0, 1, 1},
		{}, {0, 1, 1, 0}, {0, 1, 1, 1}, {0, 1, 0, 0}, {0, 0, 1, 1}},
	{{}, {1, 0, 0, 1}, {}, {1, 0, 0, 1}, {},
		{1, 0, 0, 1}, {}, {1, 0, 0, 1}, {}, {1, 0, 0, 1}},
	{{}, {0, 0, 1, 1}, {0, 1, 0, 0}, {0, 1, 1, 1}, {0, 1, 1, 0},
		{}, {0, 0, 1, 1}, {0, 1, 0, 0}, {0, 1, 1, 1}, {0, 0, 0, 1}},
	{{}, {1, 0, 1, 1}, {0, 1, 1, 0}, {1, 0, 0, 0}, {0, 1, 1, 1},
		{1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 1, 1}, {0, 0, 1, 1}, {1, 1, 0, 0}},
	{{}, {0, 1, 1, 1}, {0, 0, 1, 1}, {0, 1, 1, 0}, {0, 1, 0, 0},
		{}, {0, 1, 1, 1}, {0, 0, 1, 1}, {0, 0, 0, 1}, {0, 1, 0, 0}},
	{{}, {1, 1, 1, 1}, {0, 1, 1, 1}, {1, 0, 1, 1}, {0, 1, 0, 1},
		{1, 0, 0, 1}, {0, 1, 1, 0}, {1, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}},
}

func multstat() string {
	s := fmt.Sprintf("%d ", stage)
	for i, _ := range multff {
		if multff[i] {
			s += "1"
		} else {
			s += "0"
		}
	}
	if reset1ff {
		s += " 1"
	} else {
		s += " 0"
	}
	if reset3ff {
		s += " 1"
	} else {
		s += " 0"
	}
	return s
}

func multreset() {
	for i := 0; i < 24; i++ {
		multin[i] = nil
		multout[i] = nil
		multff[i] = false
		iersw[i] = 0
		iercl[i] = 0
		icandsw[i] = 0
		icandcl[i] = 0
		sigsw[i] = 0
		placsw[i] = 0
		prodsw[i] = 0
	}
	for i := 0; i < 5; i++ {
		R[i] = nil
		D[i] = nil
	}
	A = nil
	S = nil
	AS = nil
	AC = nil
	SC = nil
	ASC = nil
	RS = nil
	DS = nil
	F = nil
	lhppI = nil
	lhppII = nil
	rhppI = nil
	rhppII = nil
	stage = 0
	reset1ff = false
	reset3ff = false
	Multl = false
	Multr = false
	buffer61 = false
	f44 = false
	multupdate <- 1
}

func multclear() {
}

func multplug(jack string, ch chan pulse) {
	switch {
	case jack == "Rα", jack == "Ra", jack == "rα", jack == "ra":
		R[0] = ch
	case jack == "Rβ", jack == "Rb", jack == "rβ", jack == "rb":
		R[1] = ch
	case jack == "Rγ", jack == "Rg", jack == "rγ", jack == "rg":
		R[2] = ch
	case jack == "Rδ", jack == "Rd", jack == "rδ", jack == "rd":
		R[3] = ch
	case jack == "Rε", jack == "Re", jack == "rε", jack == "re":
		R[4] = ch
	case jack == "Dα", jack == "Da", jack == "dα", jack == "da":
		D[0] = ch
	case jack == "Dβ", jack == "Db", jack == "dβ", jack == "db":
		D[1] = ch
	case jack == "Dγ", jack == "Dg", jack == "dγ", jack == "dg":
		D[2] = ch
	case jack == "Dδ", jack == "Dd", jack == "dδ", jack == "dd":
		D[3] = ch
	case jack == "Dε", jack == "De", jack == "dε", jack == "de":
		D[4] = ch
	case jack == "A", jack == "a":
		A = ch
	case jack == "S", jack == "s":
		S = ch
	case jack == "AS", jack == "as":
		AS = ch
	case jack == "AC", jack == "ac":
		AC = ch
	case jack == "SC", jack == "sc":
		SC = ch
	case jack == "ASC", jack == "asc":
		ASC = ch
	case jack == "RS", jack == "rs":
		RS = ch
	case jack == "DS", jack == "ds":
		DS = ch
	case jack == "F", jack == "f":
		F = ch
	case jack == "LHPPI", jack == "lhppi", jack == "lhppI":
		lhppI = ch
	case jack == "LHPPII", jack == "lhppii", jack == "lhppII":
		lhppII = ch
	case jack == "RHPPI", jack == "rhppi", jack == "rhppI":
		rhppI = ch
	case jack == "RHPPII", jack == "rhppii", jack == "rhppII":
		rhppII = ch
	default:
		prog, err := strconv.Atoi(jack[:len(jack)-1])
		if err != nil {
			fmt.Printf("Invalid multiplier jack %s\n", jack)
			return
		}
		switch jack[len(jack)-1] {
		case 'i':
			multin[prog-1] = ch
		case 'o':
			multout[prog-1] = ch
		}
	}
	multupdate <- 1
}

func recv2val(recv string) int {
	switch recv {
	case "α", "a", "alpha":
		return 0
	case "β", "b", "beta":
		return 1
	case "γ", "g", "gamma":
		return 2
	case "δ", "d", "delta":
		return 3
	case "ε", "e", "epsilon":
		return 4
	case "0":
		return 5
	}
	return 5
}

func multctl(ch chan [2]string) {
	products := [7]string{"A", "S", "AS", "0", "AC", "SC", "ASC"}
	for {
		ctl := <-ch
		switch {
		case len(ctl[0]) > 6 && ctl[0][:6] == "ieracc":
			prog, _ := strconv.Atoi(ctl[0][6:])
			iersw[prog-1] = recv2val(ctl[1])
		case len(ctl[0]) > 5 && ctl[0][:5] == "iercl":
			prog, _ := strconv.Atoi(ctl[0][5:])
			if ctl[1] == "C" {
				iercl[prog-1] = 1
			} else {
				iercl[prog-1] = 0
			}
		case len(ctl[0]) > 8 && ctl[0][:8] == "icandacc":
			prog, _ := strconv.Atoi(ctl[0][8:])
			icandsw[prog-1] = recv2val(ctl[1])
		case len(ctl[0]) > 7 && ctl[0][:7] == "icandcl":
			prog, _ := strconv.Atoi(ctl[0][7:])
			if ctl[1] == "C" {
				icandcl[prog-1] = 1
			} else {
				icandcl[prog-1] = 0
			}
		case len(ctl[0]) > 2 && ctl[0][:2] == "sf":
			prog, _ := strconv.Atoi(ctl[0][2:])
			val, _ := strconv.Atoi(ctl[1])
			if val == 0 {
				val = 1
			}
			sigsw[prog-1] = 10 - val
		case len(ctl[0]) > 5 && ctl[0][:5] == "place":
			prog, _ := strconv.Atoi(ctl[0][5:])
			val, _ := strconv.Atoi(ctl[1])
			placsw[prog-1] = val - 2
		case len(ctl[0]) > 4 && ctl[0][:4] == "prod":
			prog, _ := strconv.Atoi(ctl[0][4:])
			for i, p := range products {
				if p == ctl[1] {
					prodsw[prog-1] = i
					break
				}
			}
		default:
			fmt.Printf("Invalid multiplier switch %s\n", ctl[0])
		}
	}
}

func multargs(prog int) {
	resp1 := make(chan int)
	resp2 := make(chan int)
	ier := iersw[prog]
	icand := icandsw[prog]
	if ier < 5 && R[ier] != nil {
		R[ier] <- pulse{1, resp1}
	}
	if icand < 5 && D[icand] != nil {
		D[icand] <- pulse{1, resp2}
	}
	if ier < 5 && R[ier] != nil {
		<-resp1
	}
	if icand < 5 && D[icand] != nil {
		<-resp2
	}
	multff[prog] = true
	buffer61 = true
}

func shiftprod(lhpp, rhpp int, resp1, resp2, resp3, resp4 chan int) {
	if lhppI != nil && lhpp != 0 {
		lhppI <- pulse{lhpp >> uint(stage-2), resp1}
	}
	if lhppII != nil && lhpp != 0 {
		lhppII <- pulse{(lhpp << uint(12-stage)) & 0x3ff, resp2}
	}
	if rhppI != nil && rhpp != 0 {
		rhppI <- pulse{rhpp >> uint(stage-1), resp3}
	}
	if rhppII != nil && rhpp != 0 {
		rhppII <- pulse{(rhpp << uint(11-stage)) & 0x3ff, resp4}
	}
	if lhppI != nil && lhpp != 0 {
		<-resp1
	}
	if lhppII != nil && lhpp != 0 {
		<-resp2
	}
	if rhppI != nil && rhpp != 0 {
		<-resp3
	}
	if rhppII != nil && rhpp != 0 {
		<-resp4
	}
}

func multunit(cyctrunk chan pulse) {
	var ier, icand string
	var sigfig int

	multupdate = make(chan int)
	resp1 := make(chan int)
	resp2 := make(chan int)
	resp3 := make(chan int)
	resp4 := make(chan int)
	go multunit2()
	for {
		c := <-cyctrunk
		switch {
		case c.val&Cpp != 0:
			if f44 {
				stage = 1
				f44 = false
			} else if stage == 12 {
				reset1ff = true
				reset3ff = true
				handshake(1, F, resp1)
				stage++
			} else if stage == 13 {
				which := -1
				for i, f := range multff {
					if f {
						which = i
						break
					}
				}
				if which != -1 {
					handshake(1, multout[which], resp1)
					multff[which] = false
					switch prodsw[which] {
					case 0:
						handshake(1, A, resp1)
					case 1:
						handshake(1, S, resp1)
					case 2:
						handshake(1, AS, resp1)
					case 4:
						handshake(1, AC, resp1)
					case 5:
						handshake(1, SC, resp1)
					case 6:
						handshake(1, ASC, resp1)
					}
				}
				reset1ff = false
				reset3ff = false
				stage = 0
			} else if stage != 0 {
				minplace := 10
				for i := 0; i < 24; i++ {
					if multff[i] && placsw[i]+2 < minplace {
						minplace = placsw[i] + 2
					}
				}
				if stage == minplace+1 {
					if ier[0] == 'M' {
						handshake(1, DS, resp1)
					}
					if icand[0] == 'M' {
						handshake(1, RS, resp1)
					}
					Multl = false
					Multr = false
					stage = 12
				} else {
					stage++
				}
			}
		case c.val&Ccg != 0 && stage == 13:
			which := -1
			for i, f := range multff {
				if f {
					which = i
					break
				}
			}
			if iercl[which] == 1 {
				accclear(8)
			}
			if icandcl[which] == 1 {
				accclear(9)
			}
		case c.val&Onep != 0 && stage == 1:
			Multl = true
			Multr = true
			sigfig = -1
			for i := 0; i < 24; i++ {
				if multff[i] {
					sigfig = sigsw[i]
				}
			}
			if sigfig == 0 && lhppII != nil {
				handshake(1<<9, lhppII, resp1)
			} else if sigfig > 0 && sigfig < 9 && lhppI != nil {
				handshake(1<<uint(sigfig-1), lhppI, resp1)
			}
		case c.val&Fourp != 0 && stage == 1:
			if sigfig == 0 && lhppII != nil {
				handshake(1<<9, lhppII, resp1)
			} else if sigfig > 0 && sigfig < 9 && lhppI != nil {
				handshake(1<<uint(sigfig-1), lhppI, resp1)
			}
		case c.val&Onep != 0 && stage >= 2 && stage < 12:
			ier = accstat(8)
			icand = accstat(9)
			lhpp := 0
			rhpp := 0
			for i := 0; i < 10; i++ {
				ps10 := table10[ier[stage]-'0'][icand[i+2]-'0']
				ps1 := table1[ier[stage]-'0'][icand[i+2]-'0']
				if ps10.one == 1 {
					lhpp |= 1 << uint(9-i)
				}
				if ps1.one == 1 {
					rhpp |= 1 << uint(9-i)
				}
			}
			shiftprod(lhpp, rhpp, resp1, resp2, resp3, resp4)
		case c.val&Twop != 0 && stage >= 2 && stage < 12:
			lhpp := 0
			rhpp := 0
			for i := 0; i < 10; i++ {
				ps10 := table10[ier[stage]-'0'][icand[i+2]-'0']
				ps1 := table1[ier[stage]-'0'][icand[i+2]-'0']
				if ps10.two == 1 {
					lhpp |= 1 << uint(9-i)
				}
				if ps1.two == 1 {
					rhpp |= 1 << uint(9-i)
				}
			}
			shiftprod(lhpp, rhpp, resp1, resp2, resp3, resp4)
		case c.val&Twopp != 0 && stage >= 2 && stage < 12:
			lhpp := 0
			rhpp := 0
			for i := 0; i < 10; i++ {
				ps10 := table10[ier[stage]-'0'][icand[i+2]-'0']
				ps1 := table1[ier[stage]-'0'][icand[i+2]-'0']
				if ps10.twop == 1 {
					lhpp |= 1 << uint(9-i)
				}
				if ps1.twop == 1 {
					rhpp |= 1 << uint(9-i)
				}
			}
			shiftprod(lhpp, rhpp, resp1, resp2, resp3, resp4)
		case c.val&Fourp != 0 && stage >= 2 && stage < 12:
			lhpp := 0
			rhpp := 0
			for i := 0; i < 10; i++ {
				ps10 := table10[ier[stage]-'0'][icand[i+2]-'0']
				ps1 := table1[ier[stage]-'0'][icand[i+2]-'0']
				if ps10.four == 1 {
					lhpp |= 1 << uint(9-i)
				}
				if ps1.four == 1 {
					rhpp |= 1 << uint(9-i)
				}
			}
			shiftprod(lhpp, rhpp, resp1, resp2, resp3, resp4)
		case c.val&Onepp != 0 && stage >= 2 && stage < 12:
			minplace := 10
			for i := 0; i < 24; i++ {
				if multff[i] && placsw[i]+2 < minplace {
					minplace = placsw[i] + 2
				}
			}
			if stage == minplace+1 && ier[0] == 'M' && icand[0] == 'M' {
				handshake(1<<10, rhppI, resp1)
			}
		case c.val&Rp != 0 && buffer61:
			buffer61 = false
			f44 = true
		}
		if c.resp != nil {
			c.resp <- 1
		}
	}
}

func multunit2() {
	var p pulse

	for {
		p.resp = nil
		select {
		case <-multupdate:
		case p = <-multin[0]:
			multargs(0)
		case p = <-multin[1]:
			multargs(1)
		case p = <-multin[2]:
			multargs(2)
		case p = <-multin[3]:
			multargs(3)
		case p = <-multin[4]:
			multargs(4)
		case p = <-multin[5]:
			multargs(5)
		case p = <-multin[6]:
			multargs(6)
		case p = <-multin[7]:
			multargs(7)
		case p = <-multin[8]:
			multargs(8)
		case p = <-multin[9]:
			multargs(9)
		case p = <-multin[10]:
			multargs(10)
		case p = <-multin[11]:
			multargs(11)
		case p = <-multin[12]:
			multargs(12)
		case p = <-multin[13]:
			multargs(13)
		case p = <-multin[14]:
			multargs(14)
		case p = <-multin[15]:
			multargs(15)
		case p = <-multin[16]:
			multargs(16)
		case p = <-multin[17]:
			multargs(17)
		case p = <-multin[18]:
			multargs(18)
		case p = <-multin[19]:
			multargs(19)
		case p = <-multin[20]:
			multargs(20)
		case p = <-multin[21]:
			multargs(21)
		case p = <-multin[22]:
			multargs(22)
		case p = <-multin[23]:
			multargs(23)
		}
		if p.resp != nil {
			p.resp <- 1
		}
	}
}
