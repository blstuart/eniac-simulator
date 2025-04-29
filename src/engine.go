package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	Handshake = iota
	Timed
)

var engineproc *os.Process
var engout io.WriteCloser
var ppunch chan string
var astat [20]string
var mpdstat [20]byte
var mpsstat [10]byte
var mpistat [10]byte
var ftarg [3]string
var ftring [3]string
var ftadd [3]string
var ftsub [3]string
var ftsetup [3]string
var conslast string
var mullast string
var dsqlast string

func engineshutdown() {
	if(engineproc == nil) {
		return
	}
	engout.Close()
	time.Sleep(100 * time.Millisecond)
	engineproc.Kill()
}

func enginedisplay(engcmd *string) {
	var runmode int
	var engmsg string

	cmlast := -1
	argv := strings.Split(*engcmd, " ")
	cmd := exec.Command(argv[0])
	if cmd.Err != nil {
		fmt.Fprintf(os.Stderr, "visualizion: %s\n", cmd.Err.Error())
		os.Exit(1)
	}
	cmd.Args = argv
	engout, _ = cmd.StdinPipe()
	engin, _ := cmd.StdoutPipe()
	cmd.Stderr = os.Stderr
	cmd.Start()
	engineproc = cmd.Process

	ppunch = make(chan string)
	go func() {
		for {
			card := <-ppunch
			if card == "exit" {
				break
			}
			fmt.Fprintf(engout, "punch %s\n", card)
		}
	} ()

	cyclast := 0
	for i := 0; i < 20; i++ {
		astat[i] = "P 0000000000 0000000000 0 000000000000"
	}
	for i := 0; i < 20; i++ {
		mpdstat[i] = 0
	}
	for i := 0; i < 10; i++ {
		mpsstat[i] = 0
		mpistat[i] = 0
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

	es := bufio.NewScanner(engin)
	for {
		es.Scan()
		engmsg = es.Text()
		if engmsg == "ready" {
			runmode = Timed
			break
		} else if engmsg == "update" {
			runmode = Handshake
			break
		}
	}

	handchan := make(chan int)
	go func() {
		for {
			if !es.Scan() {
				break
			}
			engmsg = es.Text()
			if engmsg == "exit" {
				break
			} else if engmsg == "update" {
				handchan <- 1
			} else if engmsg == "refresh" {
				cmlast = -1
				cyclast = 0
				for i := 0; i < 20; i++ {
					astat[i] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
				}
				for i := 0; i < 20; i++ {
					mpdstat[i] = 10
				}
				for i := 0; i < 10; i++ {
					mpsstat[i] = 10
					mpistat[i] = 10
				}
				for i := 0; i < 3; i++ {
					ftarg[i] = "x"
					ftring[i] = "x"
					ftadd[i] = "x"
					ftsub[i] = "x"
					ftsetup[i] = "x"
				}
				conslast = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
				mullast = "xx000000000000000000000000xxxx"
				dsqlast = "xxxx00000000x000000000000000000000000000000"
			} else {
				proccmd(engmsg)
			}
		}
	} ()

	go func() {
		for {
			needupd := false
			for i := 0; i < 20; i++ {
				s := accstat(i)
				if s != astat[i] {
					if s[0] == 'P' && astat[i][0] != 'P' {
						fmt.Fprintf(engout, "ad %d 0 0\n", i)
						needupd = true
					} else if s[0] == 'M' && astat[i][0] != 'M' {
						fmt.Fprintf(engout, "ad %d 0 3\n", i)
						needupd = true
					}
					for j := 0; j < 10; j++ {
						if s[j+2] != astat[i][j+2] {
							fmt.Fprintf(engout, "ad %d %d %c\n", i, j + 1, s[j+2])
							needupd = true
						}
						if s[j+13] != astat[i][j+13] {
							fmt.Fprintf(engout, "ac %d %d %c\n", i, j + 1, s[j+13])
							needupd = true
						}
					}
					if s[24] != astat[i][24] {
						fmt.Fprintf(engout, "ar %d %c\n", i, s[24])
						needupd = true
					}
					for j := 0; j < 12; j++ {
						if s[j+26] != astat[i][j+26] {
							fmt.Fprintf(engout, "af %d %d %c\n", i, j, s[j + 26])
							needupd = true
						}
					}
					astat[i] = s
				}
			}
			s := cycstat()
			n, _ := strconv.Atoi(s)
			if n != cyclast {
				fmt.Fprintf(engout, "cy %d\n", n)
				needupd = true
				cyclast = n
			}
			if cmlast != cmode {
				cmlast = cmode
				switch cmlast {
				case Pulse:
					fmt.Fprintf(engout, "cm P\n")
				case Add:
					fmt.Fprintf(engout, "cm A\n")
				case Cont:
					fmt.Fprintf(engout, "cm C\n")
				}
				needupd = true
			}
			s = mpstat()
			for i := 0; i < 10; i++ {
				b := s[i] - '0'
				if mpsstat[i] != b {
					fmt.Fprintf(engout, "mps %d %d\n", i, b)
					needupd = true
					mpsstat[i] = b
				}
			}
			for i := 0; i < 20; i++ {
				b := s[i+11] - '0'
				if mpdstat[i] != b {
					fmt.Fprintf(engout, "mpd %d %d\n", i, b)
					needupd = true
					mpdstat[i] = b
				}
			}
			for i := 0; i < 10; i++ {
				b := s[i+32] - '0'
				if mpistat[i] != b {
					fmt.Fprintf(engout, "mpi %d %d\n", i, b)
					needupd = true
					mpistat[i] = b
				}
			}
			for i := 0; i < 3; i++ {
				s = ftstat(i)
				ftinfo := strings.Split(s, " ")
				if ftinfo[1] != ftarg[i] {
					fmt.Fprintf(engout, "ftar %d %s\n", i, ftinfo[1])
					needupd = true
					ftarg[i] = ftinfo[1]
				}
				if ftinfo[2] != ftring[i] {
					fmt.Fprintf(engout, "ftr %d %s\n", i, ftinfo[2])
					needupd = true
					ftring[i] = ftinfo[2]
				}
				if ftinfo[3] != ftadd[i] {
					fmt.Fprintf(engout, "ftad %d %s\n", i, ftinfo[3])
					needupd = true
					ftadd[i] = ftinfo[3]
				}
				if ftinfo[4] != ftsub[i] {
					fmt.Fprintf(engout, "ftsu %d %s\n", i, ftinfo[4])
					needupd = true
					ftsub[i] = ftinfo[4]
				}
				if ftinfo[5] != ftsetup[i] {
					fmt.Fprintf(engout, "ftse %d %s\n", i, ftinfo[5])
					needupd = true
					ftsetup[i] = ftinfo[5]
				}
			}
			s = consstat()
			if s != conslast {
				for i := 0; i < 30; i++ {
					if s[i] != conslast[i] {
						fmt.Fprintf(engout, "ct %d %c\n", i, s[i])
						needupd = true
					}
				}
				conslast = s
			}
			s = multstat()
			if s != mullast {
				mullast = s
				fmt.Fprintf(engout, "m %s\n", mullast)
				needupd = true
			}
			s = divsrstat()
			if s != dsqlast {
				dsqlast = s
				fmt.Fprintf(engout, "d %s\n", dsqlast)
				needupd = true
			}
			switch runmode {
			case Timed:
				time.Sleep(30 * time.Millisecond)
			case Handshake:
				if needupd {
					fmt.Fprintf(engout, "up\n")
					<- handchan
				}
			}
		}
	} ()
}
