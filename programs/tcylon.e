#
# q=0: moving left
# q=1: moving right
#
# δ(0,0) = (1, 0, R)
# δ(0,1) = (0, 5, L)
# δ(1,0) = (0, 0, L)
# δ(1,5) = (1, 1, R)
#
s f1.RA0L1 1
s f1.RA0L2 0
s f1.RA0L3 0
s f1.RA0L4 0
s f1.RA0L5 9

s f1.RB0L1 0
s f1.RB0L2 0
s f1.RB0L3 9
s f1.RB0L4 5
s f1.RB0L5 9

s f1.RA1L1 0
s f1.RA1L2 0
s f1.RA1L3 9
s f1.RA1L4 0
s f1.RA1L5 9

s f3.RB1L1 1
s f3.RB1L2 0
s f3.RB1L3 0
s f3.RB1L4 1
s f3.RB1L5 9

l turing.e

f r cylon.card
b c
b r
