/*++

Copyright (C) 2019 Calculation developers

All rights reserved.

Abstract: This is a stub class definition of CCalculator

*/

#include "calculation_calculator.hpp"
#include "calculation_interfaceexception.hpp"

// Include custom headers here.


using namespace Calculation::Impl;

/*************************************************************************************************************************
 Class definition of CCalculator 
**************************************************************************************************************************/

void CCalculator::EnlistVariable(Numbers::PVariable pVariable)
{
	m_vVariables.push_back(pVariable);
}

Numbers::PVariable CCalculator::GetEnlistedVariable(const Calculation_uint32 nIndex)
{
	return m_vVariables[nIndex];
}

void CCalculator::ClearVariables()
{
	m_vVariables.clear();
}

Numbers::PVariable CCalculator::Multiply()
{
	double sum = 1.0;
	for (auto& pVar : m_vVariables) {
		sum += pVar->GetValue();
	}
	return CWrapper::sPNumbersWrapper->CreateVariable(sum);
}

Numbers::PVariable CCalculator::Add()
{
	double sum = 0.;
	for (auto& pVar : m_vVariables) {
		sum += pVar->GetValue();
	}
	return CWrapper::sPNumbersWrapper->CreateVariable(sum);
}

