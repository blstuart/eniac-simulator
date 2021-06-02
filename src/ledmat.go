package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var ledacc = []int{ 1, 2, 3, 4 }
var ledproc *os.Process

func leddisplay() {
	var astat [4]string

	cmd := exec.Command("ledmat")
	matpipe, _ := cmd.StdinPipe()
	cmd.Start()
	ledproc = cmd.Process
	fmt.Fprintf(matpipe, "c\n")
	for i := 0; i < 4; i++ {
		fmt.Fprintf(matpipe, "s %d 10 1\n", i*16+2)
		for j := 0; j < 10; j++ {
			fmt.Fprintf(matpipe, "s %d 10 1\n", i*16+3+j)
		}
		astat[i] = "P 0000000000 0000000000"
	}
	for {
		for i := 0; i < 4; i++ {
			s := accstat(ledacc[i] - 1)
			if s[0] != astat[i][0] {
				if s[0] == 'M' {
					fmt.Fprintf(matpipe, "s %d 10 0\n", i*16+2)
					fmt.Fprintf(matpipe, "s %d 7 1\n", i*16+2)
				} else {
					fmt.Fprintf(matpipe, "s %d 7 0\n", i*16+2)
					fmt.Fprintf(matpipe, "s %d 10 1\n", i*16+2)
				}
			}
			for j := 0; j < 10; j++ {
				if s[j+2] != astat[i][j+2] {
					fmt.Fprintf(matpipe, "s %d %d 1\n",
						i*16+3+j, 10 - (s[j+2]-'0'))
					fmt.Fprintf(matpipe, "s %d %d 0\n",
						i*16+3+j, 10 - (astat[i][j+2]-'0'))
				}
			}
			for j := 0; j < 10; j++ {
				if s[j+13] != astat[i][j+13] {
					if s[j+13] == '1' {
						fmt.Fprintf(matpipe, "s %d 14 1\n", i*16+3+j)
					} else {
						fmt.Fprintf(matpipe, "s %d 14 0\n", i*16+3+j)
					}
				}
			}
			astat[i] = s
		}
		time.Sleep(30 * time.Millisecond)
	}
}
