#
# An attempt to recreate the demo at the unveiling of the ENIAC
# on Feb 15, 1946
#

#
# The initiate button is used to start each part of the demo
# by pulsing the A input on the Master Programmer.  The
# end of each part advances the A stepper.
#

#
# Before that starts, the very first thing is to load a card
# of constants using the Reader Start switch.  The completion
# of the read, also transfers the first constant to Acc 1.
#
# The card contains the data:
# 0000097367000001397500000981750099000000-00000321700097500000003007903
#
f r demo.card
p i.Ro 1-2
p c.o 1
p 1 a1.α
p 1-2 a1.1i
p 1-2 c.1i
s c.s1 Alr
s a1.op1 α
# s c.j10 9
# s c.j9 7
# s c.j8 3
# s c.j7 6
# s c.j6 7

#
# This first part is setting up the Master Programmer.
#
p i.Io 1-1
s p.cA 6
s p.cB 6
s p.cC 6
s p.cH 3
s p.a20 B
s p.a18 B
s p.a14 C
s p.a8 G
s p.a4 H
p 1-1 p.Ai
p 1-11 p.Adi
p p.A1o 1-6
p p.A2o 1-3
p p.A3o 1-4
p 1-3 p.Bi
p p.B1o 2-1
p p.B2o 2-6
p p.B3o 2-2
p p.B4o 2-11
p p.B5o 3-1
p p.B6o 3-10
s p.d20s1 1
s p.d19s1 0
s p.d18s1 0
s p.d20s2 0
s p.d19s2 0
s p.d18s2 1
s p.d20s3 1
s p.d19s3 0
s p.d18s3 0
s p.d20s4 0
s p.d19s4 0
s p.d18s4 1
s p.d20s5 1
s p.d19s5 0
s p.d18s5 0
s p.d20s6 0
s p.d19s6 0
s p.d18s6 1
p 1-6 p.Ci
p p.C1o 1-5
p p.C3o 1-7
p p.C4o 1-8
p p.C5o 1-9
p p.C6o 1-11
s p.d17s1 1
s p.d16s1 0
s p.d15s1 0
s p.d14s1 0
s p.d17s2 0
s p.d16s2 0
s p.d15s2 0
s p.d14s2 1
s p.d17s3 0
s p.d16s3 0
s p.d15s3 0
s p.d14s3 1
s p.d17s4 0
s p.d16s4 0
s p.d15s4 0
s p.d14s4 1
s p.d17s5 0
s p.d16s5 4
s p.d15s5 9
s p.d14s5 9
s p.d17s6 0
s p.d16s6 0
s p.d15s6 0
s p.d14s6 1
p 1-4 p.Hi
p p.H1o 5-1
p p.H2o 2-5
p p.H3o 5-10
s p.d7s1 0
s p.d6s1 1
s p.d5s1 0
s p.d4s1 0
s p.d7s2 0
s p.d6s2 0
s p.d5s2 0
s p.d4s2 1
s p.d7s3 0
s p.d6s3 0
s p.d5s3 0
s p.d4s3 1

#
# The first demo is adding the number 97367 to itself 5000
# times for a total of 486,835,000.
#
# Then the whole process is run again with a 100Hz clock.
# We'll use Stepper C for the 5000 counter by setting it to
# 1000 and having the adds repeat 5 times each.
#
p 1-5 a1.5i
p 1-5 a2.5i
p a1.5o 1-6
p a1.A 1
p 1 a2.α
s a1.op5 A
s a1.cc5 0
s a1.rp5 5
s a2.op5 α
s a2.cc5 0
s a2.rp5 5
p 1-7 a1.2i
p 1-7 a2.2i
s a1.op2 0
s a1.cc2 C
s a2.op2 0
s a2.cc2 C

#
# The second demo is of the multiplier.  In the first step,
# the number 13975 is loaded into Acc 9 and 10 and
# multiplied.  Then on another initiate pulse, 499 more
# multiplications are performed.
#
p 1 a9.α
p 1 a10.α
p 1-7 a9.7i
p 1-7 a10.7i
p 1-7 c.2i
s a9.op7 α
s a9.rp7 1
s a10.op7 α
s a10.rp7 1
s c.s2 Blr
# s c.j5 1
# s c.j4 3
# s c.j3 9
# s c.j2 7
# s c.j1 5
p 1-8 m.1i
s m.ieracc1 0
s m.icandacc1 0
s m.sf1 0
s m.place1 10
s m.prod1 0
s m.iercl1 0
s m.icandcl1 0
p 1-9 m.2i
p m.2o 1-6
s m.ieracc2 0
s m.icandacc2 0
s m.sf2 0
s m.place2 10
s m.prod2 0
s m.iercl2 0
s m.icandcl2 0

