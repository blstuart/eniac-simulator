# Test for the 20-digit accumulator support

# First some basic addition and subtraction
# Result: M 9 999 999 999 in Acc 15 and
# P 1 358 024 688 in Acc 16
# - set up pairing of 15-16 and 17-18 using PX-5-110
# - PX-5-109 load boxes and PX-5-121 are default
p a16.st1 a15.st2
p a16.su1 a15.su2
p a18.st1 a17.st2
p a18.su1 a17.su2
# - load low order accumulators with large values
p i.io 1-1
p c.o 1
p 1 a16.α
p 1 a18.α
p 1-1 c.25i
p 1-1 a15.1i
p c.25o 1-2
p 1-2 c.26i
p 1-2 a18.1i
p c.26o 1-3
s c.s25 Jlr
s c.s26 Klr
s a15.op1 α
s a18.op1 α
s c.j10 1
s c.j9 2
s c.j8 3
s c.j7 4
s c.j6 5
s c.j5 6
s c.j4 7
s c.j3 8
s c.j2 9
s c.j1 9
s c.k10 9
s c.k9 8
s c.k8 7
s c.k7 6
s c.k6 5
s c.k5 4
s c.k4 3
s c.k3 2
s c.k2 1
s c.k1 1
# - add 17-18 to 15-16
p a17.A 2
p a18.A 3
p 2 a15.β
p 3 a16.β
p 1-3 a15.5i
p 1-3 a18.5i
p a15.5o 1-4
s a15.op5 β
s a18.op5 A
# - subtract 17-18 from 15-16 twice
p a17.S 2
p a18.S 3
p 1-4 a16.5i
p 1-4 a17.5i
p a16.5o 1-5
s a16.op5 β
s a16.rp5 2
s a17.op5 S
s a17.rp5 2

# Next step is to multiply them.
# Results: P 1 219 326 320 in Acc 13 and
# P 1 386 983 689 in Acc 14
# Multiplier wiring
# p m.Rα 9-1
# p 9-1 a9.1i
# p m.Dα 9-2
# p 9-2 a10.1i
p m.A 9-3
p 9-3 a13.3i
p m.RS 9-4
p 9-4 a9.6i
p 9-4 a11.2i
p m.DS 9-5
p 9-5 a10.6i
p 9-5 a13.2i
p m.F 9-6
p 9-6 a11.1i
p 9-6 a13.1i
p a9.S 4
p 4 a11.β
p a10.S 5
p a11.A 5
p 5 a13.β
p m.lhppI 7
p 7 a11.α
p m.lhppII 6
p 6 a12.α
p m.rhppI 9
p 9 a13.α
p m.rhppII 8
p 8 a14.α
p a12.A 8
p 8 a14.β
p a11.st1 m.l
p a12.st1 a11.st2
p a12.su1 a11.su2
p a13.st1 m.r
p a14.st1 a13.st2
p a14.su1 a13.su2
s m.ieracc1 0
s m.iercl1 C
s m.icandacc1 0
s m.icandcl1 C
s m.sf1 0
s m.place1 10
s m.prod1 0
s a9.op1 α
s a9.op6 S
s a9.rp6 1
s a10.op1 α
s a10.op6 S
s a10.rp6 1
s a11.op1 A
s a11.cc1 C
s a11.op2 β
s a13.op1 β
s a13.op2 β
s a13.op3 A

p 1 a9.α
p 1 a10.α
p 1-1 a9.1i
p 1-2 a10.1i
p 1-5 m.1i
p m.1o 1-6

