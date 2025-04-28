package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type gstat struct {
	guimode int
	proc *os.Process
}

var guistate gstat

var msgin *bufio.Scanner
var width, height int
var usecontrol, demomode, tkkludge *bool

var ftuoff = []int{4, 28, 30}
var accoff = []int{6 * 642, 7 * 642, 9 * 642, 10 * 642, 11 * 642, 12 * 642, 13 * 642, 14 * 642, 15 * 642,
	16 * 642, 20 * 642, 21 * 642, 22 * 642, 23 * 642,
	24 * 642, 25 * 642, 26 * 642, 27 * 642, 32 * 642, 33 * 642}

func loadmenu(gpipe io.Writer) {
	f, err := os.Open("programs")
	if err != nil {
		fmt.Println("program list: ", err)
		return
	}
	l, err := f.Readdirnames(0)
	f.Close()
	if err != nil {
		fmt.Println("program list: ", err)
		return
	}
	fmt.Fprintln(gpipe, "menu .mlib -tearoff false")
	p := sort.StringSlice(l)
	p.Sort()
	for _, n := range p {
		fmt.Fprintf(gpipe, ".mlib add command -label %s "+
			"-command {puts stdout \"l %s\"}\n", n, n)
	}
	fmt.Fprintln(gpipe, "tk_popup .mlib 10 10")
}

func resetbuts(gpipe io.Writer, newstate int) {
	switch guistate.guimode {
	case 0:
		fmt.Fprintln(gpipe, ".perspbut configure -relief raised")
	case 1:
		fmt.Fprintln(gpipe, ".s1but configure -relief raised")
	case 2:
		fmt.Fprintln(gpipe, ".s2but configure -relief raised")
	case 3:
		fmt.Fprintln(gpipe, ".s3but configure -relief raised")
	case 4:
		fmt.Fprintln(gpipe, ".s4but configure -relief raised")
	case 5:
		fmt.Fprintln(gpipe, ".s5but configure -relief raised")
	}
	switch newstate {
	case 0:
		fmt.Fprintln(gpipe, ".perspbut configure -relief sunken")
	case 1:
		fmt.Fprintln(gpipe, ".s1but configure -relief sunken")
	case 2:
		fmt.Fprintln(gpipe, ".s2but configure -relief sunken")
	case 3:
		fmt.Fprintln(gpipe, ".s3but configure -relief sunken")
	case 4:
		fmt.Fprintln(gpipe, ".s4but configure -relief sunken")
	case 5:
		fmt.Fprintln(gpipe, ".s5but configure -relief sunken")
	}
	guistate.guimode = newstate
}

