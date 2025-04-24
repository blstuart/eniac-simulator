package main

import (
	"fmt"
	"strconv"
	"time"
)

var gate66, gate69 int
var prff, printphase1, printphase2, rdff, rdilock, rdsync, rdfinish bool
var initjack [18]chan pulse
var initclrff [6]bool
var initupdate chan int
var prttmr time.Time

func initstat() string {
	s := ""
	for _, f := range initclrff {
		s += b2is(f)
	}
	s += b2is(rdff)
	s += b2is(prff)
	s += b2is(rdfinish)
	s += b2is(rdilock)
	s += b2is(rdsync)
	s += "00"
	s += fmt.Sprintf("%d%d", gate66, gate69)
	return s
}

func initreset() {
	gate66 = 0
	gate69 = 0
	prff = false
	rdff = false
	rdilock = false
	rdsync = false
	rdfinish = false
	for i := 0; i < 18; i++ {
		initjack[i] = nil
	}
	for i := 0; i < 6; i++ {
		initclrff[i] = false
	}
	initupdate <- 1
}

func initplug(jack string, ch chan pulse) {
	switch jack[0] {
	case 'c', 'C':
		set, _ := strconv.Atoi(jack[2:])
		if set >= 1 && set <= 6 {
			switch jack[1] {
			case 'i':
				initjack[2*(set-1)] = ch
			case 'o':
				initjack[2*(set-1)+1] = ch
			}
		}
	case 'i', 'I':
		initjack[17] = ch
	case 'p', 'P':
		switch jack[1] {
		case 'i':
			initjack[15] = ch
		case 'o':
			initjack[16] = ch
		}
	case 'r', 'R':
		switch jack[1] {
		case 'l':
			initjack[12] = ch
		case 'i':
			initjack[13] = ch
		case 'o':
			initjack[14] = ch
		}
	default:
		fmt.Println("Initiate unit jack syntax: i.jack")
	}
	initupdate <- 1
}

func initiateunit(cyctrunk chan pulse, button chan int) {
	var lastread time.Time

	resp := make(chan int)
	go initiateunit2()
	for {
		select {
		case p := <-cyctrunk:
			cyc := p.val
			if cyc&Cpp != 0 {
				if gate69 == 1 {
					gate66 = 0
					gate69 = 0
					handshake(1, initjack[17], resp)
				} else if gate66 == 1 {
					gate69 = 1
				}
				for i, ff := range initclrff {
					if ff {
						if initjack[2*i+1] != nil {
							initjack[2*i+1] <- pulse{1, resp}
							<-resp
						}
						initclrff[i] = false
					}
				}
				if rdsync {
					handshake(1, initjack[14], resp)
					rdff = false
					rdilock = false
					rdsync = false
					rdfinish = false
				}
				if rdff && time.Since(lastread) > 375*time.Millisecond {
					if cardscanner != nil {
						if cardscanner.Scan() {
							card := cardscanner.Text()
							proccard(card)
							lastread = time.Now()
							rdfinish = true
						} else {
							cardscanner = nil
						}
					}
				}
				if rdfinish && rdilock {
					rdsync = true
				}
				if printphase1 && time.Since(prttmr) > 150*time.Millisecond {
					s := doprint()
					if punchwriter != nil {
						punchwriter.WriteString(s)
						punchwriter.WriteByte('\n')
						punchwriter.Flush()
					} else {
						fmt.Println(s)
					}
					if ppunch != nil {
						ppunch <- s
					}
					handshake(1, initjack[16], resp)
					prttmr = time.Now()
					printphase1 = false
					printphase2 = true
					prff = false
				}
				if printphase2 && time.Since(prttmr) > 450*time.Millisecond {
					if prff {
						prttmr = time.Now()
						printphase1 = true
					}
					printphase2 = false
				}
			}
			if p.resp != nil {
				p.resp <- 1
			}
		case bu := <-button:
			switch bu {
			//caes 3:
			case 4:
				gate66 = 1
			case 5:
				mpclear()
				for i := 0; i < 20; i++ {
					accclear(i)
				}
				divclear()
				multclear()
			case 3:
				rdff = true
				rdilock = true
			}
		}
	}
}

func initiateunit2() {
	initupdate = make(chan int)
	for {
		select {
		case <-initupdate:
		case p := <-initjack[12]:
			rdilock = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[13]:
			rdff = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[15]:
			if !printphase1 {
				prff = true
				if !printphase2 {
					printphase1 = true
					prttmr = time.Now()
				}
			}
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[0]:
			initclrff[0] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[2]:
			initclrff[1] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[4]:
			initclrff[2] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[6]:
			initclrff[3] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[8]:
			initclrff[4] = true
			if p.resp != nil {
				p.resp <- 1
			}
		case p := <-initjack[10]:
			initclrff[5] = true
			if p.resp != nil {
				p.resp <- 1
			}
		}
	}
}
