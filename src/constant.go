package main

import (
	"fmt"
	"strconv"
	"unicode"
)

var consupdate chan int
var consel [30]int
var conscard [8][10]byte
var signcard [8][2]byte
var consj [10]byte
var signj [2]byte
var consk [10]byte
var signk [2]byte
var consout chan pulse
var conspin [30]chan pulse
var consinff1, consinff2 [30]bool
var conspout [30]chan pulse

func consstat() string {
	s := ""
	for _, f := range consinff2 {
		s += b2is(f)
	}
	return s
}

func consreset() {
	for i := 0; i < 30; i++ {
		consel[i] = 0
		conspin[i] = nil
		consinff1[i] = false
		consinff2[i] = false
		conspout[i] = nil
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			conscard[i][j] = 0
		}
		signcard[i][0] = 0
		signcard[i][1] = 0
	}
	for i := 0; i < 10; i++ {
		consj[i] = 0
		consk[i] = 0
	}
	signj[0] = 0
	signj[1] = 0
	signk[0] = 0
	signk[1] = 0
	consout = nil
	consupdate <- 1
}

func consplug(jack string, ch chan pulse) {
	var prog int
	var ilk rune

	if jack == "o" {
		consout = ch
	} else {
		fmt.Sscanf(jack, "%d%c", &prog, &ilk)
		if ilk == 'i' {
			conspin[prog-1] = ch
		} else {
			conspout[prog-1] = ch
		}
	}
	consupdate <- 1
}

func pm2int(ch string) byte {
	if ch == "p" || ch == "P" {
		return 0
	} else {
		return 1
	}
}

func consctl(ch chan [2]string) {
	var n int

	selector := map[string]int{"l": 0, "r": 1, "lr": 2}
	for {
		ctl := <-ch
		switch ctl[0][0] {
		case 's':
			prog, _ := strconv.Atoi(ctl[0][1:])
			switch ctl[1][0] {
			case 'a', 'A', 'c', 'C', 'e', 'E', 'g', 'G', 'j', 'J':
				n = 0
			case 'b', 'B', 'd', 'D', 'f', 'F', 'h', 'H', 'k', 'K':
				n = 3
			}
			consel[prog-1] = n + selector[ctl[1][1:]]
		case 'j', 'J':
			if ctl[0][1] == 'l' {
				signj[0] = pm2int(ctl[1])
			} else if ctl[0][1] == 'r' {
				signj[1] = pm2int(ctl[1])
			} else {
				digit, _ := strconv.Atoi(ctl[0][1:])
				n, _ := strconv.Atoi(ctl[1])
				consj[digit-1] = byte(n)
			}
		case 'k', 'K':
			if ctl[0][1] == 'l' {
				signk[0] = pm2int(ctl[1])
			} else if ctl[0][1] == 'r' {
				signk[1] = pm2int(ctl[1])
			} else {
				digit, _ := strconv.Atoi(ctl[0][1:])
				n, _ := strconv.Atoi(ctl[1])
				consk[digit-1] = byte(n)
			}
		default:
			fmt.Println("Invalid constant switch: ", ctl[0])
		}
	}
}

func proccard(c string) {
	l := len(c)
	if l > 80 {
		l = 80
	}
	for i := 0; i < l/10; i++ {
		procfield(i, c[10*i:10*i+10])
	}
}

func procfield(i int, f string) {
	if len(f) < 10 || f == "          " {
		return
	}
	bank := i / 2
	tendig := true
	for j := bank * 6; j < bank*6+6; j++ {
		switch consel[j] {
		case 0, 1:
			if i%2 == 0 && conspin[j] != nil {
				tendig = false
			}
		case 3, 4:
			if i%2 == 1 && conspin[j] != nil {
				tendig = false
			}
		}
	}
	if tendig {
		donum(i, 0, f)
	} else {
		donum(i, 0, f[:5])
		donum(i, 5, f[5:])
	}
}

func donum(i, off int, f string) {
	var nz int

	neg := byte(0)
	for _, c := range f {
		if c == '-' || c == ']' || c == '}' || c >= 'J' && c <= 'R' {
			neg = 1
		}
	}
	if off == 0 {
		signcard[i][0] = neg
	} else {
		signcard[i][1] = neg
	}
	if neg == 0 {
		for j, c := range f {
			if unicode.IsDigit(c) {
				conscard[i][9-(j+off)] = charval(c)
			} else {
				conscard[i][9-(j+off)] = 0
			}
		}
		return
	}
	l := len(f)
	for nz = l - 1; nz >= 0 && f[nz] == '0'; nz-- {
		conscard[i][9-(nz+off)] = '0'
	}
	for ; nz >= 0; nz-- {
		conscard[i][9-(nz+off)] = 9 - charval(rune(f[nz]))
	}
}

func charval(c rune) byte {
	if c == '-' || c == ']' || c == '}' {
		return 0
	}
	if c >= 'J' && c <= 'R' {
		return byte(c - 'J' + 1)
	}
	return byte(c - '0')
}

