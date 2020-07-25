# Divider test
# Divider wiring
p d.ans 8
p 8 a2.α
p 8 a5.γ
p a5.A 9
p a5.S 9
p a3.A 9
p a7.A 9
p 9 a3.γ
p 9 ad.s.1.1
p ad.s.1.1 a7.α
p a2.st1 d.su2q
p a3.st1 d.sv
p a5.st1 d.su3
p a7.st1 d.su2s
s d.nr1 α
s d.nc1 C
s d.dr1 α
s d.dc1 C
s d.pl1 D4
# s d.pl1 D10
s d.ro1 RO
s d.an1 1
s d.il1 NI

# answer disposal: transfer to ACC1
p d.1o 1-3
p 1-3 a1.3i
p a2.A 7
p 7 a1.γ
s a1.op3 γ

# load the argument accumulators
# preload J constant into ACC 1
p i.io 1-1
p 1-1 a1.1i
p 1-1 c.25i
p c.o 1
p 1 a1.α
s a1.op1 α
s c.s25 Jlr
# transfer from ACC 1 to ACC 3, and load
# constant K into ACC 5
p c.25o 1-2
p 1-2 d.1i
p 1-2 a1.2i
p a1.A 2
p 2 a3.α
s a1.op2 A
s a1.cc2 C
p 1 a5.α
p 1-2 c.26i
s c.s26 Klr
# set the J and K constants
# s c.j10 0
# s c.j9 3
# s c.j8 0
# s c.k10 0
# s c.k9 1
# s c.k8 6
# s c.k7 0
# Example from Goldstine's tech manual
# Answer: P 0 091 000 000 in Acc 1
s c.j10 0
s c.j9 2
s c.j8 0
s c.j7 9
s c.j6 0
s c.j5 7
s c.j4 0
s c.j3 0
s c.j2 0
s c.j1 0
s c.k10 0
s c.k9 2
s c.k8 3
s c.k7 0
s c.k6 0
s c.k5 0
s c.k4 0
s c.k3 0
s c.k2 0
s c.k1 0
# go into single add cycle mode
s cy.op 1a
