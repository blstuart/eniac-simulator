package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var engineproc *os.Process

func enginedisplay(engcmd *string) {
	var astat [20]string
	var mpdstat [20]byte
	var mpsstat [10]byte
	var ftarg [3]string
	var ftring [3]string
	var ftadd [3]string
	var ftsub [3]string
	var ftsetup [3]string
	var conslast string
	var engmsg string
	var mullast string
	var dsqlast string

	cmd := exec.Command(*engcmd)
	engout, _ := cmd.StdinPipe()
	engin, _ := cmd.StdoutPipe()
	cmd.Start()
	engineproc = cmd.Process

	cyclast := 0
	for i := 0; i < 20; i++ {
		astat[i] = "P 0000000000 0000000000 0 0000000000"
	}
	for i := 0; i < 20; i++ {
		mpdstat[i] = 0
	}
	for i := 0; i < 10; i++ {
		mpsstat[i] = 0
	}
	for i := 0; i < 3; i++ {
		ftarg[i] = "0"
		ftring[i] = "0"
		ftadd[i] = "0"
		ftsub[i] = "0"
		ftsetup[i] = "0"
	}
	conslast = "000000000000000000000000000000"
	mullast = "0 000000000000000000000000 0 0"
	dsqlast = "0 0 00000000 000000000000000000000000000000"

	for {
		n, _ := fmt.Fscanln(engin, &engmsg);
		if n == 1 {
			if engmsg == "ready" {
				break
			}
		}
	}

	go func() {
		for {
			for i := 0; i < 20; i++ {
				s := accstat(i)
				if s[:23] != astat[i][:23] {
					if s[0] == 'P' && astat[i][0] == 'M' {
						fmt.Fprintf(engout, "ad %d 0 0\n", i)
					} else if s[0] == 'M' && astat[i][0] == 'P' {
						fmt.Fprintf(engout, "ad %d 0 3\n", i)
					}
					for j := 0; j < 10; j++ {
						if s[j+2] != astat[i][j+2] {
							fmt.Fprintf(engout, "ad %d %d %c\n", i, j + 1, s[j+2])
						}
						if s[j+13] != astat[i][j+13] {
							fmt.Fprintf(engout, "ac %d %d %c\n", i, j + 1, s[j+13])
						}
					}
					astat[i] = s
				}
			}
			s := cycstat()
			n, _ := strconv.Atoi(s)
			if n != cyclast {
				fmt.Fprintf(engout, "cy %d\n", n)
				cyclast = n
			}
			s = mpstat()
			for i := 0; i < 10; i++ {
				b := s[i] - '0'
				if mpsstat[i] != b {
					fmt.Fprintf(engout, "mps %d %d\n", i, b)
					mpsstat[i] = b
				}
			}
			for i := 0; i < 20; i++ {
				b := s[i+11] - '0'
				if mpdstat[i] != b {
					fmt.Fprintf(engout, "mpd %d %d\n", i, b)
					mpdstat[i] = b
				}
			}
			for i := 0; i < 3; i++ {
				s = ftstat(i)
				ftinfo := strings.Split(s, " ")
				if ftinfo[1] != ftarg[i] {
					fmt.Fprintf(engout, "ftar %d %s\n", i, ftinfo[1])
					ftarg[i] = ftinfo[1]
				}
				if ftinfo[2] != ftring[i] {
					fmt.Fprintf(engout, "ftr %d %s\n", i, ftinfo[2])
					ftring[i] = ftinfo[2]
				}
				if ftinfo[3] != ftadd[i] {
					fmt.Fprintf(engout, "ftad %d %s\n", i, ftinfo[3])
					ftadd[i] = ftinfo[3]
				}
				if ftinfo[4] != ftsub[i] {
					fmt.Fprintf(engout, "ftsu %d %s\n", i, ftinfo[4])
					ftsub[i] = ftinfo[4]
				}
				if ftinfo[5] != ftsetup[i] {
					fmt.Fprintf(engout, "ftse %d %s\n", i, ftinfo[5])
					ftsub[i] = ftinfo[5]
				}
			}
			s = consstat()
			if s != conslast {
				for i := 0; i < 30; i++ {
					if s[i] != conslast[i] {
						fmt.Fprintf(engout, "ct %d %c\n", i, s[i])
					}
				}
				conslast = s
			}
			s = multstat()
			if s != mullast {
				mullast = s
				fmt.Fprintf(engout, "m %s\n", mullast)
			}
			s = divsrstat()
			if s != dsqlast {
				dsqlast = s
				fmt.Fprintf(engout, "d %s\n", dsqlast)
			}
			time.Sleep(30 * time.Millisecond)
		}
	} ()
}
