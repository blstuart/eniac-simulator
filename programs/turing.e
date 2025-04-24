# Set-up for generalized Turing Machine
#
# Accumulators
# a1-a11: Tape
# a12: Current tape position
#    A -> 2 -> (dp1 -> 2-1, dp2 -> 2-2)
# a13: Tape symbols ping-pong w/ a15,16
#    A -> 7
# a14: Countdown for ping-pong
#	A -> 4 -> dp3 -> 4-4
#	S -> 5 -> dp4 -> 4-5
# a15: Ping-pong w/ a13,16
#    A -> 7 ->s-1 -> 13γ
# a16: Dummy programs (9, 11, 12), ping-pong 2/ a13,15
# a17: Current state (FT arg)
#    A -> 3
# a18: Staging of transition function
#    A -> 1
# a19: Staging of current symbol for selecting function table
#    A -> 9 -> dp1 -> 4-9 -> Cdi
# a20: Dummy programs, hold halt indicator
#
# Digit Trunks
# 1: load init tape
# 2: xfer a12 digits to MP
# 3: FT args
# 4: transition transmission
# 6, 7: FT outputs
# 7: a13/15/16 ping-pong
# 8: xfer from tape to a13
# 9: feed tape symbol into MP
#

#
# We're defining the operation in terms of what we'll call a "Turing
# Cycle" which represents one application of the transition function.
# It follows the phases:
#	I.	Copy the accumulator holding the current tape cell to a13
#		The 10s digit of the head pointer (a12) identifies which acc
#	II.	Ping-pong between a13 and a15 to shift current symbol
#		The 1s digit of the had pointer (a12) identifes which digit
#	III.	Use current symbol to select FT
#	IV.	Use current state as FT argument
#	V.	Take FT output with new state, new sym, head dir
#	VI.	Ping-pong to load new sym into tape
#	VII.	Store the new tape accumulator
#	VIII. Inc/dec head pointer and store new state
#

#
# Set -1 on constant transmitter J for decrementing
#
s c.jl M
s c.jr P
s c.j1 9
s c.j2 9
s c.j3 9
s c.j4 9
s c.j5 9
s c.j6 9
s c.j7 9
s c.j8 9
s c.j9 9
s c.j10 9
s c.s25 Jlr
s c.s26 Jlr
s c.s27 Jr

#
# Prior to running the machine, we can preload the tape from a
# punched card.  If a card is loaded we will immediately transfer
# its contents into the first 8 accumulators
#
p c.o 1
p 1 a1.α
p 1 a2.α
p 1 a3.α
p 1 a4.α
p 1 a5.α
p 1 a6.α
p 1 a7.α
p 1 a8.α
p i.Ro 1-2
p 1-2 c.1i
s c.s1 Alr
p 1-2 a1.1i
s a1.op1 α
p c.1o 1-3
p 1-3 c.2i
s c.s2 Blr
p 1-3 a2.1i
s a2.op1 α
p c.2o 1-4
p 1-4 c.7i
s c.s7 Clr
p 1-4 a3.1i
s a3.op1 α
p c.7o 1-5
p 1-5 c.8i
s c.s8 Dlr
p 1-5 a4.1i
s a4.op1 α
p c.8o 1-6
p 1-6 c.13i
s c.s13 Elr
p 1-6 a5.1i
s a5.op1 α
p c.13o 1-7
p 1-7 c.14i
s c.s14 Flr
p 1-7 a6.1i
s a6.op1 α
p c.14o 1-8
p 1-8 c.19i
s c.s19 Glr
p 1-8 a7.1i
s a7.op1 α
p c.19o 1-9
p 1-9 c.20i
s c.s20 Hlr
p 1-9 a8.1i
s a8.op1 α

p i.io 1-1

#
# Phase I:
#	Read accumulator containing the current tape symbol.
#		Directly adapted from Ch 10 ENIAC Technical
#		Manual Part 1, by Adele G. Goldstine
#	Result goes into a13
#

