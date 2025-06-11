/*++

Copyright (C) 2025 Thread-safe library developers

All rights reserved.

Abstract: This is a stub class definition of CStringReturner

*/

#include "libthreadsafe_stringreturner.hpp"
#include "libthreadsafe_interfaceexception.hpp"

// Include custom headers here.


using namespace LibThreadSafe::Impl;

/*************************************************************************************************************************
 Class definition of CStringReturner 
**************************************************************************************************************************/

std::string CStringReturner::GetString()
{
	return std::string("Get random string");
}

void CStringReturner::ThreadSafetyCheck()
{
	values.emplace_back(values.size());
	std::reverse(values.begin(), values.end());
}

