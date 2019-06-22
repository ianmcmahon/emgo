#!/bin/sh

set -e
 
arch=cortexm4f
if [ -z "$arch" ]; then
	arch=$EGARCH
fi

arm-none-eabi-objcopy -O binary $arch.elf $arch.bin

addr=0x20000000
if grep -q '^INCLUDE.*/loadflash' script.ld; then
	addr=0x8000000
fi

echo "Loading at $addr..."
st-flash --reset write $arch.bin $addr