s a12.sf 8
p a12.A 2
p 1-1 a12.7i
s a12.op7 A
s a12.cc7 0
s a12.rp7 1
p 2 ad.dp.1.3
p 2 ad.dp.2.2
p ad.dp.1.3 2-1
p ad.dp.2.2 2-2
p a12.7o 2-3
p 2-4 a12.5i
p a12.5o 2-6
s a12.op5 0
s a12.rp5 2
p 2-5 a12.6i
p a12.6o 2-6
s a12.op6 0
s a12.rp6 2

s p.a20 B
s p.cA 6
p 2-2 p.Adi
p 2-5 p.Ai
p 2-6 p.Acdi
p p.A1o 3-6
p p.A2o 3-7
p p.A3o 3-8
p p.A4o 3-9
p p.A5o 3-10
p p.A6o 3-11

s p.a12 D
s p.cE 2
s p.d11s1 1
s p.d11s2 1
p 2-1 p.Edi
p 2-3 p.Ei
p 2-6 p.Ecdi
p p.E1o 2-4
p p.E2o 2-5

s p.a10 G
s p.cF 6
p 2-2 p.Fdi
p 2-4 p.Fi
p 2-6 p.Fcdi
p p.F1o 3-2
p p.F2o 3-3
p p.F3o 3-4
p p.F4o 3-5
p p.F6o 3-1

#
# When reading out from a tape acc, we clear it in preparation
# for writing it back at the end of the Turing cycle
#
p a1.A 8
p 3-1 a1.5i
s a1.op5 A
s a1.cc5 0
s a1.rp5 1

p a2.A 8
p 3-2 a2.5i
s a2.op5 A
s a2.cc5 0
s a2.rp5 1

p a3.A 8
p 3-3 a3.5i
s a3.op5 A
s a3.cc5 0
s a3.rp5 1

p a4.A 8
p 3-4 a4.5i
s a4.op5 A
s a4.cc5 0
s a4.rp5 1

p a5.A 8
p 3-5 a5.5i
s a5.op5 A
s a5.cc5 0
s a5.rp5 1

p a6.A 8
p 3-6 a6.5i
s a6.op5 A
s a6.cc5 0
s a6.rp5 1

p a7.A 8
p 3-7 a7.5i
s a7.op5 A
s a7.cc5 0
s a7.rp5 1

p a8.A 8
p 3-8 a8.5i
s a8.op5 A
s a8.cc5 0
s a8.rp5 1

p a9.A 8
p 3-9 a9.5i
s a9.op5 A
s a9.cc5 0
s a9.rp5 1

p a10.A 8
p 3-10 a10.5i
s a10.op5 A
s a10.cc5 0
s a10.rp5 1

p a11.A 8
p 3-11 a11.5i
s a11.op5 A
s a11.cc5 0
s a11.rp5 1

p 8 a13.beta
p 2-3 a13.5i
s a13.op5 β
s a13.rp5 3
p a13.5o 4-1

#
# Phase II:
#	Ping-pong between a13 and a15 to shift the set of
#	10 cells so that the one we want is in the 1s place
#
#	Load counter (a14) from 1s digit of a12
#	repeat:
#		copy a13 to a15 and decrement a14 (4-1)
#		if a14 < 0 break (4-2, 4-4, 4-5)
#		copy a15 to a13 shifting right by 1 (4-3)
#

# Stepper B of MP set the total number of shifts
s p.a18 B
s p.cB 2
s p.d19s1 1
s p.d18s1 0
s p.d18s2 1
s p.d18s3 1
s p.d18s4 1
s p.d18s5 1
s p.d18s6 1
p 4-1 p.Bi
p p.B1o 4-2
p p.B2o 4-3

#Setting counter
p 2 ad.d.2.-9 
p ad.d.2.-9 a14.α
p 1 a14.β
p 1-1 a14.1i
s a14.op1 α

