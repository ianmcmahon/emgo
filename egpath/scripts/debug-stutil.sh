#!/bin/sh

set -e

arch=cortexm4f #`grep 'EGARCH=' ../build.sh |sed 's/.*EGARCH=\([[:alnum:]_]\+\).*/\1/g'`
if [ -z "$arch" ]; then
	arch=$EGARCH
fi

brkpnt=6
wchpnt=4

case "$arch" in
cortexm0)
	brkpnt=4
	wchpnt=2
	;;
cortexm7)
	brkpnt=8
	;;
esac

setsid st-util >/dev/null 2>&1 </dev/null &

trap /bin/true INT

arm-none-eabi-gdb --tui \
	-ex 'target extended-remote localhost:4242' \
	-ex 'set mem inaccessible-by-default off' \
	-ex "set remote hardware-breakpoint-limit $brkpnt" \
	-ex "set remote hardware-watchpoint-limit $wchpnt" \
	-ex 'set history save on' \
	-ex 'set history filename ~/.gdb-history-emgo' \
	-ex 'set history size 1000' \
	$arch.elf

killall st-util
