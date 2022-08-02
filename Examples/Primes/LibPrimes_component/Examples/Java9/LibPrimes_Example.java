/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.7.0-develop.

Abstract: This is an autogenerated Java application that demonstrates the
 usage of the Java bindings of Prime Numbers Library

Interface version: 1.2.0

*/


import libprimes.*;

public class LibPrimes_Example {

	public static String libpath = ""; // TODO add the location of the shared library binary here

	public static void main(String[] args) throws LibPrimesException {
		LibPrimesWrapper wrapper = new LibPrimesWrapper(libpath);
		
		LibPrimesWrapper.GetVersionResult version = wrapper.getVersion();
		System.out.print("LibPrimes version: " + version.Major + "." + version.Minor + "." + version.Micro);
		System.out.println();

		FactorizationCalculator factorization = wrapper.createFactorizationCalculator();
		factorization.setValue(735);
		factorization.calculate();
		PrimeFactor[] primeFactors = factorization.getPrimeFactors();
		String productString = "*";
		System.out.print(factorization.getValue() + " = ");
		for (int i = 0; i < primeFactors.length; i++) {
			PrimeFactor pF = primeFactors[i];
			if (i == primeFactors.length - 1) {
				productString = "\n";
			}
				
			System.out.print(" " + pF.Prime + "^" + pF.Multiplicity + " " + productString);
		}	
	}
}
