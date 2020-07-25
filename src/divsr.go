package main

import (
	"fmt"
	"strconv"
)

const (
	svα   = 1 << 0
	svβ   = 1 << 1
	svγ   = 1 << 2
	svA   = 1 << 5
	svCLR = 1 << 8

	su2qα   = 1 << 0
	su2qA   = 1 << 3
	su2qS   = 1 << 4
	su2qCLR = 1 << 5
	su2sα   = 1 << 1
	su2sA   = 1 << 2
	su2sCLR = 1 << 6

	su3α   = 1 << 0
	su3β   = 1 << 1
	su3γ   = 1 << 2
	su3A   = 1 << 3
	su3S   = 1 << 4
	su3CLR = 1 << 5
)

// Lots of state vars and I don't want to worry about name collisions
var divsr struct {
	divupdate                                                                  chan int
	progin, progout, ilock                                                     [8]chan pulse
	answer                                                                     chan pulse
	numarg, numcl, denarg, dencl, roundoff, places, ilocksw, anssw             [8]int
	preff, progff                                                              [8]bool
	placering, progring                                                        int
	divff, clrff, ilockff, coinff, dpγ, nγ, psrcff, pringff, denomff, numrplus bool
	numrmin, qα, sac, m2, m1, nac, da, nα, dα, dγ, npγ, p2, p1, sα, ds, nβ, dβ bool
	ans1, ans2, ans3, ans4                                                     bool
	curprog, divadap, sradap                                                   int
	sv, su2, su3                                                               int
}

func divsrstat() string {
	s := fmt.Sprintf("%d %d ", divsr.placering, divsr.progring)
	for i := 0; i < 8; i++ {
		if divsr.progff[i] {
			s += "1"
		} else {
			s += "0"
		}
	}
	s += " " + b2is(divsr.divff) + b2is(divsr.clrff) + b2is(divsr.coinff) + b2is(divsr.dpγ) +
		b2is(divsr.nγ) + b2is(divsr.psrcff) + b2is(divsr.pringff) + b2is(divsr.denomff) +
		b2is(divsr.numrplus) + b2is(divsr.numrmin) + b2is(divsr.qα) + b2is(divsr.sac) +
		b2is(divsr.m2) + b2is(divsr.m1) + b2is(divsr.nac) + b2is(divsr.da) + b2is(divsr.nα) +
		b2is(divsr.dα) + b2is(divsr.dγ) + b2is(divsr.npγ) + b2is(divsr.p2) + b2is(divsr.p1) +
		b2is(divsr.sα) + b2is(divsr.ds) + b2is(divsr.nβ) + b2is(divsr.dβ) + b2is(divsr.ans1) +
		b2is(divsr.ans2) + b2is(divsr.ans3) + b2is(divsr.ans4)
	return s
}

func divsrstat2() string {
	s := fmt.Sprintf("%d %d ", divsr.placering, divsr.progring)
	for i := 0; i < 8; i++ {
		if divsr.progff[i] {
			s += "1"
		} else {
			s += "0"
		}
	}
	if divsr.divff {
		s += " divff"
	}
	if divsr.clrff {
		s += " clrff"
	}
	if divsr.coinff {
		s += " coinff"
	}
	if divsr.dpγ {
		s += " dpg"
	}
	if divsr.nγ {
		s += " ng"
	}
	if divsr.psrcff {
		s += " psrcff"
	}
	if divsr.denomff {
		s += " denomff"
	}
	if divsr.numrplus {
		s += " n+"
	}
	if divsr.numrmin {
		s += " n-"
	}
	if divsr.qα {
		s += " qa"
	}
	if divsr.sac {
		s += " SAC"
	}
	if divsr.m2 {
		s += " -2"
	}
	if divsr.m1 {
		s += " -1"
	}
	if divsr.nac {
		s += " NAC"
	}
	if divsr.da {
		s += " dA"
	}
	if divsr.nα {
		s += " na"
	}
	if divsr.dα {
		s += " da"
	}
	if divsr.dγ {
		s += " dg"
	}
	if divsr.npγ {
		s += " npg"
	}
	if divsr.p2 {
		s += " +2"
	}
	if divsr.p1 {
		s += " +1"
	}
	if divsr.sα {
		s += " sa"
	}
	if divsr.ds {
		s += " dS"
	}
	if divsr.nβ {
		s += " nb"
	}
	if divsr.dβ {
		s += " db"
	}
	if divsr.ans1 {
		s += " A1"
	}
	if divsr.ans2 {
		s += " A2"
	}
	if divsr.ans3 {
		s += " A3"
	}
	if divsr.ans4 {
		s += " A4"
	}
	return s
}