# copy a16 to a15, clear a16
p a16.A 7
p 7 a15.α
p 4-2 a16.1i
p 4-2 a15.5i
s a15.op5 α
s a15.cc5 0
s a15.rp5 1
s a16.op1 A
s a16.cc1 C
p a15.5o 4-4

# copy a15 back to a16 left shifting by 1
p a15.A 7
p 7 ad.s.1.1
p ad.s.1.1 a16.α
p 4-4 a15.6i
p 4-4 a16.2i
s a15.op6 A
s a15.cc6 C
s a15.rp6 1
s a16.op2 α
s a16.cc2 0
p a15.6o 4-5

# copy a13 to a15, clear a13, decrement a14
p a13.A 7
p 4-5 a13.1i
s a13.op1 A
s a13.cc1 C
p 4-5 a15.1i
s a15.op1 α
p 4-5 a14.6i
s a14.op6 β
s a14.rp6 1
p 4-5 c.25i
p a14.6o 4-6

# Copy a15 to a13 shifting left by 1
p 7 ad.s.2.1
p ad.s.2.1 a13.γ
p 4-6 a13.2i
s a13.op2 γ
p 4-6 a15.7i
s a15.op7 A
s a15.cc7 0
s a15.rp7 1
p a15.7o 4-7

# Send the sign of the a14 counter
p a14.A ad.dp.3.11
p a14.S ad.dp.4.11
p 4-7 a14.7i
s a14.op7 AS
s a14.rp7 1
p ad.dp.3.11 4-8
p 4-8 a16.11i
s a16.op11 0
s a16.rp11 1 
p ad.dp.4.11 4-9
p 4-9 a16.12i
s a16.op12 0
s a16.rp12 1
p a16.12o 4-11
p a16.11o 4-10

# Send high digit of a15 to low digit of a16
p 7 ad.s.3.-9
p ad.s.3.-9 a16.β
p 4-11 a15.8i
p 4-11 a16.3i
s a15.op8 A
s a15.cc8 C
s a15.rp8 1
s a16.op3 β
s a16.cc3 0
p a15.8o 4-1

#
# Phase III:
#	Send from a15 to a19 then to the MP
#	0: FT1 A
#	1: FT1 B
#	2: FT2 A
#	3: FT2 B
#	4: FT3 A
#	5: FT3 B
#
p 4-10 a15.9i
s a15.op9 A
s a15.cc9 C
s a15.rp9 1 
p 7 ad.s.5.-9
p ad.s.5.-9 a19.α
p 4-10 a19.5i
s a19.op5 α
s a19.rp5 1
p a19.5o 5-7
p a19.A 9
p 5-7 a19.6i
s a19.op6 A
s a19.rp6 1
s a19.cc6 C
p a19.6o 5-8
p 9 ad.dp.7.1
p ad.dp.7.1 5-9
p 5-9 p.Cdi
s p.a14 D
s p.cC 6
p 5-8 p.Ci
p 6-2 p.Ccdi
p p.C1o 5-1
p 5-1 f1.1i
p p.C2o 5-2
p 5-2 f1.2i
p p.C3o 5-3
p 5-3 f2.1i 
p p.C4o 5-4
p 5-4 f2.2i
p p.C5o 5-5
p 5-5 f3.1i
p p.C6o 5-6
p 5-6 f3.2i

#
# Phase IV:
#
# Send state number(a17) as argument
p a17.A 3
p 6-1 a17.1i
s a17.op1 A
s a17.cc1 C
p 3 f1.arg
p f1.C 6-1
p 3 f2.arg
p f2.C 6-1
p 3 f3.arg
p f3.C 6-1

s f1.op1 A0
s f1.cl1 C
s f1.rp1 1
s f1.op2 A0
s f1.cl2 C
s f1.rp2 1
s f2.op1 A0
s f2.cl1 C
s f2.rp1 1
s f2.op2 A0
s f2.cl2 C
s f2.rp2 1
s f3.op1 A0
s f3.cl1 C
s f3.rp1 1
s f3.op2 A0
s f3.cl2 C
s f3.rp2 1

