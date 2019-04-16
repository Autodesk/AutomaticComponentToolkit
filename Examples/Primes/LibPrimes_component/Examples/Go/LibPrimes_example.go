/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.6.0-develop.

Abstract: This is an autogenerated Go application that demonstrates the
 usage of the Go bindings of Prime Numbers Library

Interface version: 1.2.0

*/


package libprimesexample // TODO: change this to "main" to build

import (
	"fmt"
	"log"
	"../../Bindings/Go"
)

func main() {
	wrapper, err := libprimes.LibPrimesLoadWrapper("../../Bin/libprimes.") // TODO: add library-path and ending here
	if (err != nil) {
		log.Fatal(err)
	}
	nMajor, nMinor, nMicro, err := wrapper.GetVersion()
	if (err != nil) {
		log.Fatal(err)
	}
	versionString := fmt.Sprintf("libprimes.version = %d.%d.%d", nMajor, nMinor, nMicro)
	fmt.Println(versionString)

	calc, err := wrapper.CreateSieveCalculator()
	if (err != nil) {
		log.Fatal(err)
	}

	err = calc.SetValue(10)
	if (err != nil) {
		log.Fatal(err)
	}

	err = calc.Calculate()
	if (err != nil) {
		log.Fatal(err)
	}

	values, err := calc.GetPrimes()
	if (err != nil) {
		log.Fatal(err)
	}
	log.Fatal(values)
}

