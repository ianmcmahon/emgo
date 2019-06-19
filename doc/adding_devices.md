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

do it in egpath/src:

  emgo/egpath/src#  unifdef -k -D STM32TARGET stm32/xgen/stm32f4xx/stm32f469xx.h|stm32xgen stm32/o/f469xx &> stm32/o/f469xx/stm32xgen_output.log

it's critical to provide the pkgpath the same way that it'll be an import path

run xgen on all the files to produce xgen_file.go

  xgen stm32/o/f469xx/*/*.go

now somehow those files need to make their way into stm32/hal/raw/<peripheral>/f69xx--*.go
Rnot sure if there's an automated tool to do this but this works:

  for i in stm32/o/f469xx/*/*.go; do FNAME=`basename $i`; P=`dirname $i`; DIRNAME=`basename $P`; D=`dirname $P`; DEVICE=`basename $D`; echo "cp $i stm32/hal/raw/$DIRNAME/$DEVICE--$FNAME"; done
