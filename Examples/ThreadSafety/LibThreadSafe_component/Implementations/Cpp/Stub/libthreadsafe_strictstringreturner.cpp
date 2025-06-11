/*++

Copyright (C) 2025 Thread-safe library developers

All rights reserved.

Abstract: This is a stub class definition of CStrictStringReturner

*/

#include "libthreadsafe_strictstringreturner.hpp"
#include "libthreadsafe_interfaceexception.hpp"

// Include custom headers here.


using namespace LibThreadSafe::Impl;

/*************************************************************************************************************************
 Class definition of CStrictStringReturner 
**************************************************************************************************************************/

std::string CStrictStringReturner::GetString()
{
	return std::string("Get random string");
}

void CStrictStringReturner::ThreadSafetyCheck()
{
	values.emplace_back(values.size());
	std::reverse(values.begin(), values.end());
}

