set -e

tput clear
echo "----------------------------------"
PATH=$PATH:$HOME/git/z88dk/bin
export ZCCCFG=$HOME/git/z88dk/lib/config
# make

rm -f a.*
# zcc +rc2014 -subtype=basic --no-crt test.asm.m4 -create-app
m4 camel80.asm.m4 >  camel80.asm
z88dk-z80asm -mz80 -l camel80.asm

zcc +rc2014 -v -subtype=cpm --no-crt -Ca '-l' -Cz '--org 0x9000' camel80.asm -create-app
# hd a.bin

# z88dk-z80asm -mz80 -b -d  -o"a.bin" -m -s -L.  -L"/home/jamesb/git/z88dk/lib/config/../..//libsrc/_DEVELOPMENT/lib/sdcc_iy"   -D__SDCC_IY -I"/home/jamesb/git/z88dk/lib/config/../..//libsrc/_DEVELOPMENT/target/rc2014" -I"/tmp/tmpzccXXV65EvV" -D__SDCC -lrc2014    "/tmp/tmpXXjBlsJV.o"
# z88dk-appmake --org 0x9000 +rom --ihex  -b "a.bin" -c "/tmp/tmpXXjBlsJV"

