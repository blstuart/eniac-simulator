# Initiating unit
p 1-1 i.Ci1
p i.Co1 1-7
p 1-1 i.Ri
p 1-7 i.Rl
p i.Ro 1-8
p 1-5 i.Pi
p i.Po 1-2
p i.Io 1-6

# Master Programmer
p 1-2 p.Ai
p p.A1o 1-3
p p.A2o 1-1
p p.A3o 1-3
p p.A4o 1-6
p 1-6 p.Ci
p p.C1o 1-1
p 1-3 p.Di
p p.D1o 1-4
p p.D2o 1-5
s p.a20 A
s p.a18 C
s p.a14 C
s p.a12 D
s p.cA 4
s p.cC 2
s p.cD 2
s p.d20s1 6
s p.d20s2 1
s p.d20s3 1
s p.d20s4 1
# s p.d16s1 2
s p.d15s1 0
# s p.d14s1 0
s p.d14s1 7
s p.d14s2 1
s p.d13s1 1
s p.d12s1 0
s p.d12s2 1

# Accumulator 13
p 1 a13.α
p a13.A 1
p 1-8 a13.1i
p 1-4 a13.5i
p a13.5o 1-10
p 1-10 a13.6i
p a13.6o 1-3
s a13.sc SC
s a13.op1 α
s a13.op5 A
s a13.rp5 1
s a13.op6 α
s a13.rp6 1

# Accumulator 14
p 1 a14.α
p 1-4 a14.5i
s a14.sc SC
s a14.op5 α
s a14.rp5 1

# Accumulator 16
p 1 a16.α
p a16.A 2
p 1-9 a16.1i
p 1-4 a16.2i
s a16.sc SC
s a16.op1 α
s a16.op2 A

# Accumulator 17
p 2 a17.α
p 1-4 a17.2i
s a17.sc SC
s a17.op2 α

# Constant Transmitter
p c.o 1
p 1-8 c.1i
p c.1o 1-9
p 1-9 c.2i
p c.2o 1-2
p 1-10 c.25i
s c.s1 Blr
s c.s2 Alr
s c.s25 Jlr
s c.jl M
s c.jr M
s c.j10 9
s c.j9 9
s c.j8 9
s c.j7 9
s c.j6 9
s c.j5 9
s c.j4 9
s c.j3 9
s c.j2 6
s c.j1 8

# Printer
s pr.2 P
s pr.3 P
s pr.4 P
s pr.5 P
s pr.7 P
s pr.8 P
s pr.9 P
s pr.10 P

f r ch10b.card
