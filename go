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

zcc +rc2014 -o camel80 -subtype=cpm --no-crt -Ca '-l' -Cz '--org 0x9000' camel80.asm -create-app
xclip -i camel80.ihx
# python scmload.py $HOME/rc2014 a.ihx
