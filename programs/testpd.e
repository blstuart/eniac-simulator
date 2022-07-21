# Load Acc1 from CT

p i.Io 1-1
p 1 a1.α
p 1-1 a1.1i
s a1.op1 α
s a1.cc1 0
p c.o 1
p 1-1 c.25i
s c.s25 Jlr
p c.25o 1-2

s c.j10 1
s c.j9 2
s c.j8 3
s c.j7 4
s c.j6 5
s c.j5 6
s c.j4 7
s c.j3 8
s c.j2 9
s c.j1 0

# Send from Acc1 to Acc2

p 1-2 a1.2i
s a1.op2 A
s a1.cc2 0
p 1-2 a2.1i
s a2.op1 α
s a2.cc1 0

# Swap the upper and lower 5 digits going from Acc1 to trunk

p a1.A ad.dp.1.1
p ad.dp.1.11 ad.pd.1.11
p ad.dp.1.10 ad.pd.1.5
p ad.dp.1.9 ad.pd.1.4
p ad.dp.1.8 ad.pd.1.3
p ad.dp.1.7 ad.pd.1.2
p ad.dp.1.6 ad.pd.1.1
p ad.dp.1.5 ad.pd.1.10
p ad.dp.1.4 ad.pd.1.9
p ad.dp.1.3 ad.pd.1.8
p ad.dp.1.2 ad.pd.1.7
p ad.dp.1.1 ad.pd.1.6
p ad.pd.1.1 2
p 2 a2.α