#
# Phase V:
#
#Delaying the pulse going into a18 by 4 addition times using dummy programs
p 5-1 a20.7i
p 5-2 a20.8i
p 5-3 a20.9i
p 5-4 a20.10i
p 5-5 a20.11i
p 5-6 a20.12i

s a20.op7 0
s a20.rp7 4
s a20.op8 0
s a20.rp8 4
s a20.op9 0
s a20.rp9 4
s a20.op10 0
s a20.rp10 4
s a20.op11 0
s a20.rp11 4
s a20.op12 0
s a20.rp12 4

p a20.7o 5-10
p a20.8o 5-11
p a20.9o 5-10
p a20.10o 5-11
p a20.11o 5-10
p a20.12o 5-11

# a18 receives new value from function table's A or B output

p f1.A 7
p f1.B 6
p f2.A 7
p f2.B 6
p f3.A 7
p f3.B 6

p 7 a18.α
p 6 a18.β
p 5-10 a18.5i
s a18.op5 α
s a18.rp5 1
p 5-11 a18.6i
s a18.op6 β
s a18.rp6 1
p a18.5o 6-2
p a18.6o 6-2

# Copy the new symbol to a16(0)
# Simultaneously load a14 with a large number
p a18.A 4
p a18.S ad.dp.6.3
p ad.dp.6.3 6-3
p 6-2 a18.7i
s a18.op7 AS
s a18.cc7 C
s a18.rp7 1
p 4 ad.s.4.-3
p ad.s.4.-3 ad.d.4.-9
p ad.d.4.-9 a16.γ
p 6-2 a16.4i
s a16.op4 γ
s a16.cc4 0
p 6-2 c.27i
p 6-2 a14.10i
s a14.op10 β
s a14.cc10 0
s a14.rp10 1
p 4 a20.α
p 6-2 a20.1i
s a20.op1 α
s a20.cc1 0
p a18.7o 4-1
#
# Copy the new state to a17
#
p 4 a17.β
p 6-2 a17.2i
s a17.op2 β
s a17.cc2 0

#
# Send the decrement or Increment pulse
#
p 6-3 a19.12i
s a19.op12 0
s a19.cc12 0
s a19.rp12 1
p a19.12o 6-5
p 6-5 a12.9i
s a12.op9 ε
s a12.cc9 C
s a12.rp9 1

p 4 ad.dp.5.3
p ad.dp.5.3 a19.11i
s a19.op11 0
s a19.cc11 0
s a19.rp11 1
p a19.11o 6-6
p 1 a12.δ
p 6-6 a12.10i
s a12.op10 δ
s a12.rp10 1
p 6-6 c.26i

#
# Phase VII:
#
# Repeat the accumulator lookup for the correct
# tape accumulator.
#
s p.a8 H
s p.cH 6
p 2-2 p.Hdi
p 8-5 p.Hi
p 1-1 p.Hcdi
p p.H1o 7-6
p p.H2o 7-7
p p.H3o 7-8
p p.H4o 7-9
p p.H5o 7-10
p p.H6o 7-11

s p.a2 K
s p.cK 2
s p.d1s1 1
s p.d1s2 1
p 2-1 p.Kdi
p 4-3 p.Ki
p 1-1 p.Kcdi
p p.K1o 8-4
p p.K2o 8-5

s p.a4 J
s p.cJ 6
s p.d3s1 1
s p.d3s2 1
s p.d3s3 1
s p.d3s4 1
s p.d3s5 1
s p.d3s6 1
p 2-2 p.Jdi
p 8-4 p.Ji
p 1-1 p.Jcdi
p p.J1o 7-2
p p.J2o 7-3
p p.J3o 7-4
p p.J4o 7-5
p p.J6o 7-1

p 7 a1.β
p 7-1 a1.6i
s a1.op6 0
s a1.cc6 C
s a1.rp6 1
p a1.6o a1.7i
s a1.op7 β
s a1.cc7 0
s a1.rp7 1