func divreset() {
	for i := 0; i < 8; i++ {
		divsr.progin[i] = nil
		divsr.progout[i] = nil
		divsr.ilock[i] = nil
		divsr.numarg[i] = 0
		divsr.numcl[i] = 0
		divsr.denarg[i] = 0
		divsr.dencl[i] = 0
		divsr.roundoff[i] = 0
		divsr.places[i] = 0
		divsr.ilocksw[i] = 0
		divsr.anssw[i] = 0
		divsr.preff[i] = false
		divsr.progff[i] = false
	}
	divsr.answer = nil
	divsr.divff = false
	divsr.ilockff = false
	divsr.ans1 = false
	divsr.ans2 = false
	divsr.ans3 = false
	divsr.ans4 = false
	divsr.divadap = 0
	divsr.sradap = 0
	divclear()
	divsr.divupdate <- 1
}

func divclear() {
	divintclear()
	divsr.sv = 0
	divsr.su2 = 0
	divsr.su3 = 0
}

func divintclear() {
	divsr.progring = 0
	divsr.placering = 0
	divsr.numrplus = true
	divsr.numrmin = false
	divsr.denomff = false
	divsr.psrcff = false
	divsr.pringff = false
	divsr.curprog = -1
	divsr.coinff = false
	divsr.clrff = false
	divsr.dpγ = false
	divsr.nγ = false
	divsr.qα = false
	divsr.sac = false
	divsr.m2 = false
	divsr.m1 = false
	divsr.nac = false
	divsr.da = false
	divsr.nα = false
	divsr.dα = false
	divsr.dγ = false
	divsr.npγ = false
	divsr.p2 = false
	divsr.p1 = false
	divsr.sα = false
	divsr.ds = false
	divsr.nβ = false
	divsr.dβ = false
}

func divsrplug(jack string, ch chan pulse) {
	var prog int
	var ilk rune

	if jack == "ans" || jack == "ANS" {
		divsr.answer = ch
	} else {
		fmt.Sscanf(jack, "%d%c", &prog, &ilk)
		switch ilk {
		case 'i':
			divsr.progin[prog-1] = ch
		case 'o':
			divsr.progout[prog-1] = ch
		case 'l':
			divsr.ilock[prog-1] = ch
		}
	}
	divsr.divupdate <- 1
}

func divsrctl(ch chan [2]string) {
	for {
		ctl := <-ch
		sw, _ := strconv.Atoi(ctl[0][2:])
		switch ctl[0][:2] {
		case "da":
			divsr.divadap = int(ctl[1][0] - 'A')
		case "ra":
			divsr.sradap = int(ctl[1][0] - 'A')
		case "nr":
			switch ctl[1] {
			case "α", "a", "alpha":
				divsr.numarg[sw-1] = 0
			case "β", "b", "beta":
				divsr.numarg[sw-1] = 1
			case "0":
				divsr.numarg[sw-1] = 2
			}
		case "nc":
			if ctl[1] == "C" || ctl[1] == "c" {
				divsr.numcl[sw-1] = 1
			} else {
				divsr.numcl[sw-1] = 0
			}
		case "dr":
			switch ctl[1] {
			case "α", "a", "alpha":
				divsr.denarg[sw-1] = 0
			case "β", "b", "beta":
				divsr.denarg[sw-1] = 1
			case "0":
				divsr.denarg[sw-1] = 2
			}
		case "dc":
			if ctl[1] == "C" || ctl[1] == "c" {
				divsr.dencl[sw-1] = 1
			} else {
				divsr.dencl[sw-1] = 0
			}
		case "pl":
			offset := 5
			pl, _ := strconv.Atoi(ctl[1][1:])
			if ctl[1][0] == 'D' {
				offset = 0
			}
			if pl == 4 {
				divsr.places[sw-1] = 0 + offset
			} else {
				divsr.places[sw-1] = pl - 6 + offset
			}
		case "ro":
			if ctl[1] == "RO" || ctl[1] == "ro" {
				divsr.roundoff[sw-1] = 1
			} else {
				divsr.roundoff[sw-1] = 0
			}
		case "an":
			if ctl[1] == "OFF" || ctl[1] == "off" {
				divsr.anssw[sw-1] = 4
			} else {
				val, _ := strconv.Atoi(ctl[1])
				divsr.anssw[sw-1] = val - 1
			}
		case "il":
			if ctl[1] == "I" || ctl[1] == "i" {
				divsr.ilocksw[sw-1] = 1
			} else {
				divsr.ilocksw[sw-1] = 0
			}
		default:
			fmt.Printf("Invalid divider switch %s\n", ctl[0])
		}
	}
}

