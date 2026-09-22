# Generate sine and cosine using double integration

p c.o 1
s c.j1 0
s c.j2 0
s c.j3 0
s c.j4 0
s c.j5 0
s c.j6 0
s c.j7 0
s c.j8 0
s c.j9 0
s c.j10 1
s c.s30 Jlr
p i.io 1-1
p 1-1 c.30i

p 1 a1.alpha
p 1-1 a1.5i
s a1.op5 alpha
s a1.rp5 1
s a1.cc5 0
p a1.5o 1-2

p a1.A 1
p 1-2 a1.6i
s a1.op6 A
s a1.rp6 1
s a1.cc6 0
p a1.6o 1-3

p 1 ad.s.1.-3
p ad.s.1.-3 a2.alpha
p 1-2 a2.5i
s a2.op5 alpha
s a2.rp5 1
s a2.cc5 0

p a2.S 1
p 1-3 a2.6i
s a2.op6 S
s a2.rp6 1
s a2.cc6 0

p 1 ad.s.2.-3
p ad.s.2.-3 a1.beta
p 1-3 a1.7i
s a1.op7 beta
s a1.rp7 1
s a1.cc7 C
p a1.7o 1-2
