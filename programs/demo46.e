#
# An attempt to recreate the demo at the unveiling of the ENIAC
# on Feb 15, 1946
#

#
# The initiate button is used to start each part of the demo
# by advancing the A input on the Master Programmer.  The
# end of each part advances the A stepper.
#

p i.Io 1-1
p 1-1 p.Ai
p 1-11 p.Adi
s p.a20 B
s p.cA 6

#
# The first demo is adding the number 97367 to itself 5000
# times for a total of 486,835,000.
#
# Based on Burks' writeup describing the Feb 1, press demo,
# it appears that the first load of 97367 into an ACC is done
# in 1Add mode, and the 5000 adds are run at full speed.
# Then the whole process is run again with a 100Hz clock.
# We'll use Stepper C for the 5000 counter.
#

p c.o 1
p 1 ad.s.1.-5
p ad.s.1.-5 a1.α
p p.A1o 1-2
p 1-2 a1.1i
p 1-2 c.25i
p c.25o 1-3
s c.s25 Jl
s a1.op1 α
s c.j10 9
s c.j9 7
s c.j8 3
s c.j7 6
s c.j6 7
p 1-3 a1.12i
p a1.12o 1-4
s a1.op12 0
s a1.r12 1
p 1-4 p.Ci
p p.C1o 1-4
p p.C2o 1-11
p 1-4 a1.2i
p 1-4 a2.2i
p a1.A 1
p 1 a2.α
s a1.op2 A
s a2.op2 α
s p.cC 2
s p.a18 B
s p.a14 C
s p.d17s1 4
s p.d16s1 9
s p.d15s1 9
s p.d14s1 8

#
# The second demo is of the multiplier.  In 1Add mode, the
# number 13975 is loaded into Acc 9 and 10.
#

p p.A2o 2-1
p 1 a9.α
p 1 a10.α
p 2-1 a9.7i
p 2-1 a10.7i
p 2-1 c.26i
p c.26o 2-2
s a9.op7 α
s a10.op7 α
s c.s26 Jr
s c.j5 1
s c.j4 3
s c.j3 9
s c.j2 7
s c.j1 5
p 2-2 p.Bi
p p.B1o 2-3
s p.cB 2
s p.d20m1 4
s p.d19m1 9
s p.d18m1 8
p p.B2o 2-1
p 2-3 m.1i
s m.ieracc1 0
s m.icandacc1 0
s m.sf1 0
s m.place1 10
s m.prod1 0
s m.iercl1 0
s m.icandcl1 0

#
# Full multiplier configuration for 20-digit
# products.
#
p m.lhppI 6
p m.lhppII 7
p m.rhppI 8
p m.rhppII 9
p 6 a11.α
p 7 a12.α
p 8 a13.α
p 9 a14.α
p a9.S 7
p 7 a11.β
p a12.A 7
p 7 a14.β
p a10.S 6
p a11.A 6
p 6 a13.β

p m.Rα 9-1
p m.Rβ 9-2
p m.Rγ 9-3
p m.Rδ 9-4
p m.Rε 9-5
p m.Dα 9-6
p m.Dβ 9-7
p m.Dγ 9-8
p m.Dδ 9-9
p m.Dε 9-10
p 9-1 a9.1i
p 9-2 a9.2i
p 9-3 a9.3i
p 9-4 a9.4i
p 9-5 a9.5i
p 9-6 a10.1i
p 9-7 a10.2i
p 9-8 a10.3i
p 9-9 a10.4i
p 9-10 a10.5i

p m.A 10-1
p m.S 10-2
p m.AS 10-3
p m.AC 10-4
p m.SC 10-5
p m.ASC 10-6
p m.RS 10-7
p m.DS 10-8
p m.F 10-9
p 10-1 a13.3i
p 10-2 a13.4i
p 10-3 a13.5i
p 10-4 a13.6i
p 10-5 a13.7i
p 10-6 a13.8i
p 10-7 a9.6i
p 10-7 a11.2i
p 10-8 a10.6i
p 10-8 a13.2i
p 10-9 a11.1i
p 10-9 a13.1i
p a11.st1 m.l
p a13.st1 m.r
p a12.st1 a11.st2
p a12.su1 a11.su2
p a14.st1 a13.st2
p a14.su1 a13.su2

s a9.op1 α
s a9.op2 β
s a9.op3 γ
s a9.op4 δ
s a9.op5 ε
s a9.rp5 1
s a9.op6 S
s a9.rp6 1
s a10.op1 α
s a10.op2 β
s a10.op3 γ
s a10.op4 δ
s a10.op5 ε
s a10.rp5 1
s a10.op6 S
s a10.rp6 1
s a11.op1 A
s a11.cc1 C
s a11.op2 β
s a13.op1 β
s a13.cc1 0
s a13.op2 β
s a13.cc2 0
s a13.op3 A
s a13.cc3 0
s a13.op4 S
s a13.cc4 0
s a13.op5 AS
s a13.cc5 0
s a13.rp5 1
s a13.op6 A
s a13.cc6 C
s a13.rp6 1
s a13.op7 S
s a13.cc7 C
s a13.rp7 1
s a13.op8 AS
s a13.cc8 C
s a13.rp8 1

