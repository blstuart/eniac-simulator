package main

import (
	"time"
	"fmt"
	"strconv"
)

const (
	Cpp = 1 << iota
	Onep
	Ninep
	Tenp
	Scg
	Rp
	Onepp
	Ccg
	Twop
	Twopp
	Fourp
)

const (
	Pulse = iota
	Add
	Cont
)

var clocks = []int{
	0, Tenp, // 0
	Onep | Ninep, Tenp, // 1
	Twop | Ninep, Tenp, // 2
	Twop | Ninep, Tenp, // 3
	Twopp | Ninep, Tenp, // 4
	Twopp | Ninep, Tenp, // 5
	Fourp | Ninep, Tenp, // 6
	Fourp | Ninep, Tenp, // 7
	Fourp | Ninep, Tenp, // 8
	Fourp | Ninep, Tenp, // 9
	Onepp, 0, // 10
	Ccg, 0, // 11
	0, 0, // 12
	Rp, 0, // 13
	0, 0, // 14
	0, 0, // 15
	0, 0, // 16
	Cpp, 0, // 17
	0, 0, // 18
	Rp, 0, // 19
}

var cmode = Cont
var cyc = 0
var intbch chan int
var cyperiod = 0

func cycstat() string {
	if cyc >= len(clocks) {
		return "0"
	} else {
		return fmt.Sprintf("%d", cyc)
	}
}

func cycreset() {
	oldmode := cmode
	cmode = Cont
	cyperiod = 0
	if intbch != nil && (oldmode == Add || oldmode == Pulse) {
		intbch <- 1
	}
}

func cyclectl(cch chan [2]string) {
	for {
		x := <-cch
		switch x[0] {
		case "op":
			oldmode := cmode
			switch x[1] {
			case "1p", "1P":
				cmode = Pulse
			case "1a", "1A":
				cmode = Add
			case "co", "CO":
				cmode = Cont
				if oldmode == Add || oldmode == Pulse {
					intbch <- 1
				}
			case "cy", "CY":
				cmode = (cmode + 1) % 3
				if oldmode == Add || oldmode == Pulse {
					intbch <- 1
				}
			default:
				fmt.Println("cycle unit op swtch value: one of 1p, 1a, co, cy")
			}
		case "rate":
			cyrate, _ := strconv.Atoi(x[1])
			if cyrate == 0 {
				cyperiod = 0
			} else {
				cyperiod = 500000 / cyrate
			}
		default:
			fmt.Println("cycle unit switch: s cy.op.val")
		}
	}
}

func cycleunit(cch chan pulse, bch chan int) {
	var p pulse

	intbch = make(chan int)
	go func() {
		for {
			b := <-bch
			if cmode == Add || cmode == Pulse {
				intbch <- b
			}
		}
	}()

	p.resp = make(chan int)
	for {
		if cmode == Add {
			<-intbch
		}
		for cyc = 0; cyc < len(clocks); cyc++ {
			if cmode == Pulse {
				<-intbch
			}
			if cyc == 32 && (initclrff[0] || initclrff[1] || initclrff[2] ||
				initclrff[3] || initclrff[4] || initclrff[5]) {
				p.val = Scg
				cch <- p
				<-p.resp
			} else if clocks[cyc] != 0 {
				p.val = clocks[cyc]
				cch <- p
				if cyperiod != 0 {
					time.Sleep(time.Duration(cyperiod) * time.Microsecond)
				}
				<-p.resp
			} else if cyperiod != 0 {
				time.Sleep(time.Duration(cyperiod) * time.Microsecond)
			}
			cyc++
			if clocks[cyc] != 0 {
				p.val = clocks[cyc]
				cch <- p
				if cyperiod != 0 {
					time.Sleep(time.Duration(cyperiod) * time.Microsecond)
				}
				<-p.resp
			} else if cyperiod != 0 {
				time.Sleep(time.Duration(cyperiod) * time.Microsecond)
			}
		}
	}
}
