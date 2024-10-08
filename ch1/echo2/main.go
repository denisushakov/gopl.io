// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	//s, sep := "", ""
	for idx, arg := range os.Args {
		//s += sep + arg
		//sep = " "
		fmt.Println(idx, arg)
	}
	//fmt.Println(s)
	secs := time.Since(t).Seconds()
	fmt.Println(secs)
}

//!-
