/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

Abstract: This is a stub class definition of CLibPrimesCalculator

*/

#include "libprimes_calculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes::Impl;

/*************************************************************************************************************************
 Class definition of CLibPrimesCalculator 
**************************************************************************************************************************/

CLibPrimesCalculator::CLibPrimesCalculator()
	:m_value(0) , m_Callback(nullptr)
{

}


LibPrimes_uint64  CLibPrimesCalculator::GetValue()
{
	return m_value;
}

void CLibPrimesCalculator::SetValue(const LibPrimes_uint64  nValue)
{
	m_value = nValue;
}

void CLibPrimesCalculator::SetProgressCallback (const LibPrimesProgressCallback pProgressCallback)
{
	m_Callback = pProgressCallback;
}

