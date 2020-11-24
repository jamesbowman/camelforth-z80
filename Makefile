camel80.ihx: *.asm.m4 Makefile
	zcc +rc2014 -o camel80 -subtype=cpm --no-crt -Ca '-l' -Cz '--org 0x9000' camel80.asm.m4 -create-app