p 2-6 a9.8i
p 2-6 a10.8i
p 2-6 a13.9i
s a9.op8 0
s a9.cc8 C
s a9.rp8 1
s a10.op8 0
s a10.cc8 C
s a10.rp8 1
s a13.op9 0
s a13.cc9 C
s a13.rp9 1

#
# The next step is compute a table of squares and
# cubes, first with printing, then without.
#
p a15.A 1
p a16.A 1
p 1 a16.α
p 1 a18.α
p 2-1 a16.5i
p 2-1 a18.5i
s a16.op5 A
s a16.cc5 0
s a16.rp5 3
s a18.op5 α
s a18.cc5 0
s a18.rp5 3
p a16.5o 2-3
p 2-3 a15.6i
p 2-3 a16.6i
p 2-3 a18.6i
s a15.op6 A
s a15.cc6 0
s a15.rp6 3
s a16.op6 α
s a16.cc6 0
s a16.rp6 2
s a18.op6 α
s a18.cc6 0
s a18.rp6 3
p a15.6o 2-4
p 2-4 a15.7i
p 2-4 a16.7i
p 2-4 a18.7i
s a15.op7 ε
s a15.cc7 C
s a15.rp7 1
s a16.op7 ε
s a16.cc7 C
s a16.rp7 1
s a18.op7 ε
s a18.cc7 C
s a18.rp7 1
p a16.7o 2-5
p 2-5 i.Pi
p i.Po 1-1

p 2-6 a15.1i
p 2-6 a16.1i
p 2-6 a18.1i
s a15.op1 0
s a15.cc1 C
s a16.op1 0
s a16.cc1 C
s a18.op1 0
s a18.cc1 C

p 2-2 a16.8i
p 2-2 a18.8i
s a16.op8 A
s a16.cc8 0
s a16.rp8 3
s a18.op8 α
s a18.cc8 0
s a18.rp8 3
p a16.8o 2-7
p 2-7 a15.9i
p 2-7 a16.9i
p 2-7 a18.9i
s a15.op9 A
s a15.cc9 0
s a15.rp9 3
s a16.op9 α
s a16.cc9 0
s a16.rp9 2
s a18.op9 α
s a18.cc9 0
s a18.rp9 3
p a15.9o 2-8
p 2-8 a15.10i
p 2-8 a16.10i
p 2-8 a18.10i
s a15.op10 ε
s a15.cc10 C
s a15.rp10 1
s a16.op10 ε
s a16.cc10 C
s a16.rp10 1
s a18.op10 ε
s a18.cc10 C
s a18.rp10 1
p a16.10o 1-3

s pr.11-12 C
s pr.15-16 C
s pr.6 P
s pr.8 P
s pr.11 P
s pr.12 P
s pr.15 P
s pr.16 P

#
# The penultimate demonstration is a computation of
# the values of the sine and cosine of 100 angles.
# B=0.000098175
#
p 2-11 a15.2i
p 2-11 a16.2i
p 2-11 a18.2i
p 2-11 a20.2i
p a18.A 2
p 2 ad.s.2.3
p ad.s.2.3 a20.α
s a15.op2 0
s a15.cc2 C
s a16.op2 0
s a16.cc2 C
s a18.op2 A
s a18.cc2 C
s a20.op2 α
s a20.cc2 0
p 2-11 a10.9i
p 1 ad.s.3.1
p ad.s.3.1 a10.β
s a10.op9 β
s a10.cc9 0
s a10.rp9 1
p 2-11 c.7i
s c.s7 Clr
# s c.k5 9
# s c.k4 8
# s c.k3 1
# s c.k2 7
# s c.k1 5

p a20.A 1
p 1 a17.α
p 3-1 a17.1i
p 3-1 a20.5i
s a17.op1 α
s a17.cc1 0
s a20.op5 A
s a20.cc5 0
s a20.rp5 1
p a20.5o 3-2
p 3-2 m.3i
p 2 a9.β
s m.ieracc3 β
s m.icandacc3 0
s m.sf3 10
s m.place3 10
s m.prod3 SC
s m.iercl3 C
s m.icandcl3 0
p m.3o 3-5
p 3-2 a18.3i
s a18.op3 A
s a18.cc3 0
p 3-5 a20.6i
s a20.op6 β
s a20.cc6 C
s a20.rp6 1
p a13.S 1
p 1 a20.β
p a20.6o 3-3
p 3-3 m.4i
s m.ieracc4 γ
s m.icandacc4 0
s m.sf4 10
s m.place4 10
s m.prod4 AC
s m.iercl4 C
s m.icandcl4 0
p 3-3 a17.2i
p a17.A 2
p 2 a9.γ
s a17.op2 A
s a17.cc2 C
p a13.A 1
p m.4o 3-6
p 3-6 a18.4i
s a18.op4 α
s a18.cc4 0
p 3-6 a15.5i
p a15.5o 2-5
s a15.op5 ε
s a15.cc5 C
s a15.rp5 1

