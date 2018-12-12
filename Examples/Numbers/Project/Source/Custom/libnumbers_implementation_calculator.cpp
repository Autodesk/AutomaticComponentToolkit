/*++

Copyright (C) 2018 Autodesk

All rights reserved.

Abstract: This is a stub class definition of CInternalLibNumbersCalculator
Interface version: 1.0.0

**/

#include "libnumbers_implementation_calculator.hpp"
#include "libnumbers_interfaceexception.hpp"

// Include custom headers here.


using namespace LibNumbers;

/*************************************************************************************************************************
 Class definition of CInternalLibNumbersCalculator 
**************************************************************************************************************************/

IInternalLibNumbersNumber * CInternalLibNumbersCalculator::AddNumbers (IInternalLibNumbersNumber* pValue1, IInternalLibNumbersNumber* pValue2)
{
	throw ELibNumbersInterfaceException (LIBNUMBERS_ERROR_NOTIMPLEMENTED);
}

