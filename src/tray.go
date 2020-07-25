package main

import "fmt"

type trunk struct {
	xmit    [16]chan pulse
	recv    []chan pulse
	started bool
	update  chan int
}

var dtrays [20]trunk
var ctrays [121]trunk

func treset(t *trunk) {
	for i, _ := range t.xmit {
		t.xmit[i] = nil
	}
	t.recv = nil
	t.started = false
	if t.update != nil {
		t.update <- -1
		t.update = nil
	}
}

func trayreset() {
	for i, _ := range dtrays {
		treset(&dtrays[i])
	}
	for i, _ := range ctrays {
		treset(&ctrays[i])
	}
}

func handshake(val int, ch chan pulse, resp chan int) {
	if ch != nil {
		ch <- pulse{val, resp}
		<-resp
	}
}

func dotrunk(t *trunk) {
	var x, p pulse

	p.resp = make(chan int)
	for {
		select {
		case q := <-t.update:
			if q == -1 {
				return
			}
			continue
		case x = <-t.xmit[0]:
		case x = <-t.xmit[1]:
		case x = <-t.xmit[2]:
		case x = <-t.xmit[3]:
		case x = <-t.xmit[4]:
		case x = <-t.xmit[5]:
		case x = <-t.xmit[6]:
		case x = <-t.xmit[7]:
		case x = <-t.xmit[8]:
		case x = <-t.xmit[9]:
		case x = <-t.xmit[10]:
		case x = <-t.xmit[11]:
		case x = <-t.xmit[12]:
		case x = <-t.xmit[13]:
		case x = <-t.xmit[14]:
		case x = <-t.xmit[15]:
		}
		p.val = x.val
		if x.val != 0 {
			needresp := 0
			for _, c := range t.recv {
				if c != nil {
				pulseloop:
					for {
						select {
						case c <- p:
							needresp++
							break pulseloop
						case <-p.resp:
							needresp--
						}
					}
				}
			}
			for needresp > 0 {
				<-p.resp
				needresp--
			}
		}
		if x.resp != nil {
			x.resp <- 1
		}
	}
}

func trunkxmit(ilk, n int, ch chan pulse) {
	var t *trunk

	if ilk == 0 {
		t = &dtrays[n]
	} else {
		t = &ctrays[n]
	}
	if !t.started {
		t.update = make(chan int)
		go dotrunk(t)
		t.started = true
	}
	for i, c := range t.xmit {
		if c == nil {
			t.xmit[i] = ch
			t.update <- 1
			return
		}
	}
	fmt.Printf("Too many transmitters on %d:%d\n", ilk, n)
}

func trunkrecv(ilk, n int, ch chan pulse) {
	var t *trunk

	if ilk == 0 {
		t = &dtrays[n]
	} else {
		t = &ctrays[n]
	}
	if t.recv == nil {
		t.recv = make([]chan pulse, 0, 20)
	}
	for i, c := range t.recv {
		if c == nil {
			t.recv[i] = ch
			return
		}
	}
	if len(t.recv) != cap(t.recv) {
		t.recv = append(t.recv, ch)
	}
}
