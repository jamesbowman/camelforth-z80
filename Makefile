camel80.ihx: *.asm.m4 Makefile
	zcc +rc2014 -o camel80 -subtype=none --no-crt -Ca '-l' -Cz '--org 0x8000' camel80.asm.m4 -create-app
