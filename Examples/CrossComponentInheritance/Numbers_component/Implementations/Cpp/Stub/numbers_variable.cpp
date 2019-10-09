/*++

Copyright (C) 2019 Numbers developers

All rights reserved.

Abstract: This is a stub class definition of CVariable

*/

#include "numbers_variable.hpp"
#include "numbers_interfaceexception.hpp"

// Include custom headers here.


using namespace Numbers::Impl;

/*************************************************************************************************************************
 Class definition of CVariable 
**************************************************************************************************************************/

Numbers_double CVariable::GetValue()
{
	return m_dValue;
}

void CVariable::SetValue(const Numbers_double dValue)
{
	m_dValue = dValue;
}

