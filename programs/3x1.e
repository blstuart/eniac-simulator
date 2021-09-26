# 3x+1 sequence
#
# while x >= 2:
#     print(x)
#     if x*5(0) == 0:
#         x *= 3
#         x += 1
#     else:
#         x = x / 2

# 1-1           a13.5   a 0 1 2-1 (a13 = x)
# 2-1           print   4-3
# 4-3           a3.7    a C 1 4-2
# 4-3           a13.1   S 0 1
# 4-2           a3.2    A C 1 (fires if a3 <=0, i.e. a13 >= 2)
# a3.A.dp.11    a3.6    A C 1 1-2

# 1-2           a1.5    a 0 5 (a1 = 5x)
# 1-2           a13.6   A 0 5 1-3
# 1-3           a1.2    A C 1
# 1-3           p.Acdi
# 1-3           a3.5    0 0 2 2-2
# a1.A.dp.1     a2.6    0 0 1 2-3 (2-3 fires if we're odd)
# (2-3 causes 1-4 to fire, otherwise 3-1 fires)
# 1-4           a2.5    a 0 3 (a2 = 3x)
# 1-4           a13.7   A C 3 1-5
# 1-5           a13.8   a C 1 (a13 = 3x+1)
# 1-5           a2.3    A C 1 2-1

# if even
# 3-1           a1.6    a 0 5
# 3-1           a13.9   A C 5 3-2
# 3-2           a13.10  b 0 1 2-1 (a13 = a1 * 5 / 10)
# 3-2           a1.3    A C 1

# initialize
p i.io 1-1
p 1-1 c.30i
p 1-1 a13.5i
p a13.5o 2-1

# top of loop, print
p 2-1 i.Pi
p i.Po 4-3
p 4-3 a13.1i
p 4-3 a3.7i
p a3.7o 4-2
p 4-2 a3.2i
p ad.dp.3.11 a3.6i
p a3.6o 1-2

# >= 2
p 1-2 a1.5i
p 1-2 a13.6i
p a13.6o 1-3
p 1-3 a1.2i
p 1-3 p.Acdi
p 1-3 a3.5i
p a3.5o 2-2
p 2-2 p.Ai
p ad.dp.1.1 a2.6i
p a2.6o 2-3
p 2-3 p.20di
p p.A1o 3-1
p p.A2o 1-4

# x is odd
p 1-4 a2.5i
p 1-4 a13.7i
p a13.7o 1-5
p 1-5 a2.3i
p 1-5 a13.8i
p a13.8o 2-1

# x is even
p 3-1 a1.6i
p 3-1 a13.9i
p a13.9o 3-2
p 3-2 a13.10i
p 3-2 a1.3i
p a13.10o 2-1

# digit bus setup
p a13.A 1
p 1 a1.a
p a1.A 2
p 2 ad.dp.1.1
p 1 a2.a
p a2.A 3
p 3 a13.a
p c.o 3
p 2 ad.s.1.-1
p ad.s.1.-1 a13.b
p a3.A ad.dp.3.11
p a13.S 4
p 4 a3.a

# switches
s pr.2 P
s pr.3 P

s c.s30 Klr
s c.k1 7

s p.cA 2
s p.d20s1 1

s a13.op1 S
s a13.op5 a
s a13.op6 A
s a13.rp6 5
s a13.op7 A
s a13.cc7 C
s a13.rp7 3
s a13.op8 a
s a13.cc8 C
s a13.op9 A
s a13.cc9 C
s a13.rp9 5
s a13.op10 b

s a1.op2 A
s a1.cc2 C
s a1.op3 A
s a1.cc3 C
s a1.op5 a
s a1.rp5 5
s a1.op6 a
s a1.rp6 5

s a2.op3 A
s a2.cc3 C
s a2.op5 a
s a2.rp5 3

s a3.op2 A
s a3.cc2 C
s a3.rp5 2
s a3.op6 A
s a3.cc6 C
s a3.op7 a
s a3.cc7 C

