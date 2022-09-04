package main

import (
    "fmt"
)

var unit string

// man pages
var initman string
var cyman string
var ctman string
var prman string
var accman string
var mulman string
var divman string
var ftman string
var mpman string
var aman string

func cmdref() {
    fmt.Println("b: Press a button\n" +
                "d: Display a unit status\n" +
                "D: Display all unit statuses\n" +
                "f: Associate a file with the punch or reader\n"+
                "l: Load a configuration file\n" +
                "h: Display the command reference\n" +
                "p: Plug in a jumper\n" +
                "q: Quit\n" +
                "r: Reset a single unit\n" +
                "R: Reset all units\n" +
                "s: Set a switch setting\n" +
                "u: Declare units being used\n" +
                "dt: Declare the number of data trunks in use\n" +
                "pt: Declare the number of program trunks in use\n" +
                "^C: Switch cycling unit to 1 pulse mode (pause execution)")
}
func unitsref() {
    fmt.Println("init: Initialization unit\n" +
                "cy: Cycling unit\n" +
                "ct: Constant transmitter\n" +
                "pr: Printer\n"+
                "acc: Accumulator\n" +
                "mul: High Speed Multiplier\n" +
                "div: Divider and Square Rooter\n" +
                "ft: Function Tables\n" +
                "mp: Master Programmer\n" +
                "a: Adapters")
}

