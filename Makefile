test.ihx: test.azm Makefile
	z80asm -l -b test.azm
	z88dk-appmake +hex --org 0x9000 -b test.bin
	cat test.ihx