func getval(sel int) (sgn byte, val []byte, pos1pp int) {
	if sel >= 24 {
		pos1pp = -1
		switch consel[sel] {
		case 0:
			sgn = signj[0]
			val = make([]byte, 10)
			copy(val[5:], consj[5:])
			return
		case 1:
			sgn = signj[1]
			val = make([]byte, 10)
			if sgn != 0 {
				for i := 5; i < 10; i++ {
					val[i] = 9
				}
			}
			copy(val[:5], consj[:5])
			return
		case 2:
			return signj[0], consj[:], -1
		case 3:
			sgn = signk[0]
			val = make([]byte, 10)
			copy(val[5:], consk[5:])
			return
		case 4:
			sgn = signk[1]
			val = make([]byte, 10)
			if sgn != 0 {
				for i := 5; i < 10; i++ {
					val[i] = 9
				}
			}
			copy(val[:5], consk[:5])
			return
		case 5:
			return signk[0], consk[:], -1
		}
	} else {
		bank := sel / 6
		switch consel[sel] {
		case 0:
			sgn = signcard[2*bank][1]
			val = make([]byte, 10)
			if sgn != 0 {
				for i := 0; i < 5; i++ {
					val[i] = 9
				}
			}
			copy(val[5:], conscard[2*bank][5:])
			pos1pp = 0
			return
		case 1:
			sgn = signcard[2*bank][0]
			val = make([]byte, 10)
			copy(val[:5], conscard[2*bank][:5])
			pos1pp = 5
			return
		case 2:
			return signcard[2*bank][0], conscard[2*bank][:], 0
		case 3:
			sgn = signcard[2*bank+1][1]
			val = make([]byte, 10)
			if sgn != 0 {
				for i := 0; i < 5; i++ {
					val[i] = 9
				}
			}
			copy(val[5:], conscard[2*bank+1][5:])
			pos1pp = 0
			return
		case 4:
			sgn = signcard[2*bank+1][0]
			val = make([]byte, 10)
			copy(val[:5], conscard[2*bank+1][:5])
			pos1pp = 5
			return
		case 5:
			return signcard[2*bank+1][0], conscard[2*bank+1][:], 0
		}
	}
	return
}

var digitcons = []int{0, Onep, Twop, (Onep | Twop), Fourp, (Onep | Fourp),
	(Twop | Fourp), (Onep | Twop | Fourp), (Twop | Twopp | Fourp),
	(Onep | Twop | Twopp | Fourp)}

func consunit(cyctrunk chan pulse) {
	var val []byte
	var sign byte
	var pos1pp int
	var whichrp bool

	consupdate = make(chan int)
	go consunit2()
	resp := make(chan int)
	for {
		p := <-cyctrunk
		sending := -1
		for i := 0; i < 30; i++ {
			if consinff2[i] {
				sending = i
				sign, val, pos1pp = getval(i)
				break
			}
		}
		cyc := p.val
		if cyc&Ccg != 0 {
			whichrp = false
		} else if cyc&Rp != 0 {
			if whichrp {
				for i := 0; i < 30; i++ {
					if consinff1[i] {
						consinff1[i] = false
						consinff2[i] = true
					}
				}
				whichrp = false
			} else {
				whichrp = true
			}
		}
		if sending > -1 {
			if cyc&Cpp != 0 {
				handshake(1, conspout[sending], resp)
				consinff2[sending] = false
				sending = -1
			} else if cyc&Ninep != 0 {
				n := 0
				for i := uint(0); i < uint(10); i++ {
					if cyc&digitcons[val[i]] != 0 {
						n |= 1 << i
					}
				}
				if sign == 1 {
					n |= 1 << 10
				}
				if n != 0 {
					handshake(n, consout, resp)
				}
			} else if cyc&Onepp != 0 && pos1pp >= 0 && sign == 1 {
				handshake(1<<uint(pos1pp), consout, resp)
			}
		}
		if p.resp != nil {
			p.resp <- 1
		}
	}
}

func consunit2() {
	var prog pulse

	for {
		select {
		case <-consupdate:
		case prog = <-conspin[0]:
			if prog.val == 1 {
				consinff1[0] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[1]:
			if prog.val == 1 {
				consinff1[1] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[2]:
			if prog.val == 1 {
				consinff1[2] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[3]:
			if prog.val == 1 {
				consinff1[3] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[4]:
			if prog.val == 1 {
				consinff1[4] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[5]:
			if prog.val == 1 {
				consinff1[5] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[6]:
			if prog.val == 1 {
				consinff1[6] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[7]:
			if prog.val == 1 {
				consinff1[7] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[8]:
			if prog.val == 1 {
				consinff1[8] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[9]:
			if prog.val == 1 {
				consinff1[9] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[10]:
			if prog.val == 1 {
				consinff1[10] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[11]:
			if prog.val == 1 {
				consinff1[11] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[12]:
			if prog.val == 1 {
				consinff1[12] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[13]:
			if prog.val == 1 {
				consinff1[13] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[14]:
			if prog.val == 1 {
				consinff1[14] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[15]:
			if prog.val == 1 {
				consinff1[15] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[16]:
			if prog.val == 1 {
				consinff1[16] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[17]:
			if prog.val == 1 {
				consinff1[17] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[18]:
			if prog.val == 1 {
				consinff1[18] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[19]:
			if prog.val == 1 {
				consinff1[19] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[20]:
			if prog.val == 1 {
				consinff1[20] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[21]:
			if prog.val == 1 {
				consinff1[21] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[22]:
			if prog.val == 1 {
				consinff1[22] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[23]:
			if prog.val == 1 {
				consinff1[23] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[24]:
			if prog.val == 1 {
				consinff1[24] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[25]:
			if prog.val == 1 {
				consinff1[25] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[26]:
			if prog.val == 1 {
				consinff1[26] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[27]:
			if prog.val == 1 {
				consinff1[27] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[28]:
			if prog.val == 1 {
				consinff1[28] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		case prog = <-conspin[29]:
			if prog.val == 1 {
				consinff1[29] = true
			}
			if prog.resp != nil {
				prog.resp <- 1
			}
		}
	}
}
