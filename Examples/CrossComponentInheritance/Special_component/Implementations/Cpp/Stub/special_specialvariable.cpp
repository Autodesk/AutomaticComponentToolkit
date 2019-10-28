/*++

Copyright (C) 2019 Special Numbers developers

All rights reserved.

Abstract: This is a stub class definition of CSpecialVariable

*/

#include "special_specialvariable.hpp"
#include "special_interfaceexception.hpp"

// Include custom headers here.


using namespace Special::Impl;

/*************************************************************************************************************************
 Class definition of CSpecialVariable 
**************************************************************************************************************************/

Special_double CSpecialVariable::GetValue()
{
	return Special_double(GetSpecialValue());
}

void CSpecialVariable::SetValue(const Special_double dValue)
{
	m_dValue = dValue;
}

Special_int64 CSpecialVariable::GetSpecialValue()
{
	return Special_int64(m_dValue);
}

