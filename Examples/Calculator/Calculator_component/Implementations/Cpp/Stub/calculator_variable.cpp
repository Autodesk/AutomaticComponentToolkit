/*++

Copyright (C) 2019 Calculator developers

All rights reserved.

Abstract: This is a stub class definition of CVariable

*/

#include "calculator_variable.hpp"
#include "calculator_interfaceexception.hpp"

// Include custom headers here.


using namespace Calculator::Impl;

/*************************************************************************************************************************
 Class definition of CVariable 
**************************************************************************************************************************/

CVariable::CVariable(Calculator_double dValue)
	: m_dValue(dValue)
{

}

Calculator_double CVariable::GetValue()
{
	return m_dValue;
}

void CVariable::SetValue(const Calculator_double dValue)
{
	m_dValue = dValue;
}

