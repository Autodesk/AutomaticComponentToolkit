/*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is a stub class definition of CCalculator

*/

#include "calculator_calculator.hpp"
#include "calculator_interfaceexception.hpp"

// Include custom headers here.
#include "calculator_variable.hpp"

using namespace Calculator::Impl;

/*************************************************************************************************************************
 Class definition of CCalculator 
**************************************************************************************************************************/

void CCalculator::EnlistVariable(IVariable* pVariable)
{
	m_pVariables.push_back(PIVariable(pVariable));
}

IVariable * CCalculator::GetEnlistedVariable(const Calculator_uint32 nIndex)
{
	if (nIndex < m_pVariables.size()) {
		return m_pVariables[nIndex].getCoOwningPtr();
	}
	else {
		throw ECalculatorInterfaceException(CALCULATOR_ERROR_INVALIDPARAM);
	}
}

void CCalculator::ClearVariables()
{
	m_pVariables.clear();
}

IVariable * CCalculator::Multiply()
{
	Calculator_double dValue = 1.0;
	for (const PIVariable & pVariable : m_pVariables) {
		dValue *= pVariable->GetValue();
	}
	return new CVariable(dValue);
}

IVariable * CCalculator::Add()
{
	Calculator_double dValue = 0.0;
	for (const PIVariable & pVariable : m_pVariables) {
		dValue += pVariable->GetValue();
	}
	return new CVariable(dValue);
}

