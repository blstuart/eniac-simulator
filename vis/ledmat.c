#include <u.h>
#include <libc.h>
#include <bio.h>

void drive(void *);
void cleardisp(void);
void setpixel(int, int, int);

uchar buf[16][8];

void
main()
{
	int x, y, v;
	Biobuf bbuf;
	char *cmdstr, *toks[10];

	cleardisp();
	if(rfork(RFPROC|RFMEM)) {
		drive(nil);
		exits(nil);
	}
	Binit(&bbuf, 0, OREAD);
	while(1) {
		cmdstr = Brdstr(&bbuf, '\n', 1);
		if(cmdstr == 0)
			exits(nil);
		tokenize(cmdstr, toks, 10);
		if(toks[0][0] == 'c') {
			cleardisp();
		}
		else if(toks[0][0] == 's') {
			x = atoi(toks[1]);
			y = atoi(toks[2]);
			v = atoi(toks[3]);
			setpixel(x, y, v);
		}
		else {
			fprint(2, "unrecognized command %s\n", toks[0]);
		}
		free(cmdstr);
	}
}

void
cleardisp(void) {
	memset(buf, 0xff, 16*8);
}

void
setpixel(int x, int y, int v)
{
	int row, col, bit;

	row = y;
	col = x / 8;
	bit = 7 -(x % 8);
	if(v) {
		buf[row][col] &= ~(1 << bit);
	}
	else {
		buf[row][col] |= 1 << bit;
	}
}

/*
 * current wiring
 * GPIO	LED
 * 17		LA
 * 27		LB
 * 22		LC
 * 18		LD
 * 24		LAT
 * 25		EN
 *
 * SPI
 * SCLK	CLK
 * MOSI	R1/G1
 */
void
drive(void *)
{
	int fdg, spictl, spidat;
	int row;

	fdg = open("#G/gpio", OWRITE);
	if(fdg < 0) {
		perror("gpio open");
		exits(nil);
	}
	/* Just to be safe in case they got overridden */
	fprint(fdg, "function 10 alt0\n");
	fprint(fdg, "function 11 alt0\n");
	spictl = open("#π/spictl", ORDWR);
	spidat = open("#π/spi0", ORDWR);
	if(spictl < 0 || spidat < 0) {
		perror("spi open");
		exits(nil);
	}
	fprint(fdg, "function 17 out\n");
	fprint(fdg, "function 27 out\n");
	fprint(fdg, "function 22 out\n");
	fprint(fdg, "function 18 out\n");
	fprint(fdg, "function 24 out\n");
	fprint(fdg, "function 25 out\n");
	while(1) {
		for(row = 0; row < 16; row++) {
			if(row & 01)
				fprint(fdg, "set 17 1\n");
			else
				fprint(fdg, "set 17 0\n");
			if(row & 02)
				fprint(fdg, "set 27 1\n");
			else
				fprint(fdg, "set 27 0\n");
			if(row & 04)
				fprint(fdg, "set 22 1\n");
			else
				fprint(fdg, "set 22 0\n");
			if(row & 0x08)
				fprint(fdg, "set 18 1\n");
			else
				fprint(fdg, "set 18 0\n");
			pwrite(spidat, buf[row], 8, 0);
			fprint(fdg, "set 24 1\n");
			fprint(fdg, "set 24 0\n");
			fprint(fdg, "set 25 0\n");
			fprint(fdg, "set 25 1\n");
		}
		sleep(1);
	}
}