func man(unit string){
    initman = `
    Initiating Unit

TERMINALS
    Cin: Selective clear input (1 ≤ n ≤ 6)
    Con: Selective clear ouput (1 ≤ n ≤ 6)
    Ri: Reader input
    Ro: Reader output
    Rl: Reader interlock
    Pi: Printer/punch input
    Po: Printer/punch output
    Pl: Printer/punch interlock
    Io: Initiating pulse output

BUTTONS
    c: Initial clear
    i: Initiating pulse
    r: Reader start switch

EXAMPLES
    1. Press the initiating pulse button: b i
    2. Plug a program jumper from the initiating pulse output of the initiating
    unit to program trunk 1, line 2: p i.Io 1-2

NOTES
    1. Although the clear, reader, printer, and initiate mnemonics are show as
    upper case here and on most of the original diagrams, either upper or
    accepted by the simulator.
    2. The terminals for Rs and Is are not implemented as these are
    nected to the hand-held control unit.
    3. The Start, Stop, and Door Shut switches are not implemented.`

    cyman = `
    Cycling unit

SWITCH
    op: Cycling unit operating mode: 1a (1 add), 1p (1 pulse), or co (continuous)

BUTTON
    p: 1 pulse & 1 add time button

EXAMPLES
    1. Set the cycling unit to 1 add mode: s cy.op 1a
    2. Single step in 1a or 1p modes: b p

NOTES
    1. The PA, Ext, 1A, and Cont terminals are unimplemented.
    2. The heater switch is unimplemented.
    3. The oscilloscope and oscilloscope input switch are unimplemented.
    4. Values for the operation switch may be specified in either upper or lower case.`

    ctman = `
    Constant Transmitter

TERMINALS
    o: Digit output terminal
    ni: Program input terminals (1≤n≤30)
    no: Program output terminals (1≤n≤30)

SWITCHES
    sn: Constant selector (1≤n≤30): Al, Ar, Alr, ..., Kl, Kr, Klr
    jl, jr:
    kl, kr: PM set: P or M
    jn, kn: Constant set (1≤n≤10): 0–9`

    prman = `
    Printer

SWITCHES
    n–m: Coupling (1≤n≤8, m=n+1): 0 or C
    n: Printing (1≤n≤16): O or P`

    accman = `
    Accumulator

TERMINALS
    α–ϵ: Input data terminals
    A: Additive output data terminal
    S: Subtractive output data terminal
    ni: Program input terminal (1≤n≤12)
    no: Program output terminal (5≤n≤12)
    Isn: Accumulator interconnect terminals (s ∈ {1, r}, n ∈ {1,2})
    lb: Special notation for a load block on interconnects

SWITCHES
    opn: Operation select (1≤n≤12): α–ϵ, 0, A, AS, or S
    rpn: Repeat (5≤n≤12): 1–9
    ccn: Clear correct (1≤n≤12): 0 or C
    sf: Significant figures: 0–10
    sc: Selective clear: 0 or SC

EXAMPLES
    1. Plug jumper from data trunk 3 to the gamma input on accumulator 12: p 3 a12.
    2. Set program 5 on accumulator 9 to output additively: s a9.op5 A
    3. Set accumulator 2 program 7 repeat to 4: s a2.rp7 4

NOTES
    1. Carry clear and selective clear switch settings may be specified in either
       upper or lower case.
    2. The input data terminals and input operation switch settings may be specified
       with either a Unicode Greek letter, the lower case spelled-out name of the
       Greek letter, or the corresponding lower case Roman letter.`

   mulman = `
    High Speed Multiplier

TERMINALS
    lhppI:
    lhppII:
    rhppI:
    rhppII: Partial product data terminals
    Rα–Rǫ: Multiplier accumulator program pulse outputs
    Dα–Dǫ: Multiplicand accumulator program pulse outputs
    A, S, AS:
    AC, SC, ASC: Product disposal terminals
    RS, DS, F: Internal operating terminals
    ni: Program input terminals (1 ≤ n ≤ 24)
    no: Program output terminals (1 ≤ n ≤ 24)

SWITCHES
    ieracc(n): Multiplier accumulator receive switch (1 ≤ n ≤ 24): α–ǫ or 0
    icandacc(n): Multiplicand accumulator receive switch (1 ≤ n ≤ 24): α–ǫ or 0
    sf(n): Significant figures switch (1 ≤ n ≤ 24): 0, 2–10
    place(n): Multiplier places switch (1 ≤ n ≤ 24): 2–10
    prod(n): Product disposal switch (1 ≤ n ≤ 24): A, S, AS, 0, AC, SC, or ASC
    iercl(n): Multiplier accumulator clear switch (1 ≤ n ≤ 24): 0 or C
    icandcl(n): Multiplicand accumulator clear switch (1 ≤ n ≤ 24): 0 or C

    Consult online reference for typical multiplier setup`

    divman = `
    Divider & Square Rooter

TERMINALS
    ni: Program input terminals (1 ≤ n ≤ 8)
    no: Program output terminals (1 ≤ n ≤ 8)
    nl: Program interlock terminals (1 ≤ n ≤ 8)
    ans: Digit answer output terminal

ADAPTERS (specified as switches)
    da: Divider adapter set: A, B, or C
    ra: Square Rooter adapter set: A, B, or C

SWITCHES
    nr(n): Numerator accumulator receive (1 ≤ n ≤ 8): α, β, or 0
    nc(n): Numerator accumulator clear (1 ≤ n ≤ 8): 0 or C
    dr(n): Denominator-square-root accumulator receive (1 ≤ n ≤ 8): α, β, or 0
    dc(n): Denominator-square-root accumulator clear (1 ≤ n ≤ 8): 0 or C
    pl(n): Places (1 ≤ n ≤ 8): D4, D7–10, S4, or S7–10
    ro(n): Round-off (1 ≤ n ≤ 8): RO or NRO
    an(n): Answer disposal (1 ≤ n ≤ 8): 1–4, or OFF
    il(n): Interlock (1 ≤ n ≤ 8): I or NI

    Consult online reference for typical divider setup

NOTES
    1. Only 10-digit numerators, denominators, and radicands are supported.
    2. Accumulator 2 is assumed for the quotient, Accumulator 3 for the numer-
    ator or radicand, Accumulator 5 for the denominator or square root, and
    Accumulator 7 for the shift accumulator.`

    ftman = `
    Function Tables

TERMINALS
    arg: Argument input termal
    A, B: Function output terminals
    NC, C: Argument reception NC and C program pulse output terminals
    ni: Program input terminals (1 ≤ n ≤ 11)
    no: Program output terminals (1 ≤ n ≤ 11)

SWITCHES
    op(n): Operation (1 ≤ n ≤ 11): A-2, A-1, ..., A+2, S-2, ..., S+2
    cl(n): Argument reception (1 ≤ n ≤ 11): 0, C, or NC
    rp(n): Operation repeat (1 ≤ n ≤ 11): 1–9
    mpm(n): Master PM switch (1 ≤ n ≤ 2): PorM
    AnC:
    BnC: Constant digit (1 ≤ n ≤ 4): 0–9, PM1, or PM2
    AnD:
    BnD: Digit delete (1 ≤ n ≤ 4): O or D
    AnS:
    BnS: Subtract pulse (4 ≤ n ≤ 10): 0 or S
    RArLd:
    RBrLd: Switch panel digit (−2 ≤ r ≤ 101, 1 ≤ d ≤6): 0–9
    RArS:
    RBrS: Switch panel sign (−2 ≤ r ≤ 101): P or M`

    mpman = `
    Master Programmer

TERMINALS
    ndi: Decade direct input (1 ≤ n ≤ 20)
    Adi–Kdi: Stepper direct input
    Ai–Ki: Stepper input
    Acdi–Kcdi: Stepper clear input
    Ano–Kno: Stepper output (1 ≤ n ≤ 6)

SWITCHES
    a(n): Decade associator switch (n ∈ {2, 4, 8, 10, 12, 14, 18, 20}): A–K
    dns(m): Decade switch (1 ≤ n ≤ 20, 1 ≤ m ≤ 6): 0–9
    cA–cK: Stepper clear: 1–6`

    aman = `
    Adapters

TYPES
    dp: Digit program pulse, e.g. ad.dp.1.11 for sign
    s: Shift, < 0 left, > 0 right
    d: Deleter, e.g. ad.d.1.7 for xxxxxxxxxx → xxxxxxx000
    sd: Special digit (used in A.G. Chapter 7)`

    switch unit {
        case "init": fmt.Println(initman)
        case "cy": fmt.Println(cyman)
        case "ct": fmt.Println(ctman)
        case "pr": fmt.Println(prman) 
        case "acc": fmt.Println(accman)
        case "mul": fmt.Println(mulman)
        case "div": fmt.Println(divman)
        case "ft": fmt.Println(ftman)
        case "mp": fmt.Println(mpman)
        case "a": fmt.Println(aman)
    }
}