/*++

Copyright (C) 2020 ADSK

All rights reserved.

Abstract: This is a stub class definition of CTiger

*/

#include "rtti_tiger.hpp"
#include "rtti_interfaceexception.hpp"

// Include custom headers here.


using namespace RTTI::Impl;

/*************************************************************************************************************************
 Class definition of CTiger 
**************************************************************************************************************************/

void CTiger::Roar()
{
	throw ERTTIInterfaceException(RTTI_ERROR_NOTIMPLEMENTED);
}

