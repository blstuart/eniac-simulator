C++_FLAGS=-I/usr/pkg/include/irrlicht -L/usr/pkg/lib

ansi: ansi.go
	go build -o ansi ansi.go

eniac3d: eniac3d.o
	g++ $(C++_FLAGS) -o eniac3d eniac3d.o -lIrrlicht

eniac3d.o: eniac3d.cpp
	g++ $(C++_FLAGS) -c eniac3d.cpp

eniacfp: eniacfp.o
	g++ $(C++_FLAGS) -o eniacfp eniacfp.o -lIrrlicht

eniacfp.o: eniacfp.cpp
	g++ $(C++_FLAGS) -c eniacfp.cpp

eniact5: eniact5.o
	g++ $(C++_FLAGS) -o eniac5t eniac5t.o -lIrrlicht

eniact5.o: eniact5.cpp
	g++ $(C++_FLAGS) -I/usr/pkg/include -I../../../packages/tiltfive_sdk_1.4.1/Native -c eniact5.cpp

eniacrlfp: eniacrlfp.o
	cc -L/usr/pkg/lib -o eniacrlfp eniacrlfp.o -lraylib

eniacrlfp.o: eniacrlfp.c
	cc -I/usr/pkg/include -c eniacrlfp.c

eniactk: eniactk.go
	go build -o eniactk eniactk.go

irrtest: irrtest.o
	g++ $(C++_FLAGS) -o irrtest irrtest.o -lIrrlicht

irrtest.o: irrtest.cpp
	g++ $(C++_FLAGS) -c irrtest.cpp

irrtest2: irrtest2.o
	g++ $(C++_FLAGS) -o irrtest2 irrtest2.o -lIrrlicht

irrtest2.o: irrtest2.cpp
	g++ $(C++_FLAGS) -c irrtest2.cpp