func guicmd(sc *bufio.Scanner, gpipe io.Writer) {
	for sc.Scan() {
		s := sc.Text()
		if s == "loadmenu" {
			loadmenu(gpipe)
		} else if s == "s1view" {
			resetbuts(gpipe, 1)
			fmt.Fprintln(gpipe, ".eniac configure -image s1img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		} else if s == "s2view" {
			resetbuts(gpipe, 2)
			fmt.Fprintln(gpipe, ".eniac configure -image s2img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		} else if s == "s3view" {
			resetbuts(gpipe, 3)
			fmt.Fprintln(gpipe, ".eniac configure -image s3img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		} else if s == "s4view" {
			resetbuts(gpipe, 4)
			fmt.Fprintln(gpipe, ".eniac configure -image s4img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		} else if s == "s5view" {
			resetbuts(gpipe, 5)
			fmt.Fprintln(gpipe, ".eniac configure -image s5img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		} else if s == "perspview" {
			resetbuts(gpipe, 0)
			fmt.Fprintln(gpipe, ".eniac configure -image enimg")
			clearstate()
			setneonsize(gpipe, "neon")
			drawfixed(gpipe)
		} else if len(s) > 2 && s[:2] == "l " {
			fmt.Fprintln(gpipe, "destroy .mlib")
			fmt.Println(s)
		} else {
			fmt.Println(s)
		}
	}
}

func clearstate() {
	fmt.Println("refresh")
}

func setneonsize(gpipe io.Writer, img string) {
	fmt.Fprintln(gpipe, ".ih1 configure -image", img)
	fmt.Fprintln(gpipe, ".ih2 configure -image", img)
	for i := 0; i < 6; i++ {
		fmt.Fprintf(gpipe, ".initc%d configure -image %s\n", i+1, img)
	}
	fmt.Fprintln(gpipe, ".initrs configure -image", img)
	fmt.Fprintln(gpipe, ".initps configure -image", img)
	fmt.Fprintln(gpipe, ".initrf configure -image", img)
	fmt.Fprintln(gpipe, ".initri configure -image", img)
	fmt.Fprintln(gpipe, ".initrsy configure -image", img)
	fmt.Fprintln(gpipe, ".initpf configure -image", img)
	fmt.Fprintln(gpipe, ".initpsy configure -image", img)
	fmt.Fprintln(gpipe, ".initi configure -image", img)
	fmt.Fprintln(gpipe, ".initis configure -image", img)
	fmt.Fprintln(gpipe, ".cycst configure -image", img)
	fmt.Fprintln(gpipe, ".cy10p configure -image", img)
	fmt.Fprintln(gpipe, ".cyccg configure -image", img)
	fmt.Fprintln(gpipe, ".cych1 configure -image", img)
	fmt.Fprintln(gpipe, ".cych2 configure -image", img)
	fmt.Fprintln(gpipe, ".cych3 configure -image", img)
	fmt.Fprintln(gpipe, ".cych4 configure -image", img)
	for i := 1; i <= 20; i++ {
		fmt.Fprintf(gpipe, ".acc%dh1 configure -image %s\n", i, img)
		fmt.Fprintf(gpipe, ".acc%dh2 configure -image %s\n", i, img)
		fmt.Fprintf(gpipe, ".acc%ds configure -image %s\n", i, img)
		for j := 9; j >= 0; j-- {
			fmt.Fprintf(gpipe, ".a%dd%d configure -image %s\n", i, j, img)
			fmt.Fprintf(gpipe, ".aff%dd%d configure -image %s\n", i, j, img)
		}
		fmt.Fprintf(gpipe, ".acc%drep configure -image %s\n", i, img)
		for j := 0; j < 12; j++ {
			fmt.Fprintf(gpipe, ".acc%dff%d configure -image %s\n", i, j+1, img)
		}
	}
	fmt.Fprintln(gpipe, ".ph1 configure -image", img)
	fmt.Fprintln(gpipe, ".ph2 configure -image", img)
	fmt.Fprintln(gpipe, ".ph3 configure -image", img)
	fmt.Fprintln(gpipe, ".ph4 configure -image", img)
	for i := 0; i < 10; i++ {
		fmt.Fprintf(gpipe, ".pi%d configure -image %s\n", i, img)
		fmt.Fprintf(gpipe, ".ps%d configure -image %s\n", i, img)
	}
	for i := 0; i < 20; i++ {
		fmt.Fprintf(gpipe, ".pd%d configure -image %s\n", i, img)
	}
	fmt.Fprintln(gpipe, ".dh1 configure -image", img)
	fmt.Fprintln(gpipe, ".dh2 configure -image", img)
	fmt.Fprintln(gpipe, ".dplr configure -image", img)
	fmt.Fprintln(gpipe, ".dprr configure -image", img)
	for i := 0; i < 8; i++ {
		fmt.Fprintf(gpipe, ".dprogff%d configure -image %s\n", i+1, img)
	}
	fmt.Fprintln(gpipe, ".ddivff configure -image", img)
	fmt.Fprintln(gpipe, ".dclrff configure -image", img)
	fmt.Fprintln(gpipe, ".dilockff configure -image", img)
	fmt.Fprintln(gpipe, ".ddpgamma configure -image", img)
	fmt.Fprintln(gpipe, ".dngamma configure -image", img)
	fmt.Fprintln(gpipe, ".dpsrcff configure -image", img)
	fmt.Fprintln(gpipe, ".dpringff configure -image", img)
	fmt.Fprintln(gpipe, ".ddenomff configure -image", img)
	fmt.Fprintln(gpipe, ".dnumrplus configure -image", img)
	fmt.Fprintln(gpipe, ".dnumrmin configure -image", img)
	fmt.Fprintln(gpipe, ".dqalpha configure -image", img)
	fmt.Fprintln(gpipe, ".dsac configure -image", img)
	fmt.Fprintln(gpipe, ".dm2 configure -image", img)
	fmt.Fprintln(gpipe, ".dm1 configure -image", img)
	fmt.Fprintln(gpipe, ".dnac configure -image", img)
	fmt.Fprintln(gpipe, ".dda configure -image", img)
	fmt.Fprintln(gpipe, ".dnalpha configure -image", img)
	fmt.Fprintln(gpipe, ".ddalpha configure -image", img)
	fmt.Fprintln(gpipe, ".ddgamma configure -image", img)
	fmt.Fprintln(gpipe, ".dngamma configure -image", img)
	fmt.Fprintln(gpipe, ".dp2 configure -image", img)
	fmt.Fprintln(gpipe, ".dp1 configure -image", img)
	fmt.Fprintln(gpipe, ".dsalpha configure -image", img)
	fmt.Fprintln(gpipe, ".dds configure -image", img)
	fmt.Fprintln(gpipe, ".dnbeta configure -image", img)
	fmt.Fprintln(gpipe, ".ddbeta configure -image", img)
	fmt.Fprintln(gpipe, ".dans1 configure -image", img)
	fmt.Fprintln(gpipe, ".dans2 configure -image", img)
	fmt.Fprintln(gpipe, ".dans3 configure -image", img)
	fmt.Fprintln(gpipe, ".dans4 configure -image", img)
	for i := 0; i < 3; i++ {
		fmt.Fprintf(gpipe, ".mh%d configure -image %s\n", 2*i, img)
		fmt.Fprintf(gpipe, ".mh%d configure -image %s\n", 2*i+1, img)
	}
	for i := 0; i < 24; i++ {
		fmt.Fprintf(gpipe, ".mi%d configure -image %s\n", i+1, img)
	}
	fmt.Fprintln(gpipe, ".mr1 configure -image", img)
	fmt.Fprintln(gpipe, ".mr3 configure -image", img)
	fmt.Fprintln(gpipe, ".mstage configure -image", img)
	fmt.Fprintln(gpipe, ".conh1 configure -image", img)
	fmt.Fprintln(gpipe, ".conh2 configure -image", img)
	fmt.Fprintln(gpipe, ".conh3 configure -image", img)
	fmt.Fprintln(gpipe, ".conh4 configure -image", img)
	fmt.Fprintln(gpipe, ".prh1 configure -image", img)
	fmt.Fprintln(gpipe, ".prh2 configure -image", img)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fprintf(gpipe, ".ft%dh%d configure -image %s\n", i+1, j+1, img)
		}
		fmt.Fprintf(gpipe, ".ft%da1 configure -image %s\n", i+1, img)
		fmt.Fprintf(gpipe, ".ft%da10 configure -image %s\n", i+1, img)
		fmt.Fprintf(gpipe, ".ft%dr configure -image %s\n", i+1, img)
		fmt.Fprintf(gpipe, ".ft%daset configure -image %s\n", i+1, img)
		fmt.Fprintf(gpipe, ".ft%dadd configure -image %s\n", i+1, img)
		fmt.Fprintf(gpipe, ".ft%dsubt configure -image %s\n", i+1, img)
		for j := 0; j < 12; j++ {
			fmt.Fprintf(gpipe, ".ft%dt%d configure -image %s\n", i+1, j+1, img)
		}
	}
	for i := 0; i < 30; i++ {
		fmt.Fprintf(gpipe, ".ct%d configure -image %s\n", i+1, img)
	}
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	usecontrol = flag.Bool("c", false, "use a portable control station connected to GPIO pins")
	demomode = flag.Bool("D", false, "cycle among perspectives")
	tkkludge = flag.Bool("K", false, "work around memory issue in some Tk implementations")
	flag.Parse()

	msgin = bufio.NewScanner(os.Stdin)
	for {
		innergui()
	}
}

func findimages() string {
	_, e := os.Stat("images")
	if e == nil {
		return ""
	}
	_, e = os.Stat("../images")
	if e == nil {
		return "../"
	}
	home, ex := os.LookupEnv("home")
	if !ex  {
		home, ex = os.LookupEnv("HOME")
	}
	if ex {
		_, e = os.Stat(home + "lib/eniac/images")
		if e != nil {
			return home + "lib/eniac/"
		}
	}
	return ""
}

func innergui() {
	var nname string
	var x, y int

	guistate.guimode = 0
	widset := []int{480, 720, 1020, 1280, 1360, 1600, 1800, 1900, 2500, 3500}
	cmd := exec.Command("wish")
	gpipe, _ := cmd.StdinPipe()
	cpipe, _ := cmd.StdoutPipe()
	cmd.Start()
	guistate.proc = cmd.Process
	fmt.Fprintln(gpipe, "wm geometry . +5+5")
	sc := bufio.NewScanner(cpipe)
	fmt.Fprintln(gpipe, "puts [winfo vrootwidth .]")
	sc.Scan()
	w, _ := strconv.Atoi(sc.Text())
	fmt.Fprintln(gpipe, "puts [winfo vrootheight .]")
	sc.Scan()
	h, _ := strconv.Atoi(sc.Text())
	if w/2+120 > h {
		w = (h - 120) * 2
	}
	width = 0
	for i, wid := range widset {
		if w < wid && i > 0 {
			width = widset[i-1]
			break
		}
	}
	if width == 0 {
		width = 1900
	}
	height = width / 2
	fmt.Fprintln(gpipe, "fconfigure stdout -buffering none -blocking false")
	go guicmd(sc, gpipe)
	if *demomode {
		go rundemo(gpipe)
	}
	if width > 480 {
		fmt.Fprintf(gpipe, "wm geometry . %dx%d\n", width, height+120)
	} else {
		fmt.Fprintf(gpipe, "wm geometry . %dx%d\n", width, height+32)
	}
	prefix := findimages()
	fmt.Fprintf(gpipe, "image create photo enimg -file %simages/e%d.ppm\n", prefix, width)
	fmt.Fprintf(gpipe, "image create photo s1img -file %simages/e%ds1.ppm\n", prefix, width)
	fmt.Fprintf(gpipe, "image create photo s2img -file %simages/e%ds2.ppm\n", prefix, width)
	fmt.Fprintf(gpipe, "image create photo s3img -file %simages/e%ds3.ppm\n", prefix, width)
	fmt.Fprintf(gpipe, "image create photo s4img -file %simages/e%ds4.ppm\n", prefix, width)
	fmt.Fprintf(gpipe, "image create photo s5img -file %simages/e%ds5.ppm\n", prefix, width)
	fmt.Fprintln(gpipe, "label .eniac -borderwidth 0 -image enimg")
	fmt.Fprintln(gpipe, "place .eniac -x 0 -y 0")
	fmt.Fprintf(gpipe, "image create photo ctlimg -file %simages/control.ppm\n", prefix)
	fmt.Fprintln(gpipe, "label .control -borderwidth 2 -image ctlimg")
	fmt.Fprintln(gpipe, "label .cmode -font [list Clean 10] -borderwidth 1 -width 7 -height 1")
	butparam := "-borderwidth 0 -font [list Clean 10] -pady 0"
	if runtime.GOOS == "darwin" {
		butparam += " -padx 10"
	} else {
		butparam += " -padx 1"
	}
	fmt.Fprintln(gpipe, "button .copc "+butparam+" -text Co -command {puts stdout \"s cy.op co\"}")
	fmt.Fprintln(gpipe, "button .copp "+butparam+" -text 1P -command {puts stdout \"s cy.op 1p\"}")
	fmt.Fprintln(gpipe, "button .copa "+butparam+" -text 1A -command {puts stdout \"s cy.op 1a\"}")
	fmt.Fprintln(gpipe, "button .clbut "+butparam+" -text CLEAR -command {puts stdout \"b c\"}")
	fmt.Fprintln(gpipe, "button .readbut "+butparam+" -text READ -command {puts stdout \"b r\"}")
	fmt.Fprintln(gpipe, "button .initbut "+butparam+" -text INIT -command {puts stdout \"b i\"}")
	fmt.Fprintln(gpipe, "button .pulbut "+butparam+" -text PULSE -command {puts stdout \"b p\"}")
	if !*usecontrol {
		if width > 480 {
			fmt.Fprintf(gpipe, "place .control -x %d -y %d\n", width/12, height-82)
			fmt.Fprintf(gpipe, "place .cmode -x %d -y %d\n",
				width/12+12, height-82+5)
			fmt.Fprintf(gpipe, "place .copc -x %d -y %d\n",
				width/12+5, height-82+27)
			fmt.Fprintf(gpipe, "place .copp -x %d -y %d\n",
				width/12+30, height-82+27)
			fmt.Fprintf(gpipe, "place .copa -x %d -y %d\n",
				width/12+55, height-82+27)
			fmt.Fprintf(gpipe, "place .clbut -x %d -y %d\n",
				width/12+5, height-82+53)
			fmt.Fprintf(gpipe, "place .readbut -x %d -y %d\n",
				width/12+5, height-82+83)
			fmt.Fprintf(gpipe, "place .initbut -x %d -y %d\n",
				width/12+5, height-82+113)
			fmt.Fprintf(gpipe, "place .pulbut -x %d -y %d\n",
				width/12+5, height-82+143)
		} else {
			fmt.Fprintf(gpipe, "place .cmode -x %d -y %d\n", width/10+5, height-78)
			fmt.Fprintf(gpipe, "place .copc -x %d -y %d\n", width/10+5, height-65)
			fmt.Fprintf(gpipe, "place .copp -x %d -y %d\n", width/10+30, height-65)
			fmt.Fprintf(gpipe, "place .copa -x %d -y %d\n", width/10+55, height-65)
			fmt.Fprintf(gpipe, "place .clbut -x %d -y %d\n", width/10+40, height-52)
			fmt.Fprintf(gpipe, "place .readbut -x %d -y %d\n", width/10+40, height-31)
			fmt.Fprintf(gpipe, "place .initbut -x %d -y %d\n", width/10+40, height-10)
			fmt.Fprintf(gpipe, "place .pulbut -x %d -y %d\n", width/10+40, height+11)
		}
	}
	butparam = "-borderwidth 2 -font [list Clean 10] -width 5 -pady 0"
	if runtime.GOOS == "darwin" {
		butparam += " -padx 8"
	} else {
		butparam += " -padx 2"
	}
	fmt.Fprintln(gpipe, "button .resetbut "+butparam+" -text RESET "+
		"-command {puts stdout \"R\"}\n")
	fmt.Fprintln(gpipe, "button .loadbut "+butparam+
		" -text \"LOAD\\nCONF\" -command {puts stdout \"loadmenu\"}\n")
	if width > 480 {
		fmt.Fprintf(gpipe, "place .resetbut -anchor c -x %d -y %d\n", width/24, height+25)
		fmt.Fprintf(gpipe, "place .loadbut -anchor c -x %d -y %d\n", width/24, height+70)
	} else {
		fmt.Fprintf(gpipe, "place .resetbut -x 3 -y %d\n", height+11)
		fmt.Fprintf(gpipe, "place .loadbut -x 3 -y %d\n", height-30)
	}
	fmt.Fprintf(gpipe, "button .s1but -borderwidth 2 -pady 0 -font [list Clean 10] "+
		"-text \"S1 View\" -command {puts stdout \"s1view\"}\n")
	fmt.Fprintf(gpipe, "button .s2but -borderwidth 2 -pady 0 -font [list Clean 10] "+
		"-text \"S2 View\" -command {puts stdout \"s2view\"}\n")
	fmt.Fprintf(gpipe, "button .s3but -borderwidth 2 -pady 0 -font [list Clean 10] "+
		"-text \"S3 View\" -command {puts stdout \"s3view\"}\n")
	fmt.Fprintf(gpipe, "button .s4but -borderwidth 2 -pady 0 -font [list Clean 10] "+
		"-text \"S4 View\" -command {puts stdout \"s4view\"}\n")
	fmt.Fprintf(gpipe, "button .s5but -borderwidth 2 -pady 0 -font [list Clean 10] "+
		"-text \"S5 View\" -command {puts stdout \"s5view\"}\n")
	fmt.Fprintf(gpipe, "button .perspbut -borderwidth 2 -pady 0 -font [list Clean 10] "+
		"-text \"Perspective View\" -relief sunken -command {puts stdout \"perspview\"}\n")
	if width > 480 {
		if *usecontrol {
			fmt.Fprintf(gpipe, "place .s1but -anchor c -x %d -y %d\n",
				width/8, height-48)
			fmt.Fprintf(gpipe, "place .s2but -anchor c -x %d -y %d\n",
				width/3, height-48)
		} else {
			fmt.Fprintf(gpipe, "place .s1but -anchor c -x %d -y %d\n",
				width/12+130, height-48)
			if width > 720 {
				fmt.Fprintf(gpipe, "place .s2but -anchor c -x %d -y %d\n",
					width/3, height-48)
			} else {
				fmt.Fprintf(gpipe, "place .s2but -anchor c -x %d -y %d\n",
					width/12+215, height-48)
			}
		}
		fmt.Fprintf(gpipe, "place .s3but -anchor c -x %d -y %d\n", width/2, height-48)
		fmt.Fprintf(gpipe, "place .s4but -anchor c -x %d -y %d\n", 2*width/3, height-48)
		fmt.Fprintf(gpipe, "place .s5but -anchor c -x %d -y %d\n", 7*width/8, height-48)
		fmt.Fprintf(gpipe, "place .perspbut -anchor c -x %d -y %d\n", width/2, height-79)
	} else {
		fmt.Fprintf(gpipe, "place .s1but -x %d -y %d\n", width/10+100, height-73)
		fmt.Fprintf(gpipe, "place .s2but -x %d -y %d\n", width/10+100, height-52)
		fmt.Fprintf(gpipe, "place .s3but -x %d -y %d\n", width/10+100, height-31)
		fmt.Fprintf(gpipe, "place .s4but -x %d -y %d\n", width/10+100, height-10)
		fmt.Fprintf(gpipe, "place .s5but -x %d -y %d\n", width/10+100, height+11)
		fmt.Fprintf(gpipe, "place .perspbut -x %d -y %d\n", width/10+182, height+11)
	}
	fmt.Fprintf(gpipe, "image create photo neon -file %simages/orange.ppm\n", prefix)
	if width >= 1600 {
		fmt.Fprintf(gpipe, "image create photo neon2 -file %simages/orange3.ppm\n", prefix)
	} else {
		fmt.Fprintf(gpipe, "image create photo neon2 -file %simages/orange2.ppm\n", prefix)
	}
	fmt.Fprintf(gpipe, "image create photo gpilot1 -file %simages/gpilot.ppm\n", prefix)
	fmt.Fprintf(gpipe, "image create photo apilot1 -file %simages/apilot.ppm\n", prefix)
	// Initiating unit
	fmt.Fprintln(gpipe, "image create photo gpilot2")
	fmt.Fprintf(gpipe, "gpilot2 copy gpilot1 -subsample %d %d\n", 16000/width, 13000/width)
	fmt.Fprintln(gpipe, "label .gpilot -image gpilot2 -borderwidth 0")
	fmt.Fprintln(gpipe, "image create photo apilot2")
	fmt.Fprintf(gpipe, "apilot2 copy apilot1 -subsample %d %d\n", 16000/width, 13000/width)
	fmt.Fprintln(gpipe, "label .apilot -image apilot2 -borderwidth 0")
	fmt.Fprintln(gpipe, "label .ih1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ih2 -borderwidth 0 -image neon")
	for i := 0; i < 6; i++ {
		fmt.Fprintf(gpipe, "label .initc%d -borderwidth 0 -image neon\n", i+1)
	}
	fmt.Fprintln(gpipe, "label .initrs -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initps -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initrf -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initri -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initrsy -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initpf -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initpsy -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initi -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .initis -borderwidth 0 -image neon")
	// Cycling unit
	fmt.Fprintln(gpipe, "label .cycst -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .cy10p -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .cyccg -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .cych1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .cych2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .cych3 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .cych4 -borderwidth 0 -image neon")
	// Accumulators
	for i := 1; i <= 20; i++ {
		fmt.Fprintf(gpipe, "label .acc%dh1 -borderwidth 0 -image neon\n", i)
		fmt.Fprintf(gpipe, "label .acc%dh2 -borderwidth 0 -image neon\n", i)
		fmt.Fprintf(gpipe, "label .acc%ds -borderwidth 0 -image neon\n", i)
		x, y, _ = ray(accoff[i-1]+59, 149, 1)
		fmt.Fprintf(gpipe, "place .acc%ds -x %d -y %d\n", i, x, y)
		for j := 9; j >= 0; j-- {
			fmt.Fprintf(gpipe, "label .a%dd%d -borderwidth 0 -image neon\n", i, j)
			x, y = neonpos(i - 1, j, 0, accoff)
			x, y, _ = ray(x, y, 1)
			fmt.Fprintf(gpipe, "place .a%dd%d -x %d -y %d\n", i, j, x, y)
			fmt.Fprintf(gpipe, "label .aff%dd%d -borderwidth 0 -image neon\n", i, j)
		}
		fmt.Fprintf(gpipe, "label .acc%drep -borderwidth 0 -image neon\n", i)
		x, y, _ = ray(accoff[i-1] + 90, -1310, 1)
		fmt.Fprintf(gpipe, "place .acc%drep -x %d -y %d\n", i, x, y)
		for j := 0; j < 12; j++ {
			fmt.Fprintf(gpipe, "label .acc%dff%d -borderwidth 0 -image neon\n", i, j+1)
		}
	}
	// Master programmer
	fmt.Fprintln(gpipe, "label .ph1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ph2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ph3 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ph4 -borderwidth 0 -image neon")
	for i := 0; i < 10; i++ {
		fmt.Fprintf(gpipe, "label .pi%d -borderwidth 0 -image neon\n", i)
		fmt.Fprintf(gpipe, "label .ps%d -borderwidth 0 -image neon\n", i)
		if i < 5 {
			x, y, _ = ray(158+73*i+2*642, -223, 1)
			fmt.Fprintf(gpipe, "place .ps%d -x %d -y %d\n", i, x, y)
		} else {
			x, y, _ = ray(758+73*(i-5)+2*642, -223, 1)
			fmt.Fprintf(gpipe, "place .ps%d -x %d -y %d\n", i, x, y)
		}
	}
	for i := 0; i < 20; i++ {
		fmt.Fprintf(gpipe, "label .pd%d -borderwidth 0 -image neon\n", i)
		if i < 10 {
			x, y, _ = ray(160+37*i+2*642, 191, 1)
			fmt.Fprintf(gpipe, "place .pd%d -x %d -y %d\n", i, x, y)
		} else {
			x, y, _ = ray(762+37*(i-10)+2*642, 191, 1)
			fmt.Fprintf(gpipe, "place .pd%d -x %d -y %d\n", i, x, y)
		}
	}
	// Divider
	fmt.Fprintln(gpipe, "label .dh1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dh2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dplr -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dprr -borderwidth 0 -image neon")
	for i := 0; i < 8; i++ {
		fmt.Fprintf(gpipe, "label .dprogff%d -borderwidth 0 -image neon\n", i+1)
	}
	fmt.Fprintln(gpipe, "label .ddivff -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dclrff -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dilockff -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ddpgamma -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dngamma -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dpsrcff -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dpringff -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ddenomff -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dnumrplus -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dnumrmin -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dqalpha -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dsac -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dm2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dm1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dnac -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dda -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dnalpha -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ddalpha -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ddgamma -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dnpgamma -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dp2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dp1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dsalpha -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dds -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dnbeta -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .ddbeta -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dans1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dans2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dans3 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .dans4 -borderwidth 0 -image neon")
	// Multiplier
	for i := 0; i < 3; i++ {
		fmt.Fprintf(gpipe, "label .mh%d -borderwidth 0 -image neon\n", 2*i)
		fmt.Fprintf(gpipe, "label .mh%d -borderwidth 0 -image neon\n", 2*i+1)
	}
	for i := 0; i < 24; i++ {
		fmt.Fprintf(gpipe, "label .mi%d -borderwidth 0 -image neon\n", i+1)
	}
	fmt.Fprintln(gpipe, "label .mr1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .mr3 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .mstage -borderwidth 0 -image neon")
	// Constant transmitter
	fmt.Fprintln(gpipe, "label .conh1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .conh2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .conh3 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .conh4 -borderwidth 0 -image neon")
	// Printer/punch
	fmt.Fprintln(gpipe, "label .prh1 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "label .prh2 -borderwidth 0 -image neon")
	fmt.Fprintln(gpipe, "text .outdeck -width 80 -height 10")
	if width > 480 {
		fmt.Fprintf(gpipe, "place .outdeck -x %d -y %d\n", width-570, height-30)
	}
	lines := 0
	// Function tables
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Fprintf(gpipe, "label .ft%dh%d -borderwidth 0 -image neon\n", i+1, j+1)
		}
		fmt.Fprintf(gpipe, "label .ft%da1 -borderwidth 0 -image neon\n", i+1)
		x, y, _ = ray(88+ftuoff[i]*642, 40, 1)
		fmt.Fprintf(gpipe, "place .ft%da1 -x %d -y %d\n", i+1, x, y)
		fmt.Fprintf(gpipe, "label .ft%da10 -borderwidth 0 -image neon\n", i+1)
		x, y, _ = ray(347+ftuoff[i]*642, 40, 1)
		fmt.Fprintf(gpipe, "place .ft%da10 -x %d -y %d\n", i+1, x, y)
		fmt.Fprintf(gpipe, "label .ft%dr -borderwidth 0 -image neon\n", i+1)
		x, y, _ = ray(268+ftuoff[i]*642, -25, 1)
		fmt.Fprintf(gpipe, "place .ft%dr -x %d -y %d\n", i+1, x, y)
		fmt.Fprintf(gpipe, "label .ft%daset -borderwidth 0 -image neon\n", i+1)
		fmt.Fprintf(gpipe, "label .ft%dadd -borderwidth 0 -image neon\n", i+1)
		fmt.Fprintf(gpipe, "label .ft%dsubt -borderwidth 0 -image neon\n", i+1)
		for j := 0; j < 12; j++ {
			fmt.Fprintf(gpipe, "label .ft%dt%d -borderwidth 0 -image neon\n", i+1, j+1)
		}
	}
	// Constant transmitter
	for i := 0; i < 30; i++ {
		fmt.Fprintf(gpipe, "label .ct%d -borderwidth 0 -image neon\n", i+1)
	}

	drawfixed(gpipe)

	fmt.Println("update")

	// Kludge to restart wish periodically to avoid
	// it running out of memory.  Only seems to be
	// an issue on Linux/arm version of wish
	updatecnt := 0
	if *tkkludge {
		updatecnt = 1
	}
mainloop:
	for i := 0; i < 10000; i += updatecnt {
		if !msgin.Scan() {
			guistate.proc.Kill()
			os.Exit(0)
		}
		s := msgin.Text()
		p := strings.Split(s, " ")
		switch p[0] {
		case "exit":
			break mainloop
		case "up":
			fmt.Fprintln(gpipe, "update")
			fmt.Fprintln(gpipe, "puts \"update\"")

		// Initiating unit
		case "init":
			for i, f := range s[:6] {
				nname = fmt.Sprintf(".initc%d", i+1)
				neonplcl(gpipe, nname, f == '1', 90+45*i, -1150)
			}
			neonplcl(gpipe, ".initrs", s[6] == '1', 360, -1150)
			neonplcl(gpipe, ".initps", s[7] == '1', 360, -1160)
			neonplcl(gpipe, ".initrf", s[8] == '1', 405, -1150)
			neonplcl(gpipe, ".initri", s[9] == '1', 405, -1160)
			neonplcl(gpipe, ".initrsy", s[10] == '1', 450, -1150)
			neonplcl(gpipe, ".initpf", s[11] == '1', 495, -1150)
			neonplcl(gpipe, ".initpsy", s[12] == '1', 495, -1160)
			neonplcl(gpipe, ".initi", s[13] == '1', 540, -1150)
			neonplcl(gpipe, ".initis", s[14] == '1', 540, -1160)

		// Cycle unit
		case "cy":
			n, _ := strconv.Atoi(p[1])
			neonplcl(gpipe, ".cycst", true, 122+642+(n/2)*20, 40)
			neonplcl(gpipe, ".cyccg", n >= 22 && n <= 36, 1010, -18)
			neonplcl(gpipe, ".cy10p", n < 20 && n%2 == 1, 884, -18)
		case "cm":
			switch p[1] {
			case "P":
				fmt.Fprintln(gpipe, ".cmode configure -text \"1 Pulse\"")
			case "A":
				fmt.Fprintln(gpipe, ".cmode configure -text \"1 Add\"")
			case "C":
				fmt.Fprintln(gpipe, ".cmode configure -text \"Cont.\"")
			}

		// Accumulators
		case "ad":
			unit, _ := strconv.Atoi(p[1])
			dig, _ := strconv.Atoi(p[2])
			d, _ := strconv.Atoi(p[3])
			if dig == 0 {
				neonplcl(gpipe, fmt.Sprintf(".acc%ds", unit + 1), true, accoff[unit]+59, 149 + 38 * d)
			} else {
				x, y = neonpos(unit, dig, d, accoff)
				neonplcl(gpipe, fmt.Sprintf(".a%dd%d", unit + 1, dig - 1), true, x, y)
			}
		case "ac":
			unit, _ := strconv.Atoi(p[1])
			dig, _ := strconv.Atoi(p[2])
			d, _ := strconv.Atoi(p[3])
			neonplcl(gpipe, fmt.Sprintf(".aff%dd%d", unit + 1, dig), d == 1,
				accoff[unit]+dig*49+63, -31)
		case "ar":
			unit, _ := strconv.Atoi(p[1])
			rep, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, fmt.Sprintf(".acc%drep", unit), true, accoff[unit]+90, -1310+20*rep)
		case "af":
			unit, _ := strconv.Atoi(p[1])
			prog, _ := strconv.Atoi(p[2])
			val, _ := strconv.Atoi(p[3])
			switch prog {
			case 0:
				nname = fmt.Sprintf(".acc%dff1", unit + 1)
				neonplcl(gpipe, nname, val == 1, accoff[unit]+135, -1130)
			case 1:
				nname = fmt.Sprintf(".acc%dff2", unit + 1)
				neonplcl(gpipe, nname, val == 1, accoff[unit]+135, -1150)
			case 2:
				nname = fmt.Sprintf(".acc%dff3", unit + 1)
				neonplcl(gpipe, nname, val == 1, accoff[unit]+180, -1130)
			case 3:
				nname = fmt.Sprintf(".acc%dff4", unit + 1)
				neonplcl(gpipe, nname, val == 1, accoff[unit]+180, -1150)
			default:
				nname = fmt.Sprintf(".acc%dff%d", unit + 1, prog + 1)
				neonplcl(gpipe, nname, val == 1, accoff[unit]+225+45*(prog - 4), -1130)
			}

		// Divider/Square Rooter
		case "d":
			plring, _ := strconv.Atoi(p[1])
			prring, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, ".dplr", true, 8*642+88, -320+15*plring)
			neonplcl(gpipe, ".dprr", true, 8*642+534, -1346+15*prring)
			neonplcl(gpipe, ".dprogff1", p[3][0] == '1', 8*642+94, -1156)
			neonplcl(gpipe, ".dprogff2", p[3][1] == '1', 8*642+139, -1156)
			neonplcl(gpipe, ".dprogff3", p[3][2] == '1', 8*642+184, -1156)
			neonplcl(gpipe, ".dprogff4", p[3][3] == '1', 8*642+229, -1156)
			neonplcl(gpipe, ".dprogff5", p[3][4] == '1', 8*642+358, -1156)
			neonplcl(gpipe, ".dprogff6", p[3][5] == '1', 8*642+403, -1156)
			neonplcl(gpipe, ".dprogff7", p[3][6] == '1', 8*642+448, -1156)
			neonplcl(gpipe, ".dprogff8", p[3][7] == '1', 8*642+493, -1156)
			neonplcl(gpipe, ".ddivff", p[4][0] == '1', 8*642+417, 39)
			neonplcl(gpipe, ".dclrff", p[4][1] == '1', 8*642+400, 39)
			neonplcl(gpipe, ".dilockff", p[4][2] == '1', 8*642+383, 39)
			neonplcl(gpipe, ".ddpgamma", p[4][3] == '1', 8*642+356, 39)
			neonplcl(gpipe, ".dngamma", p[4][4] == '1', 8*642+339, 39)
			neonplcl(gpipe, ".dpsrcff", p[4][5] == '0', 8*642+286, 39)
			neonplcl(gpipe, ".dpringff", p[4][6] == '0', 8*642+269, 39)
			neonplcl(gpipe, ".ddenomff", p[4][7] == '0', 8*642+252, 39)
			neonplcl(gpipe, ".dnumrplus", p[4][8] == '1', 8*642+235, 39)
			neonplcl(gpipe, ".dnumrmin", p[4][0] == '1', 8*642+218, 39)
			neonplcl(gpipe, ".dqalpha", p[4][10] == '1', 8*642+140, -135)
			neonplcl(gpipe, ".dsac", p[4][11] == '1', 8*642+186, -135)
			neonplcl(gpipe, ".dm2", p[4][12] == '1', 8*642+234, -135)
			neonplcl(gpipe, ".dm1", p[4][13] == '1', 8*642+286, -135)
			neonplcl(gpipe, ".dnac", p[4][14] == '1', 8*642+334, -135)
			neonplcl(gpipe, ".dda", p[4][15] == '1', 8*642+384, -135)
			neonplcl(gpipe, ".dnalpha", p[4][16] == '1', 8*642+430, -135)
			neonplcl(gpipe, ".ddalpha", p[4][17] == '1', 8*642+480, -135)
			neonplcl(gpipe, ".dans2", p[4][27] == '1', 8*642+530, -135)
			neonplcl(gpipe, ".dans4", p[4][29] == '1', 8*642+578, -135)
			neonplcl(gpipe, ".ddgamma", p[4][18] == '1', 8*642+140, -156)
			neonplcl(gpipe, ".dnpgamma", p[4][19] == '1', 8*642+186, -156)
			neonplcl(gpipe, ".dp2", p[4][20] == '1', 8*642+234, -156)
			neonplcl(gpipe, ".dp1", p[4][21] == '1', 8*642+286, -156)
			neonplcl(gpipe, ".dsalpha", p[4][22] == '1', 8*642+334, -156)
			neonplcl(gpipe, ".dds", p[4][23] == '1', 8*642+384, -156)
			neonplcl(gpipe, ".dnbeta", p[4][24] == '1', 8*642+430, -156)
			neonplcl(gpipe, ".ddbeta", p[4][25] == '1', 8*642+480, -156)
			neonplcl(gpipe, ".dans1", p[4][26] == '1', 8*642+530, -156)
			neonplcl(gpipe, ".dans3", p[4][28] == '1', 8*642+580, -156)

		// Multiplier
		case "m":
			stage, _ := strconv.Atoi(p[1])
			mr1, _ := strconv.Atoi(p[3])
			mr3, _ := strconv.Atoi(p[4])
			neonplcl(gpipe, ".mstage", true, 18*642+188+stage*20, -36)
			neonplcl(gpipe, ".mr1", mr1 == 1, 17*642+312, 29)
			neonplcl(gpipe, ".mr3", mr3 == 1, 19*642+351, 29)
			for i := 0; i < 24; i++ {
				xpos := 642*(17+i/8) + 92 + 41*(i%8)
				if i%8 >= 4 {
					xpos += 255
				}
				nname = fmt.Sprintf(".mi%d", i+1)
				neonplcl(gpipe, nname, p[2][i] == '1', xpos, -1156)
			}

		// Master programmer
		case "mps":
			stage, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			if stage < 5 {
				neonplcl(gpipe, fmt.Sprintf(".ps%d", stage), true,
					82+99*stage+2*642, -223+val*19)
			} else {
				neonplcl(gpipe, fmt.Sprintf(".ps%d", stage), true,
					723+99*(stage-5)+2*642, -223+val*19)
			}
		case "mpi":
			stage, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			if stage < 5 {
				neonplcl(gpipe, fmt.Sprintf(".pi%d", stage), val > 0,
					95+stage*82+2*642, -1130)
			} else {
				neonplcl(gpipe, fmt.Sprintf(".pi%d", stage), val > 0,
					695+(stage-5)*82+2*642, -1130)
			}
		case "mpd":
			decade, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			if decade < 10 {
				neonplcl(gpipe, fmt.Sprintf(".pd%d", decade), true,
					131+49*decade+2*642, 191+val*19)
			} else {
				neonplcl(gpipe, fmt.Sprintf(".pd%d", decade), true,
					774+49*(decade-10)+2*642, 191+val*19)
			}

		// Function tables
		case "ftar":
			unit, _ := strconv.Atoi(p[1])
			arg, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, fmt.Sprintf(".ft%da1", unit+1), true, 88+ftuoff[unit]*642+20*(arg%10), 40)
			neonplcl(gpipe, fmt.Sprintf(".ft%da10", unit+1), true, 347+ftuoff[unit]*642+20*(arg/10), 40)
		case "ftr":
			unit, _ := strconv.Atoi(p[1])
			ring, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, fmt.Sprintf(".ft%dr", unit+1), true, 308+ftuoff[unit]*642+20*ring, -25)
		case "ftad":
			unit, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, fmt.Sprintf(".ft%dadd", unit+1), val == 1, 190+ftuoff[unit]*642, -25)
		case "ftsu":
			unit, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, fmt.Sprintf(".ft%dsubt", unit+1), val == 1, 210+ftuoff[unit]*642, -25)
		case "ftse":
			unit, _ := strconv.Atoi(p[1])
			val, _ := strconv.Atoi(p[2])
			neonplcl(gpipe, fmt.Sprintf(".ft%daset", unit+1), val == 1, 107+ftuoff[unit]*642, -25)

		// Constant Transmitter
		case "c":
			prog, _ := strconv.Atoi(p[1])
			ff, _ := strconv.Atoi(p[2])
			row := prog / 10
			col := prog % 10
			x := 90 + 34*642 + int(float32(col)*48.6+0.5)
			nname = fmt.Sprintf(".ct%d", prog + 1)
			switch row {
			case 0:
				neonplcl(gpipe, nname, ff == 1, x, 362)
			case 1:
				neonplcl(gpipe, nname, ff == 1, x, -119)
			case 2:
				neonplcl(gpipe, nname, ff == 1, x, -1150)
			}

		// Punch
		case "punch":
			if width > 480 {
				fmt.Fprintf(gpipe, ".outdeck insert end \"%s\\n\"\n", s[6:])
				if lines > 9 {
					fmt.Fprintln(gpipe, ".outdeck yview scroll 1 unit")
				}
				lines++
			}
		}
	}
	fmt.Fprintln(gpipe, "exit")
	gpipe.Close()
	cpipe.Close()
	cmd.Wait()
}

