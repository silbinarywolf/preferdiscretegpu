# Prefer Discrete GPU

[![Actions Status](https://github.com/silbinarywolf/preferdiscretegpu/workflows/Go/badge.svg)](https://github.com/silbinarywolf/preferdiscretegpu/actions)

**This is a CGo package. Unfortunately there is no way to enable this behaviour without CGo as of Go 1.14**

For laptops with multiple GPUs, its likely that applications won't launch using the discrete GPU on Windows, but rather the integrated GPU. This can lead to real-time applications like games under-performing by default, which is not the ideal user experience.

This library exports global variables for NVIDIA and AMD drivers so that they use high performance graphics rendering settings.

## Install

```
go get https://github.com/silbinarywolf/preferdiscretegpu
```

## Requirements

* Golang 1.12+

## How to use

Just import the package for side-effects. This will have no additional behaviour on non-Windows systems.

```go
package main

import (
	"fmt"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

func main() {
	// your code here
}
```

## How to test if this package is working

1) Open up the Windows Command Line (not Powershell) and run the following:

**NOTE: if you have have an older or newer version of Visual Studio installed, vcvarsall.bat will be in a different folder**
```
call "C:\Program Files (x86)\Microsoft Visual Studio 14.0\VC\vcvarsall.bat"
```

2) Build the test application that consumes this library
```
go build -v -o testapp.exe ./testapp
```

3) Run dumpbin on it:
```
dumpbin /exports testapp.exe
```

4) It should give you an output mentioning the NVIDIA/AMD constants, like so:
```
Microsoft (R) COFF/PE Dumper Version 14.00.24213.1
Copyright (C) Microsoft Corporation.  All rights reserved.


Dump of file testapp.exe

File Type: EXECUTABLE IMAGE

  Section contains the following exports for a.out.exe

    00000000 characteristics
    5F153BFB time date stamp Mon Jul 20 16:38:51 2020
        0.00 version
           1 ordinal base
           2 number of functions
           2 number of names

    ordinal hint RVA      name

          1    0 000B6C00 AmdPowerXpressRequestHighPerformance
          2    1 000B6C04 NvOptimusEnablement

  Summary

        1000 .CRT
       30000 .bss
       15000 .data
        1000 .debug_abbrev
        1000 .debug_aranges
       14000 .debug_frame
        1000 .debug_gdb_scripts
       82000 .debug_info
       35000 .debug_line
       88000 .debug_loc
        9000 .debug_pubnames
        E000 .debug_pubtypes
       30000 .debug_ranges
        1000 .debug_str
        1000 .edata
        1000 .idata
        1000 .pdata
       C7000 .rdata
        A000 .reloc
       A1000 .text
        1000 .tls
        1000 .xdata
```

## Credits

* [Hakan Guleryuz](https://groups.google.com/forum/#!topic/golang-nuts/7OHZcXUegF0) for documenting how they were able to export the NVIDIA/AMD variables in Golang and for giving information
* [SFML Contributors](https://github.com/SFML/SFML/commit/9a453ed9e3846e9f7998295b8966428a9a0b86f6#diff-93134bfcdd8e19cbd5fe05a57a658950R63) for inspiring how this feature should be consumed (ie. opt-in imported package)
