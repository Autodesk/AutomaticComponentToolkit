/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is a stub class definition of CCalculator

*/

#include "libprimes_calculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes::Impl;

/*************************************************************************************************************************
 Class definition of CCalculator 
**************************************************************************************************************************/

LibPrimes_uint64 CCalculator::GetValue()
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

ICalculator * CCalculator::GetSelf()
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

void CCalculator::SetValue(const LibPrimes_uint64 nValue)
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

void CCalculator::Calculate()
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

void CCalculator::SetProgressCallback(const LibPrimes::ProgressCallback pProgressCallback)
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

