package main

import (
	"strconv"
	"strings"
)

var prtsw = [32]int{0, 1, 0, 1, 0, 0, 1, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 1, 0, 1, 0, 1, 0}

var blank = "     "

func prreset() {
	for i := 0; i < 32; i++ {
		prtsw[i] = 0
	}
	prtsw[1] = 1
	prtsw[3] = 1
	prtsw[6] = 1
	prtsw[24] = 1
	prtsw[26] = 1
	prtsw[28] = 1
	prtsw[30] = 1
}

func doprint() (s string) {
	var raw [9]string
	var sgn [16]byte

	raw[0] = mpstat()
	for i := 1; i < 9; i++ {
		raw[i] = accstat(i + 11)
	}
	p1 := raw[0][13:18]
	sgn[0] = 'P'
	p1 += raw[1][2:12]
	sgn[1] = raw[1][0]
	sgn[2] = raw[1][0]
	p1 += raw[2][2:12]
	sgn[3] = raw[2][0]
	sgn[4] = raw[2][0]
	p1 += raw[3][7:12]
	sgn[5] = raw[3][0]
	for i := 4; i < 9; i++ {
		p1 += raw[i][2:12]
		sgn[2*(i-4)+6] = raw[i][0]
		sgn[2*(i-4)+7] = raw[i][0]
	}
	p2 := ""
	st := 0
	ed := 0
	comb := 0
	for fld := 0; fld < 16; fld++ {
		if comb == 8 {
			comb = 24
		}
		if prtsw[comb] == 0 || fld == 15 {
			ed = (fld + 1) * 5
			if sgn[fld] == 'P' {
				p2 += p1[st:ed]
			} else {
				p2 += sm2tenc(p1[st:ed])
			}
			st = ed
		}
		comb++
	}
	s = ""
	for fld := 0; fld < 16; fld++ {
		if prtsw[fld+8] == 0 {
			s += blank
		} else {
			s += p2[5*fld : 5*(fld+1)]
		}
	}
	return
}

func sm2tenc(s string) string {
	var nz int

	l := len(s)
	for nz = l - 1; nz >= 0 && s[nz] == '0'; nz-- {
	}
	if nz < 0 { // negative 0 is still 0
		return s
	} else if nz == 0 { // special case for 10's comp and 11-punch
		return string('9'+1-s[0]-1+'J') + s[1:]
	} else {
		sc := string('9' - s[0] - 1 + 'J')
		if sc == "I" {
			sc = "-"
		}
		for i := 1; i < nz; i++ {
			sc += string('9' - s[i] + '0')
		}
		sc += string('9' + 1 - s[nz] + '0')
		sc += s[nz+1:]
		return sc
	}
}

func prctl(ch chan [2]string) {
	for {
		ctl := <-ch
		n := strings.IndexRune(ctl[0], '-')
		if n == -1 {
			sw, _ := strconv.Atoi(ctl[0])
			if ctl[1][0] == 'p' || ctl[1][0] == 'P' {
				prtsw[sw+7] = 1
			} else {
				prtsw[sw+7] = 0
			}
		} else {
			sw, _ := strconv.Atoi(ctl[0][:n])
			offset := 0
			if sw >= 9 {
				offset = 16
			}
			if ctl[1] == "c" || ctl[1] == "C" {
				prtsw[sw-1+offset] = 1
			} else {
				prtsw[sw-1+offset] = 0
			}
		}
	}
}
