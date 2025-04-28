C++_FLAGS=-I/usr/pkg/include/irrlicht -L/usr/pkg/lib

eniac3d: eniac3d.o
	g++ $(C++_FLAGS) -o eniac3d eniac3d.o -lIrrlicht

eniac3d.o: eniac3d.cpp
	g++ $(C++_FLAGS) -c eniac3d.cpp

eniacfp: eniacfp.o
	g++ $(C++_FLAGS) -o eniacfp eniacfp.o -lIrrlicht

eniacfp.o: eniacfp.cpp
	g++ $(C++_FLAGS) -c eniacfp.cpp

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
