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
p 1-4 a1.5i
p 1-4 a2.5i
p a1.A 1
p 1 a2.α
s a1.op5 A
s a2.op5 α
s p.cC 2
s p.a18 C
s p.a14 D
s p.d18s1 4
s p.d17s1 9
s p.d16s1 9
s p.d15s1 8
