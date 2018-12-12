/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

Abstract: This is a stub class definition of CLibPrimesCalculator

*/

#include "libprimes_calculator.hpp"
#include "libprimes_interfaceexception.hpp"

// Include custom headers here.


using namespace LibPrimes;

/*************************************************************************************************************************
 Class definition of CLibPrimesCalculator 
**************************************************************************************************************************/

CLibPrimesCalculator::CLibPrimesCalculator()
	:m_value(0) , m_Callback(nullptr)
{

}


unsigned long long CLibPrimesCalculator::GetValue()
{
	return m_value;
}

void CLibPrimesCalculator::SetValue(const unsigned long long nValue)
{
	m_value = nValue;
}

void CLibPrimesCalculator::SetProgressCallback (const LibPrimesProgressCallback pProgressCallback)
{
	m_Callback = pProgressCallback;
}