func divargs(prog int) {
	divsr.preff[prog] = true
	if divsr.places[prog] < 5 {
		divsr.divff = true
	} else {
		divsr.divff = false
	}
	switch divsr.numarg[prog] {
	case 0:
		divsr.nα = true
		divsr.sv |= svα
	case 1:
		divsr.nβ = true
		divsr.sv |= svβ
	}
	switch divsr.denarg[prog] {
	case 0:
		divsr.dα = true
		divsr.su3 |= su3α
	case 1:
		divsr.dβ = true
		divsr.su3 |= su3β
	}
}

func doP() {
	divsr.nγ = true
	divsr.sv |= svγ
	if samesign() {
		divsr.ds = true
		divsr.su3 |= su3S
	} else {
		divsr.da = true
		divsr.su3 |= su3A
	}
}

func doS() {
	divsr.sα = true
	divsr.su2 |= su2sα
	divsr.nac = true
	divsr.sv |= svA | svCLR
	if !divsr.divff {
		if samesign() {
			divsr.m1 = true
		} else {
			divsr.p1 = true
		}
		divsr.dpγ = true
		divsr.su3 |= su3γ
	}
	p := divsr.places[divsr.curprog] % 5
	if p == 0 {
		p = 4
	} else {
		p += 6
	}
	if divsr.placering == p-2 { // Gate E6
		divsr.psrcff = true
	}
}

func samesign() bool {
	return divsr.denomff && divsr.numrmin || !divsr.denomff && divsr.numrplus
}

func overflow() bool {
	s := accstat(2)
	return s[0] == 'P' && divsr.numrmin || s[0] == 'M' && divsr.numrplus
}

