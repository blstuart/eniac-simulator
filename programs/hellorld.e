#
# Usagi Electric's Hellorld challenge
#

# EBCDIC characters in decimal on the second function table
# Mark the end of the string with a negative value and use the
# sign to determine whether we're continuing or halting.
s f2.mpm1 T
s f2.RA0S P
s f2.RA0L3 2
s f2.RA0L2 0
s f2.RA0L1 0
s f2.RA1S P
s f2.RA1L3 1
s f2.RA1L2 3
s f2.RA1L1 3
s f2.RA2S P
s f2.RA2L3 1
s f2.RA2L2 4
s f2.RA2L1 7
s f2.RA3S P
s f2.RA3L3 1
s f2.RA3L2 4
s f2.RA3L1 7
s f2.RA4S P
s f2.RA4L3 1
s f2.RA4L2 5
s f2.RA4L1 0
s f2.RA5S P
s f2.RA5L3 1
s f2.RA5L2 5
s f2.RA5L1 3
s f2.RA6S P
s f2.RA6L3 1
s f2.RA6L2 4
s f2.RA6L1 7
s f2.RA7S P
s f2.RA7L3 1
s f2.RA7L2 3
s f2.RA7L1 2
s f2.RA8S P
s f2.RA8L3 0
s f2.RA8L2 9
s f2.RA8L1 0
s f2.RA9S M
s f2.RA9L3 0
s f2.RA9L2 0
s f2.RA9L1 1

# Printing lower half of Acc 18
s pr.12 P

# Use the initial pulse to start the ball rolling
p i.Io 1-1

# Set up the data trunks
p 1 f2.arg
p f2.A 1
p a17.A 1
p 1 a18.α
p a18.A 1

# Start a function table lookup
p 1-1 f2.1i
s f2.op1 A0
s f2.cl1 NC
s f2.rp1 1
p f2.NC 1-2

# Transmit the table argument from Acc 17
p 1-2 a17.5i
s a17.op5 A
s a17.cc5 0
s a17.rp5 1
p a17.5o 1-3

# Increment the table argument in parallel with lookup
p 1-3 a17.6i
s a17.op6 ε
s a17.cc6 C
s a17.rp6 1
p a17.6o 1-4

# Clear Acc 18 in parallel with increment and lookup
p 1-3 a18.1i
s a18.op1 0
s a18.cc1 C

# Delay to end of FT loolup
p 1-4 a17.7i
s a17.op7 0
s a17.cc7 0
s a17.rp7 1
p a17.7o 1-5

# Transfer FT output to Acc 18
p 1-5 a18.5i
s a18.op5 α
s a18.cc5 0
s a18.rp5 1
p a18.5o 1-6

# Transmit Acc 18 subtractively.  The adapter connects
# the sign to control line 1-7.  Because we're transmitting
# subtractively, there will be no pulses for a negative number
# and 9 pulses for a positive number.  The pulses are sent
# through a dummy program to synchronize them with the
# central programming pulse and its output triggers the
# printing.  Completion of the printing starts the whole cycle
# over again.
p 1-6 a18.6i
s a18.op6 S
s a18.cc6 0
s a18.rp6 1
p a18.S ad.dp.1.11
p ad.dp.1.11 1-7
p 1-7 i.Pi
p i.Po 1-1
