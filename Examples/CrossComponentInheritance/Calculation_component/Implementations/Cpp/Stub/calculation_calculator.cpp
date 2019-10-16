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
	m_pVariables.push_back(pVariable);
}

Numbers::PVariable CCalculator::GetEnlistedVariable(const Calculation_uint32 nIndex)
{
	if (nIndex >= m_pVariables.size()) {
		throw ECalculationInterfaceException(CALCULATION_ERROR_INVALIDPARAM);
	}
	return m_pVariables[nIndex];
}

void CCalculator::ClearVariables()
{
	m_pVariables.clear();
}

Numbers::PVariable CCalculator::Multiply()
{
	auto pProd = CWrapper::sPNumbersWrapper->CreateVariable(1.0);
	for (auto pVar : m_pVariables) {
		pProd->SetValue(pProd->GetValue() * pVar->GetValue());
	}
	return pProd;
}

Numbers::PVariable CCalculator::Add()
{
	auto pProd = CWrapper::sPNumbersWrapper->CreateVariable(0.0);
	for (auto pVar : m_pVariables) {
		pProd->SetValue(pProd->GetValue() + pVar->GetValue());
	}
	return pProd;
}