func doGP(resp chan int) {
	if divsr.coinff { // Gate E50
		if divsr.ilocksw[divsr.curprog] == 0 || divsr.ilockff {
			divsr.coinff = false
			divsr.clrff = true
			return
		}
	} else if divsr.clrff {
		divsr.progff[divsr.curprog] = false
		handshake(1, divsr.progout[divsr.curprog], resp)
		if divsr.ilocksw[divsr.curprog] == 1 {
			divsr.ilockff = false
		}
		/*
		 * Implement the PX-4-114 adapters
		 */
		switch divsr.anssw[divsr.curprog] {
		case 0:
			divsr.ans1 = true
			divsr.su2 |= su2qA
			if divsr.divadap == 2 {
				divsr.su2 |= su2qCLR
			}
		case 1:
			divsr.ans2 = true
			switch divsr.divadap {
			case 0:
				divsr.su2 |= su2qA | su2qCLR
			case 1:
				divsr.su2 |= su2qS
			case 2:
				divsr.su2 |= su2qS | su2qCLR
			}
		case 2:
			divsr.ans3 = true
			divsr.su3 |= su3A
			if divsr.sradap == 2 {
				divsr.su3 |= su3CLR
			}
		case 3:
			divsr.ans4 = true
			switch divsr.sradap {
			case 0:
				divsr.su3 |= su3A | su3CLR
			case 1:
				divsr.su3 |= su3S
			case 2:
				divsr.su3 |= su3S | su3CLR
			}
		}
		if divsr.numcl[divsr.curprog] == 1 {
			accclear(2)
		}
		if divsr.dencl[divsr.curprog] == 1 {
			accclear(4)
		}
		divintclear()
		return
	}
	if divsr.qα {
		divsr.p1 = false
		divsr.m1 = false
		if overflow() { // Gates D9, D11, D12
			doS()
		} else {
			doP()
		}
		divsr.qα = false
		divsr.su2 &^= su2qα
	} else if divsr.nγ { //  Gates L10, G11, H11
		divsr.nγ = false
		divsr.sv &^= svγ
		if divsr.divff {
			divsr.qα = true
			divsr.su2 |= su2qα
			if divsr.ds {
				divsr.ds = false
				divsr.su3 &^= su3S
				divsr.p1 = true
			} else if divsr.da {
				divsr.da = false
				divsr.su3 &^= su3A
				divsr.m1 = true
			}
		} else {
			divsr.dγ = true
			divsr.su3 |= su3γ
			if divsr.ds {
				divsr.ds = false
				divsr.su3 &^= su3S
				divsr.p2 = true
			} else if divsr.da {
				divsr.da = false
				divsr.su3 &^= su3A
				divsr.m2 = true
			}
		}
	} else if divsr.npγ { // Gate C9
		divsr.npγ = false
		divsr.sv &^= svγ
		divsr.sac = false
		divsr.su2 &^= su2sA | su2sCLR
		divsr.m1 = false
		divsr.p1 = false
		divsr.dpγ = false
		divsr.su3 &^= su3γ
		doP()
	} else if divsr.sα { // Gates K7, L1
		divsr.sα = false
		divsr.su2 &^= su2sα
		divsr.nac = false
		divsr.sv &^= svA | svCLR
		divsr.sac = true
		divsr.su2 |= su2sA | su2sCLR
		divsr.npγ = true
		divsr.sv |= svγ
		divsr.numrplus, divsr.numrmin = divsr.numrmin, divsr.numrplus
	} else if divsr.dγ {
		divsr.dγ = false
		divsr.su3 &^= su3γ
		divsr.p2 = false
		divsr.m2 = false
		if overflow() {
			doS()
		} else {
			doP()
		}
	}
	switch divsr.progring {
	case 0:
		divsr.nα = false
		divsr.nβ = false
		divsr.sv &^= svα | svβ
		divsr.dα = false
		divsr.dβ = false
		divsr.su3 &^= su3α | su3β
		if !divsr.pringff {
			divsr.progring++
		}
	case 1: // Gate D6
		s := accstat(2)
		if s[0] == 'M' {
			divsr.numrplus, divsr.numrmin = divsr.numrmin, divsr.numrplus
		}
		s = accstat(4)
		if s[0] == 'M' {
			divsr.denomff = true
		}
		if !divsr.pringff {
			divsr.progring++
		}
	case 2: // Gate A7, B7, B8
		if divsr.divff {
			doP()
			divsr.pringff = true
			divsr.progring = 0
		} else {
			divsr.p1 = true
			divsr.dγ = true
			divsr.su3 |= su3γ
			divsr.progring++
		}
	case 3:
		divsr.p1 = false
		divsr.dγ = false
		divsr.su3 &^= su3γ
		doP()
		divsr.pringff = true
		divsr.progring = 0
	}
}

func doIIIP() {
	if divsr.npγ { // Gate C9
		divsr.npγ = false
		divsr.sv &^= svγ
		divsr.sac = false
		divsr.su2 &^= su2sA | su2sCLR
		divsr.m1 = false
		divsr.p1 = false
		divsr.dpγ = false
		divsr.su3 &^= su3γ
	} else if divsr.sα {
		divsr.sα = false
		divsr.su2 &^= su2sα
		divsr.nac = false
		divsr.sv &^= svA | svCLR
		divsr.sac = true
		divsr.su2 |= su2sA | su2sCLR
		divsr.npγ = true
		divsr.sv |= svγ
		if divsr.psrcff {
			divsr.dpγ = false
			divsr.su3 &^= su3γ
			divsr.m1 = false
			divsr.p1 = false
		}
		divsr.numrplus, divsr.numrmin = divsr.numrmin, divsr.numrplus
	} else if divsr.qα {
		divsr.qα = false
		divsr.su2 &^= su2qα
		divsr.m1 = false
		divsr.p1 = false
	} else if divsr.dγ {
		divsr.dγ = false
		divsr.su3 &^= su3γ
		divsr.m2 = false
		divsr.p2 = false
	}
	switch divsr.progring {
	case 1:
		doP()
	case 6: // Gate D4
		divsr.nγ = false
		divsr.sv &^= svγ
		divsr.da = false
		divsr.ds = false
		divsr.su3 &^= su3A | su3S
	case 7: // Gate J13
		if !overflow() && divsr.roundoff[divsr.curprog] == 1 { // Gate K12
			if divsr.divff {
				divsr.qα = true
				divsr.su2 |= su2qα
				if samesign() {
					divsr.p1 = true
				} else {
					divsr.m1 = true
				}
			} else {
				divsr.dγ = true
				divsr.su3 |= su3γ
				if samesign() {
					divsr.p2 = true
				} else {
					divsr.m2 = true
				}
			}
		}
	case 8: // Gate E3. L50
		divsr.psrcff = false
		divsr.coinff = true
	}
	divsr.progring++
}

