# Multiplier test
# Minimal multiplier wiring
p m.Rα 9-1
p 9-1 a9.1i
p m.Dα 9-2
p 9-2 a10.1i
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
p a9.S 7
p 7 a11.β
p a10.S 6
p a11.A 6
p 6 a13.β
p m.lhppI 8
p 8 a11.α
p m.rhppI 9
p 9 a13.α
p a11.st1 m.l
p a13.st1 m.r
s m.ieracc1 α
s m.iercl1 C
s m.icandacc1 α
s m.icandcl1 C
s m.sf1 O
# s m.place1.6
s m.place1 2
s m.prod1 A
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

# load the argument accumulators
# preload J constant into ACC 1
p i.io 1-1
p 1-1 a1.1i
p 1-1 c.25i
p c.o 1
p 1 a1.α
s a1.op1 α
s c.s25 Jlr
# transfer from ACC 1 to ACC 9, and load
# constant K into ACC 10
p c.25o 1-2
p 1-2 m.1i
p 9-1 a1.2i
p a1.A 2
p 2 a9.α
s a1.op2 A
s a1.cc2 C
p 1 a10.α
p 9-1 c.26i
s c.s26 Klr
# simple example to show multiplier
s c.j10 4
s c.j9 2
s c.k10 3
s c.k9 4
s c.k8 7
# go into single add cycle mode
s cy.op 1a