func rundemo(gpipe io.Writer) {
	whichscreen := 0
	for {
		time.Sleep(5 * time.Second)
		_, isclosed := fmt.Fprintln(gpipe)
		if isclosed != nil {
			return
		}
		whichscreen++
		if whichscreen > 5 {
			whichscreen = 0
		}
		switch whichscreen {
		case 1:
			resetbuts(gpipe, 1)
			fmt.Fprintln(gpipe, ".eniac configure -image s1img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		case 2:
			resetbuts(gpipe, 2)
			fmt.Fprintln(gpipe, ".eniac configure -image s2img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		case 3:
			resetbuts(gpipe, 3)
			fmt.Fprintln(gpipe, ".eniac configure -image s3img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		case 4:
			resetbuts(gpipe, 4)
			fmt.Fprintln(gpipe, ".eniac configure -image s4img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		case 5:
			resetbuts(gpipe, 5)
			fmt.Fprintln(gpipe, ".eniac configure -image s5img")
			clearstate()
			setneonsize(gpipe, "neon2")
			drawfixed(gpipe)
		case 0:
			resetbuts(gpipe, 0)
			fmt.Fprintln(gpipe, ".eniac configure -image enimg")
			clearstate()
			setneonsize(gpipe, "neon")
			drawfixed(gpipe)
		}
	}
}

func drawfixed(gpipe io.Writer) {
	x, y, vis := ray(245, -805, 0)
	if vis {
		fmt.Fprintf(gpipe, "place .gpilot -x %d -y %d\n", x, y)
	} else {
		fmt.Fprintln(gpipe, "place forget .gpilot")
	}
	x, y, vis = ray(82, -805, 0)
	if vis {
		fmt.Fprintf(gpipe, "place .apilot -x %d -y %d\n", x, y)
	} else {
		fmt.Fprintln(gpipe, "place forget .apilot")
	}
	neonplcl(gpipe, ".ih1", true, 425, -1192)
	neonplcl(gpipe, ".ih2", true, 446, -1192)
	neonplcl(gpipe, ".cycst", true, 70+642, -132)
	neonplcl(gpipe, ".cych1", true, 406+642, -1192)
	neonplcl(gpipe, ".cych2", true, 429+642, -1192)
	neonplcl(gpipe, ".cych3", true, 456+642, -1192)
	neonplcl(gpipe, ".cych4", true, 479+642, -1192)
	for i := 1; i <= 20; i++ {
		neonplcl(gpipe, fmt.Sprintf(".acc%dh1", i), true, accoff[i-1]+395, -1196)
		neonplcl(gpipe, fmt.Sprintf(".acc%dh2", i), true, accoff[i-1]+420, -1196)
	}
	neonplcl(gpipe, ".ph1", true, 320+2*642, -1198)
	neonplcl(gpipe, ".ph2", true, 341+2*642, -1198)
	neonplcl(gpipe, ".ph3", true, 927+2*642, -1198)
	neonplcl(gpipe, ".ph4", true, 950+2*642, -1198)
	neonplcl(gpipe, ".dh1", true, 342+8*642, -1198)
	neonplcl(gpipe, ".dh2", true, 366+8*642, -1198)
	for i := 0; i < 3; i++ {
		neonplcl(gpipe, fmt.Sprintf(".mh%d", 2*i), true, 396+(17+i)*642, -1196)
		neonplcl(gpipe, fmt.Sprintf(".mh%d", 2*i+1), true, 420+(17+i)*642, -1196)
	}
	neonplcl(gpipe, ".conh1", true, 427+34*642, -1197)
	neonplcl(gpipe, ".conh2", true, 448+34*642, -1197)
	neonplcl(gpipe, ".conh3", true, 1085+34*642, -1197)
	neonplcl(gpipe, ".conh4", true, 1106+34*642, -1197)
	neonplcl(gpipe, ".prh1", true, 1060+37*642, -1199)
	neonplcl(gpipe, ".prh2", true, 1082+37*642, -1199)
	for i := 0; i < 3; i++ {
		neonplcl(gpipe, fmt.Sprintf(".ft%dh1", i+1), true, 428+ftuoff[i]*642, -1182)
		neonplcl(gpipe, fmt.Sprintf(".ft%dh2", i+1), true, 450+ftuoff[i]*642, -1182)
		neonplcl(gpipe, fmt.Sprintf(".ft%dh3", i+1), true, 1028+ftuoff[i]*642, -1182)
		neonplcl(gpipe, fmt.Sprintf(".ft%dh4", i+1), true, 1052+ftuoff[i]*642, -1182)
	}
}

func neonplcl(gpipe io.Writer, name string, cond bool, x, y int) {
	if cond {
		xp, yp, vis := ray(x, y, 1)
		if vis {
			fmt.Fprintf(gpipe, "place configure %s -x %d -y %d\n", name, xp, yp)
			return
		}
	}
	fmt.Fprintf(gpipe, "place forget %s\n", name)
}

func neonpos(acc, dec, val int, accoff []int) (x, y int) {
	x = accoff[acc] + dec*49 + 63
	y = 149 + val*38
	return
}

func ray(xprime, yprime, offset int) (x, y int, vis bool) {
	switch {
	case guistate.guimode == 0:
		vis = true
		f := (80 * width) / 100
		x1 := 2568
		z0 := (x1 * f * 2) / width

		if xprime < 16*642 {
			x = width/2 - ((x1+offset*160)*f)/(xprime+z0)
			y = height/4 - (yprime*f)/(xprime+z0)
		} else if xprime < 24*642 {
			x = width/2 + ((xprime-20*642)*f)/(16*642+z0)
			y = height/4 - (yprime*f)/(16*642+z0)
		} else {
			x = width/2 + ((x1+offset*160)*f)/(40*642-xprime+z0)
			y = height/4 - (yprime*f)/(40*642-xprime+z0)
		}
	case guistate.guimode == 1 && xprime < 8*642:
		x = (xprime * width / 8) / 642
		y = height/4 - (yprime*width/8)/642
		vis = true
	case guistate.guimode == 2 && xprime >= 8*642 && xprime < 16*642:
		x = ((xprime - 8*642) * width / 8) / 642
		y = height/4 - (yprime*width/8)/642
		vis = true
	case guistate.guimode == 3 && xprime >= 16*642 && xprime < 24*642:
		x = ((xprime - 16*642) * width / 8) / 642
		y = height/4 - (yprime*width/8)/642
		vis = true
	case guistate.guimode == 4 && xprime >= 24*642 && xprime < 32*642:
		x = ((xprime - 24*642) * width / 8) / 642
		y = height/4 - (yprime*width/8)/642
		vis = true
	case guistate.guimode == 5 && xprime >= 32*642:
		x = ((xprime - 32*642) * width / 8) / 642
		y = height/4 - (yprime*width/8)/642
		vis = true
	default:
		x = 0
		y = 0
		vis = false
	}
	return
}
