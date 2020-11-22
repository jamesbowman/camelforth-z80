test.ihx: test.asm.m4 Makefile
	zcc -Ca '-l -b' test.asm.m4
	z88dk-appmake +hex --org 0x9000 -b test.bin
	cat test.ihx
