/*++

Copyright (C) 2019 PrimeDevelopers

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

void CCalculator::SetProgressCallback(const LibPrimes::ProgressCallback pProgressCallback)
{
	m_Callback = pProgressCallback;
}

LibPrimes_uint64 CCalculator::GetValue()
{
	return m_value;
}

void CCalculator::SetValue(const LibPrimes_uint64 nValue)
{
	m_value = nValue;
}

void CCalculator::Calculate()
{
	throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_NOTIMPLEMENTED);
}

