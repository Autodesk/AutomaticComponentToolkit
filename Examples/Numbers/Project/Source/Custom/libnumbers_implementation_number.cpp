/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is a stub class definition of CInternalLibNumbersNumber
Interface version: 1.0.0

**/

#include "libnumbers_implementation_number.hpp"
#include "libnumbers_interfaceexception.hpp"

// Include custom headers here.

#include <string>

using namespace LibNumbers;

/*************************************************************************************************************************
 Class definition of CInternalLibNumbersNumber 
**************************************************************************************************************************/

double CInternalLibNumbersNumber::Value ()
{
	throw ELibNumbersInterfaceException (LIBNUMBERS_ERROR_NOTIMPLEMENTED);
}

void CInternalLibNumbersNumber::SetValue (const double dValue)
{
	throw ELibNumbersInterfaceException (LIBNUMBERS_ERROR_NOTIMPLEMENTED);
}

void CInternalLibNumbersNumber::SetValueInt (const long long nValue)
{
	throw ELibNumbersInterfaceException (LIBNUMBERS_ERROR_NOTIMPLEMENTED);
}

void CInternalLibNumbersNumber::SetValueString (const std::string sValue)
{
	try
	{
		m_value = std::stod(sValue);
	}
	catch (std::invalid_argument)
	{
		throw ELibNumbersInterfaceException(LIBNUMBERS_ERROR_INVALIDPARAM);
	}
}