#
# A trajectory calculation.  For this version, we're
# using a drag table I tried to figure out and simple
# rectangular integration with a delta-t of 0.01s.
#
p 1 a4.α
p 1 a6.α
p 3 a9.δ
p 2 a10.γ
p 3 a10.δ
p 3-10 a4.4i
s a4.op4 0
s a4.cc4 C
p 3-10 a6.6i
s a6.op6 0
s a6.cc6 C
s a6.rp6 1
p 3-10 a15.3i
s a15.op3 0
s a15.cc3 C
p 3-10 a18.11i
s a18.op11 0
s a18.cc11 C
p 3-10 a20.3i
s a20.op3 0
s a20.cc3 C
p 3-10 a10.10i
s a10.op10 0
s a10.cc10 C
s a10.rp10 1
p a10.10o 3-11
p 3-11 a4.1i	# Load the initial y' and x'
p 3-11 c.19i
s a4.op1 α
s c.s19 Glr
p c.19o 4-1
p 4-1 a15.12i
s a15.op12 ε
s a15.cc12 C
s a15.rp12 1
p 4-1 a6.2i
p 4-1 c.14i
s a6.op2 α
s c.s14 Flr
p c.14o 1-11
p 5-1 a4.3i	# Compute x'^2+y'^2 to look up the
p 5-1 m.7i	# drag coefficient from a table
s a4.op3 A
s a4.cc3 0
s m.ieracc7 β
s m.iercl7 C
s m.icandacc7 γ
s m.icandcl7 C
s m.place7 10
s m.sf7 7
s m.prod7 0
p m.7o 5-2
p 5-2 a6.3i
p 5-2 m.8i
s a6.op3 A
s a6.cc3 0
s m.ieracc8 δ
s m.iercl8 C
s m.icandacc8 δ
s m.icandcl8 C
s m.place8 10
s m.sf8 0
s m.prod8 AC
p m.8o 5-3
p 1 ad.s.6.-3
p ad.s.6.-3 a7.α
p 5-3 a7.5i
s a7.op5 α
s a7.cc5 0
s a7.rp5 1
p a7.5o 5-4	# 5-4 triggers the table lookup
p 5-4 f2.1i
p f2.C 5-5	# 5-5 reads the argument
p 5-5 a7.6i
p a7.A 1
p 1 f2.arg
p f2.A 1
s a7.op6 A
s a7.cc6 C
s a7.rp6 1
s f2.op1 A0
s f2.cl1 C
s f2.rp1 1
l drag.e
p a7.6o 4-3
p a4.A 2		# Update x and y from x' and y'
p a6.A 3
p 2 ad.s.4.-2
p ad.s.4.-2 a20.δ
p 3 ad.s.5.-2
p ad.s.5.-2 a18.δ
p 4-3 a20.7i
p 4-3 a4.2i
p 4-3 a18.12i
p 4-3 a6.1i
p 4-3 a10.11i
s a20.op7 δ
s a20.cc7 0
s a20.rp7 1
s a4.op2 A
s a4.cc2 0
s a18.op12 δ
s a18.cc12 0
s a18.rp12 1
s a6.op1 A
s a6.cc1 0
s a10.op11 γ
s a10.cc11 0
s a10.rp11 1
p a20.7o 5-6	# start a dummy program to delay the
p 5-6 a7.7i	# start of multiply until after ft read is
s a7.op7 0	# done
s a7.cc7 C
s a7.rp7 1
p a7.7o 4-4
p 4-4 m.5i
p 1 ad.s.7.4
p ad.s.7.4 a9.ε
s c.s8 Dlr
s m.ieracc5 ε
s m.iercl5 0
s m.icandacc5 0
s m.icandcl5 C
s m.sf5 0
s m.place5 6
s m.prod5 SC
p m.5o 4-5
p 4-5 a4.5i
s a4.op5 α
s a4.cc5 C
s a4.rp5 1
p a4.5o 4-6
p 4-6 a4.6i
p 4-6 c.13i
s a4.op6 α
s a4.cc6 C
s a4.rp6 1
s c.s13 Elr
p a4.6o 4-7
p 4-7 m.6i
s m.ieracc6 0
s m.iercl6 C
s m.icandacc6 δ
s m.icandcl6 C
s m.sf6 0
s m.place6 6
s m.prod6 SC
p 4-7 a6.4i
s a6.op4 A
s a6.cc4 0
p m.6o 4-8
p 4-8 a6.5i
s a6.op5 α
s a6.cc5 C
s a6.rp5 1
p a6.5o 1-4
p a20.S ad.dp.1.11
p ad.dp.1.11 5-11
p 5-10 a20.12i
s a20.op12 S
s a20.cc12 0
s a20.rp12 1
p 5-10 a15.4i
s a15.op4 ε
s a15.cc4 C
p 5-11 a19.5i
s a19.op5 0
s a19.cc5 0
s a19.rp5 1
p a19.5o 1-4

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

