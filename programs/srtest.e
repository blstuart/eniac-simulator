# Square rooter test
# Square rooter basic wiring
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
s d.dr1 0
s d.dc1 0
s d.pl1 R4
# s d.pl1 R8
s d.ro1 RO
s d.an1 4
s d.il1 NI
s d.da A
s d.ra A

# answer disposal: transfer to ACC1
p d.1o 1-3
p 1-3 a1.1i
p 9 a1.α


# load constant J into ACC 3 (radicand)
p c.o 1
p 1 a3.α
p i.io 1-1
p 1-1 c.25i
p 1-1 d.1i
s c.s25 Jlr
# set the J constant
# s c.j10.0
# s c.j9.5
# s c.j8.0
# s c.j8 2
# Example from Goldstine's tech manual
# Answer: P 0 180 000 000 in Acc 1
s c.j10 0
s c.j9 0
s c.j8 8
s c.j7 1
s c.j6 3
s c.j5 6
s c.j4 0
s c.j3 4
s c.j2 0
s c.j1 0
# go into single add cycle mode
s cy.op 1a
