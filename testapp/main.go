package main

import (
	"fmt"

	_ "github.com/silbinarywolf/preferdiscretegpu"
)

func main() {
	fmt.Printf(`
	This application should not be run to determine if the NVIDIA/AMD flags were set.
	It should have "dumpbin" run against it and return something like the following:
	
	Dump of file defaultgputest.exe

	File Type: EXECUTABLE IMAGE
	
	  Section contains the following exports for a.out.exe
	
		00000000 characteristics
		5F15303A time date stamp Mon Jul 20 15:48:42 2020
			0.00 version
			   1 ordinal base
			   2 number of functions
			   2 number of names
	
		ordinal hint RVA      name
	
			  1    0 0006B820 AmdPowerXpressRequestHighPerformance
			  2    1 0006B824 NvOptimusEnablement
	
	  Summary
	
			1000 .CRT
		   30000 .bss
			4000 .data
			1000 .debug_abbrev
			1000 .debug_aranges
			C000 .debug_frame
			1000 .debug_gdb_scripts
		   56000 .debug_info
		   22000 .debug_line
		   5C000 .debug_loc
			3000 .debug_pubnames
			A000 .debug_pubtypes
		   20000 .debug_ranges
			1000 .debug_str
			1000 .edata
			1000 .idata
			1000 .pdata
		   7E000 .rdata
			6000 .reloc
		   67000 .text
			1000 .tls
			1000 .xdata

	`)
}