p 7 a2.β
p 7-2 a2.6i
s a2.op6 0
s a2.cc6 C
s a2.rp6 1
p a2.6o a2.7i
s a2.op7 β
s a2.cc7 0
s a2.rp7 1

p 7 a3.β
p 7-3 a3.6i
s a3.op6 0
s a3.cc6 C
s a3.rp6 1
p a3.6o a3.7i
s a3.op7 β
s a3.cc7 0
s a3.rp7 1

p 7 a4.β
p 7-4 a4.6i
s a4.op6 0
s a4.cc6 C
s a4.rp6 1
p a4.6o a4.7i
s a4.op7 β
s a4.cc7 0
s a4.rp7 1

p 7 a5.β
p 7-5 a5.6i
s a5.op6 0
s a5.cc6 C
s a5.rp6 1
p a5.6o a5.7i
s a5.op7 β
s a5.cc7 0
s a5.rp7 1

p 7 a6.β
p 7-6 a6.6i
s a6.op6 0
s a6.cc6 C
s a6.rp6 1
p a6.6o a6.7i
s a6.op7 β
s a6.cc7 0
s a6.rp7 1

p 7 a7.β
p 7-7 a7.6i
s a7.op6 0
s a7.cc6 C
s a7.rp6 1
p a7.6o a7.7i
s a7.op7 β
s a7.cc7 0
s a7.rp7 1

p 7 a8.β
p 7-8 a8.6i
s a8.op6 0
s a8.cc6 C
s a8.rp6 1
p a8.6o a8.7i
s a8.op7 β
s a8.cc7 0
s a8.rp7 1

p 7 a9.β
p 7-9 a9.6i
s a9.op6 0
s a9.cc6 C
s a9.rp6 1
p a9.6o a9.7i
s a9.op7 β
s a9.cc7 0
s a9.rp7 1

p 7 a10.β
p 7-10 a10.6i
s a10.op6 0
s a10.cc6 C
s a10.rp6 1
p a10.6o a10.7i
s a10.op7 β
s a10.cc7 0
s a10.rp7 1

p 7 a11.β
p 7-11 a11.6i
s a11.op6 0
s a11.cc6 C
s a11.rp6 1
p a11.6o a11.7i
s a11.op7 β
s a11.cc7 0
s a11.rp7 1

p 4-3 a16.10i
s a16.op10 0
s a16.rp10 3
s a16.cc10 0
p a16.10o 8-2
p 8-2 a16.5i
s a16.op5 A
s a16.cc5 C
s a16.rp5 1
p 8-2 a14.9i
s a14.op9 0
s a14.rp9 1
s a14.cc9 C
p a16.5o 6-11

#
# Restart the cycle
#
p a20.A ad.dp.8.5
p ad.dp.8.5 8-1
p 6-11 a20.2i
s a20.op2 A
s a20.cc2 C
p 8-1 a19.10i
s a19.op10 0
s a19.cc10 0
s a19.rp10 1
p a19.10o 1-1

#
# Set MP counters all to avoid any issues with 0 counters triggering
#
s p.d15s1 1
s p.d15s2 1
s p.d15s3 1
s p.d15s4 1
s p.d15s5 1
s p.d15s6 1
s p.d12s1 1
s p.d12s2 1
s p.d12s3 1
s p.d12s4 1
s p.d12s5 1
s p.d12s6 1
s p.d11s1 1
s p.d11s2 1
s p.d11s3 1
s p.d11s4 1
s p.d11s5 1
s p.d11s6 1
s p.d9s1 1
s p.d9s2 1
s p.d9s3 1
s p.d9s4 1
s p.d9s5 1
s p.d9s6 1
s p.d5s1 1
s p.d5s2 1
s p.d5s3 1
s p.d5s4 1
s p.d5s5 1
s p.d5s6 1
s p.d1s1 1
s p.d1s2 1
s p.d1s3 1
s p.d1s4 1
s p.d1s5 1
s p.d1s6 1

#EOF
