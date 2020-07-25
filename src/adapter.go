package main

var dpin [40]chan pulse
var dpout [40][11]chan pulse
var shiftin [40]chan pulse
var shiftout [40]chan pulse
var delin [40]chan pulse
var delout [40]chan pulse
var sdin [40]chan pulse
var sdout [40]chan pulse

func adreset() {
	for i := 0; i < 40; i++ {
		dpin[i] = nil
		for j := 0; j < 11; j++ {
			dpout[i][j] = nil
		}
		shiftin[i] = nil
		shiftout[i] = nil
		delin[i] = nil
		delout[i] = nil
		sdin[i] = nil
		sdout[i] = nil
	}
}

func adplug(ilk string, inout, which, param int, ch chan pulse) {
	switch ilk {
	case "dp":
		if inout == 0 {
			dpin[which] = ch
			go digitprog(dpin[which], which)
		} else {
			dpout[which][param-1] = ch
		}
	case "s":
		if inout == 0 {
			shiftin[which] = ch
		} else {
			shiftout[which] = ch
		}
		if shiftin[which] != nil && shiftout[which] != nil {
			go shifter(shiftin[which], shiftout[which], param)
		}
	case "d":
		if inout == 0 {
			delin[which] = ch
		} else {
			delout[which] = ch
		}
		if delin[which] != nil && delout[which] != nil {
			go deleter(delin[which], delout[which], param)
		}
	case "sd":
		if inout == 0 {
			sdin[which] = ch
		} else {
			sdout[which] = ch
		}
		if sdin[which] != nil && sdout[which] != nil {
			go specdig(sdin[which], sdout[which], uint(param))
		}
	}
}

func digitprog(in chan pulse, which int) {
	resp := make(chan int)
	for {
		d := <-in
		for i := uint(0); i < 11; i++ {
			if d.val&(1<<i) != 0 && dpout[which][i] != nil {
				dpout[which][i] <- pulse{1, resp}
				<-resp
			}
		}
		if d.resp != nil {
			d.resp <- 1
		}
	}
}

func shifter(in, out chan pulse, shift int) {
	for {
		d := <-in
		if shift >= 0 {
			d.val = (d.val & (1 << 10)) | ((d.val << uint(shift)) & ((1 << 10) - 1))
		} else {
			x := d.val >> uint(-shift)
			if d.val&(1<<10) != 0 {
				d.val = x | (((1 << 11) - 1) & ^((1 << uint(11+shift)) - 1))
			} else {
				d.val = x
			}
		}
		if d.val != 0 {
			out <- d
		} else if d.resp != nil {
			d.resp <- 1
		}
	}
}

func deleter(in, out chan pulse, which int) {
	for {
		d := <-in
		if which >= 0 {
			d.val &= ^((1 << uint(10-which)) - 1)
		} else {
			d.val &= (1 << uint(10+which)) - 1
		}
		if d.val != 0 {
			out <- d
		} else if d.resp != nil {
			d.resp <- 1
		}
	}
}

func specdig(in, out chan pulse, which uint) {
	for {
		d := <-in
		x := d.val >> which
		mask := 0x07fc
		if d.val&(1<<10) != 0 {
			d.val = x | mask
		} else {
			d.val = x & ^mask
		}
		if d.val != 0 && out != nil {
			out <- d
		} else if d.resp != nil {
			d.resp <- 1
		}
	}
}
