/*++

Copyright (C) 2025 Thread-safe library developers

All rights reserved.

Abstract: This is a stub class definition of CSoftStringReturner

*/

#include "libthreadsafe_softstringreturner.hpp"
#include "libthreadsafe_interfaceexception.hpp"

// Include custom headers here.


using namespace LibThreadSafe::Impl;

/*************************************************************************************************************************
 Class definition of CSoftStringReturner 
**************************************************************************************************************************/

std::string CSoftStringReturner::GetString()
{
	return std::string("Get random string");
}

void CSoftStringReturner::ThreadSafetyCheck()
{
	values.emplace_back(values.size());
	std::reverse(values.begin(), values.end());
}

