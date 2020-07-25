# From Adele Goldstine, ENIAC Technical Manual
# Chapter 6, Section 6.5, Figure 6-2
# Computes {\sqrt{a}+\sum_{i=1}^3 x_i^3 \over b}+cd
# Answer: P 0 531 603 500 in Acc 12

# Master Programmer
p 1-2 p.Ai
p 1-2 p.Adi
# p p.A1o 2-4
# p p.A2o 2-5
# p p.A3o 2-6
p p.A2o 2-4
p p.A3o 2-5
p p.A4o 2-6
# s p.cA 3
s p.cA 4
s p.a20 B

# ACCs 1, 4, 6, 8, 12, 14, 15
p a1.A 3
p 1-1 a1.5i
p a1.5o 2-1
s a1.op5 A
s a1.rp5 1
s a1.cc5 C

p a4.S 4
p 1-4 a4.1i
s a4.op1 S
s a4.cc1 C

# p a.6.S 1
p a6.A 1
p 1-4 a6.1i
# s a6.op1 S
s a6.op1 A
s a6.cc1 C

p a8.A 2
p 1-4 a8.1i
s a8.op1 A
s a8.cc1 C

p 1 ad.d.1.8
p ad.d.1.8 a12.α
p 5 a12.β
p a12.A 1
p a12.S 3
p 1-4 a12.1i
p 1-5 a12.2i
p 2-1 a12.5i
p 1-2 a12.6i
p 1-3 a12.7i
p a12.7o 1-4
p 1-6 a12.8i
p a12.8o 1-7
s a12.op1 S
s a12.cc1 C
s a12.op2 α
s a12.op5 A
s a12.cc5 C
s a12.rp5 2
s a12.op6 α
s a12.rp6 2
s a12.op7 β
s a12.op8 α

p a14.A 1
# p 2-4 a14.5i
p 2-4 a14.12i
p a14.12o 5-1
p 5-1 a14.5i
s a14.op5 A
s a14.cc5 C
# s a14.rp5 2

p a15.A 1
# p 2-5 a15.5i
p 2-5 a15.12i
p a15.12o 5-2
p 5-2 a15.5i
s a15.op5 A
s a15.cc5 C
# s a15.rp5 2

# Divider/Square Rooter w/ assoc ACCs: 2, 3, 5, 7
p d.ans 11
p 1-1 d.1i
p 2-6 d.1l
p d.1o 1-3
p 1-4 d.2i
p d.2o 1-6
p a2.st1 d.su2q
p a3.st1 d.sv
p a5.st1 d.su3
p a7.st1 d.su2s
s d.nr1 α
s d.nc1 C
s d.dr1 0
s d.dc1 0
s d.pl1 R8
s d.ro1 RO
s d.an1 4
s d.il1 I
s d.nr2 α
s d.nc2 C
s d.dr2 α
s d.dc2 C
s d.pl2 D8
s d.ro2 RO
s d.an2 2
s d.il2 NI
s d.da A
s d.ra A

p 11 a2.α
p a2.A 1

p 3 a3.α
p 5 a3.γ
p a3.A 5

p 4 a5.α
p 11 a5.γ
p a5.A 5
p a5.S 5

p 5 ad.s.1.1
p ad.s.1.1 a7.α
p a7.A 5

# Multiplier w/ ACCs 9, 10, 11, 13
p m.lhppI 9
p m.rhppI 8
p m.Rα a9.1i
p m.Dα a10.1i
p m.Dβ a10.2i
p 2-2 m.9i
p m.9o 2-3
p 2-3 m.10i
p m.10o 1-2
p 1-4 m.11i
p m.11o 1-5
p m.AC a13.1i
p m.SC a13.5i
p m.RS 3-1
p m.DS 3-2
p m.F 3-3
p 3-1 a9.3i
p 3-1 a11.1i
p 3-2 a10.3i
p 3-2 a13.2i
p 3-3 a11.2i
p 3-3 a13.3i
p a11.st1 m.l
p a13.st1 m.r
s m.ieracc9 α
s m.iercl9 0
s m.icandacc9 α
s m.icandcl9 C
s m.sf9 10
s m.place9 6
s m.prod9 AC
s m.ieracc10 0
s m.iercl10 C
s m.icandacc10 α
s m.icandcl10 C
s m.sf10 8
s m.place10 6
s m.prod10 SC
s m.ieracc11 α
s m.iercl11 C
s m.icandacc11 β
s m.icandcl11 C
s m.sf11 8
s m.place11 6
s m.prod11 AC

p 1 a9.α
p a9.S 7
p 2-1 a9.5i
p a9.5o 2-2
p 2-4 a9.6i
p a9.6o 2-2
p 2-5 a9.7i
p a9.7o 2-2
s a9.op1 α
s a9.op3 S
s a9.op5 0
s a9.rp5 1
s a9.op6 0
s a9.rp6 1
s a9.op7 0
s a9.cc7 C
s a9.rp7 1

p 1 a10.α
p 2 a10.β
p a10.S 6
s a10.op1 α
s a10.op2 β
s a10.op3 S

p 9 a11.α
p 7 a11.β
p a11.A 6
s a11.op1 β
s a11.op2 A
s a11.cc2 C

p 8 a13.α
p 6 a13.β
p a13.A 1
s a13.op1 A
s a13.cc1 C
s a13.op2 β
s a13.op3 β
s a13.op5 A
s a13.cc5 C
s a13.rp5 2

# Set up initial values for testing
p i.io 4-1
p c.o 10
p 4-1 c.25i
p c.25o 4-2
p 4-2 c.26i
p c.26o 4-3
p 4-3 c.27i
p c.27o 1-1
s c.s25 jlr
s c.s26 klr
s c.s27 klr
s c.j8 2
s c.k8 4

p 10 a1.α
p 4-1 a1.9i
s a1.op9 α

p 10 a4.α
p 4-2 a4.9i
p 4-3 a4.10i
s a4.op9 α
s a4.op10 α

p 10 ad.s.6.2
p ad.s.6.2 a6.α
p 4-1 a6.9i
s a6.op9 α

p 10 ad.s.2.2
p ad.s.2.2 a8.α
p 4-1 a8.9i
s a8.op9 α

p 10 ad.s.3.1
p ad.s.3.1 a12.δ
p 4-2 a12.9i
s a12.op9 δ

p 10 ad.s.4.1
p ad.s.4.1 a14.α
p 4-1 a14.9i
p 4-2 a14.10i
s a14.op9 α
s a14.op10 α

p 10 ad.s.5.1
p ad.s.5.1 a15.α
p 4-2 a15.9i
p 4-3 a15.10i
s a15.op9 α
s a15.op10 α
