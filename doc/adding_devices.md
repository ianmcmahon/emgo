My adventure in adding devices, by ianmcmahon

in egpath/stm32/xgen/stm32f4xx/ lives headers for each particular device
that look like they came from ST

xgen and stm32xgen are in tools

// stm32xgen generates STM32 peripheral files in xgen format.
//
// stm32xgen is usually used this wahy:
//  unifdef -k -f undef.h -D STM32TARGET stm32f4xx.h |stm32xgen PKGPATH

what the hell is -f? doesn't seem to work
pkgpath is egpath/src/stm32/o/f469xx/
without -f undef.h, it runs, spews a zillion errors, but produces files
