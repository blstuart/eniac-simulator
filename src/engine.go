package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var engineproc *os.Process

func enginedisplay(engcmd *string) {
	var astat [20]string

	cmd := exec.Command(*engcmd)
	engout, _ := cmd.StdinPipe()
	cmd.Start()
	engineproc = cmd.Process

	for i := 0; i < 20; i++ {
		astat[i] = "P 0000000000 0000000000 0 0000000000"
	}
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
						fmt.Fprintf(engout, "ac %d %d %c\n", i, j, s[j+13])
					}
				}
				astat[i] = s
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
