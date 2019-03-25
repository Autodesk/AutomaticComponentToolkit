/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is a stub class definition of CFactorizationCalculator

*/

#include "libprimes_factorizationcalculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes::Impl;

/*************************************************************************************************************************
 Class definition of CFactorizationCalculator 
**************************************************************************************************************************/

void CFactorizationCalculator::GetPrimeFactors(LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, LibPrimes::sPrimeFactor * pPrimeFactorsBuffer)
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