func divunit(cyctrunk chan pulse) {
	resp := make(chan int)
	divsr.divupdate = make(chan int)
	go divunit2()
	divintclear()
	for {
		p := <-cyctrunk
		switch {
		case p.val&Cpp != 0:
			if divsr.progring == 0 {
				divsr.ans1 = false
				divsr.ans2 = false
				divsr.ans3 = false
				divsr.ans4 = false
				divsr.su2 &^= su2qA | su2qS | su2qCLR
				divsr.su3 &^= su3A | su3S | su3CLR
			}
			if divsr.curprog >= 0 {
				if divsr.psrcff == false { // Gate F4
					doGP(resp)
				} else { // Gate F5
					doIIIP()
				}
			}
		case p.val&Rp != 0:
			/*
			 * Ugly hack to avoid races
			 */
			for i := 0; i < 8; i++ {
				if divsr.preff[i] {
					divsr.preff[i] = false
					divsr.progff[i] = true
					divsr.curprog = i
				}
			}
		case p.val&Onep != 0 && divsr.p1 || p.val&Twop != 0 && divsr.p2:
			if divsr.placering < 9 {
				handshake(1<<uint(8-divsr.placering), divsr.answer, resp)
			}
		case p.val&Onep != 0 && divsr.m2 || p.val&Twopp != 0 && divsr.m1:
			handshake(0x7ff, divsr.answer, resp)
		case p.val&Onep != 0 && divsr.m1 || p.val&Twopp != 0 && divsr.m2:
			if divsr.placering < 9 {
				handshake(0x7ff^(1<<uint(8-divsr.placering)), divsr.answer, resp)
			} else {
				handshake(0x7ff, divsr.answer, resp)
			}
		case (p.val&Fourp != 0 || p.val&Twop != 0) && (divsr.m1 || divsr.m2):
			handshake(0x7ff, divsr.answer, resp)
		case p.val&Onepp != 0:
			if divsr.m1 || divsr.m2 {
				handshake(1, divsr.answer, resp)
			}
			if divsr.psrcff == false && divsr.sα { // Gate L45
				divsr.placering++
			}
		}
		if p.resp != nil {
			p.resp <- 1
		}
	}
}

func divunit2() {
	var p pulse

	for {
		p.resp = nil
		select {
		case <-divsr.divupdate:
		case p = <-divsr.progin[0]:
			divargs(0)
		case p = <-divsr.progin[1]:
			divargs(1)
		case p = <-divsr.progin[2]:
			divargs(2)
		case p = <-divsr.progin[3]:
			divargs(3)
		case p = <-divsr.progin[4]:
			divargs(4)
		case p = <-divsr.progin[5]:
			divargs(5)
		case p = <-divsr.progin[6]:
			divargs(6)
		case p = <-divsr.progin[7]:
			divargs(7)
		case p = <-divsr.ilock[0]:
			divsr.ilockff = true
		case p = <-divsr.ilock[1]:
			divsr.ilockff = true
		case p = <-divsr.ilock[2]:
			divsr.ilockff = true
		case p = <-divsr.ilock[3]:
			divsr.ilockff = true
		case p = <-divsr.ilock[4]:
			divsr.ilockff = true
		case p = <-divsr.ilock[5]:
			divsr.ilockff = true
		case p = <-divsr.ilock[6]:
			divsr.ilockff = true
		case p = <-divsr.ilock[7]:
			divsr.ilockff = true
		}
		if p.resp != nil {
			p.resp <- 1
		}
	}
}
