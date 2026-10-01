# ENIAC Simulator

This is a pulse-level simulator of the ENIAC.
The current state of the simulator covers all
of the functionality as originally constructed
at the University of Pennsylvania.
Future plans include implementing the enhancements
that were made at the Ballistics Research Lab
during its operational life.

The simulator is discussed in some detail in
[Simulating the ENIAC](http://ieeexplore.ieee.org/document/8326772/).
Further information on the ENIAC can be found
[here](http://cs.drexel.edu/~bls96).

## Build Instructions

The source code is divided into two main parts: the simulator core and
the graphical visualization support.
Source code for the simulator proper is in the `src` directory and that
for the graphical support is in the `vis` directory.
The graphical support is implemented as separate programs that run as
child processes under the simulator core.
With or without graphical support, the simulator provides a command-line
interface for control.

### Building the Core Simulator

Since the core is written in Go, you'll need the Go compiler installed.
With that, go into the `src` directory and issue the command:

> `go build -o eniac`

This will result in an executable called `eniac` that you can use to invoke
the simulator
Usually, things are run from the root of the repo, so if we want to run the prime
number sieve in the programs directory, we could use the command:

> `src/eniac sieve.e`

(Note that in some environments, you might have to use `./src/eniac`.)
The effect is that you will power up the machine with the configuration described
in the file `programs/sieve.e`, and
you will be greeted with a `>` prompt.
To begin the computation, you need to press the *Initiating Pulse Switch* on the
initiating unit.
In the simulator you do this with the `b i` command.
Doing so will cause the ENIAC to puch a series of cards, each holding a prime
number in the range of 3 to 2801.

### Building Visualization Support

At the present time, there are six different graphical support programs.
Which one(s) will be useful to you depends on the details of your environment.
They are as follows:

1. `ansi`: a classic simulator interface on a VT-100 terminal.
It's written in Go and requires no special support.
The screen is split into two regions: one where the state of all the units
in the machine are displayed, and one where the user can type commands.
The display continuously updates as the state changes.

2. `eniactk`: the original graphical interface.
It requires that you have TCL/Tk installed on your system and that the program `wish`
be in your path.
It uses the graphics in the images directory.
As it is written in Go, you can compile it with the command `go build -o eniactk eniactk.go`
or if you have `make` on your system, you can `make eniactk`.

3. `eniacfp`: provides a first-person video game type of interface.
It uses the 3d model in the `obj` directory.
To use this, you must have the Irrlicht library and a C++ compiler installed.
It is built with the command `make eniacfp`.

4. `eniac3d`: very similar to `eniacfp` except that it renders the view from both
eyes and presents them both.
It's only useful if you have a video display system that supports 3d display.
These displays usually require the use of LCD shutter glasses that alternate
which eye is open, and the display switches between which frame is displayed.
The command `make eniac3d` builds this one.

5. `eniact5`: another 3d experience.
This one works with the AR system from
[Tilt Five](http://www.tiltfive.com) and requires both the Irrlicht library
and the Tilt Five native SDK.
It is built with the command `make eniact5`.

6. `ledmat`: a driver for an LED matrix connected to a Raspberrhy Pi running
Plan 9.
Some discussion and a brief demonstration of this was presented at the
[11th International Workshop on Plan 9](http://iwp9.org).

To run the simulator with visualization, add the `-v` option to the
command.
For example, if we wanted to run the prime number sieve with the TCL/Tk
graphical support, we'd use the command:

> `src/eniac -v vis/eniactk sieve.e`

(assuming all the executables are in their respective source directories).
